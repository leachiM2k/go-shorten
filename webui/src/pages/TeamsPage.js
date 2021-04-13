import React, {useContext, useEffect, useState} from "react";
import {Col, Divider, Empty, Row, Spin, Tree, Typography} from 'antd';

import TeamMembers from '../components/teams/members/TeamsPage';
import client from '../lib/client-fetch';
import {GlobalContext} from '../context/GlobalProvider';

export default function TeamsPage(props) {
    const [componentState, setComponentState] = useState({ selectedKeys: [] });
    const [teamMembers, setTeamMembers] = useState([]);

    const { state, apiGetTeamStructure } = useContext(GlobalContext);
    const teamStructure = state.teams;
    const loading = state.teamsLoading;

    useEffect(() => { apiGetTeamStructure() }, [apiGetTeamStructure]);

    const handleTreeSelect = (selectedTeam, { selected, selectedNodes, node, event }) => {
        client.get('/api/teams/' + selectedTeam + '/members')
            .then(result => {
                if (!result.data) {
                    return;
                }
                setTeamMembers({
                    ...teamMembers, [selectedTeam]: result.data.map(item => ({
                        name: item.name,
                        p: item.pnumber
                    }))
                });
            })
            .catch(err => console.error("Could not load: ", err));


        setComponentState({ selectedTeam: { key: node.key, title: node.title } });
    }

    const handleAddMembers = (newMembers) => {
        client.put('/api/teams/' + componentState.selectedTeam.key + '/members', newMembers.map(item => item.value))
            .then(result => console.log("Saved successfully", result))
            .catch(err => console.error("Could not save: ", err));
        const currentTeamMembers = (teamMembers[componentState.selectedTeam.key] || []).concat(newMembers.map(item => ({
            name: item.label,
            p: item.value
        })));
        setTeamMembers({ ...teamMembers, [componentState.selectedTeam.key]: currentTeamMembers });
    }

    const handleRemoveMember = memberToRemove => {
        client.delete('/api/teams/' + componentState.selectedTeam.key + '/members', { data: [memberToRemove.p] })
            .then(result => console.log("Saved successfully", result))
            .catch(err => console.error("Could not save: ", err));
        const currentTeamMembers = (teamMembers[componentState.selectedTeam.key] || []).filter(members => members.p !== memberToRemove.p);
        setTeamMembers({ ...teamMembers, [componentState.selectedTeam.key]: currentTeamMembers });
    }

    const renderTeamTree = () => {
        if (loading || !teamStructure) {
            return (<Spin/>);
        }
        if (teamStructure.length < 1) {
            return (<Empty description="Could not find any team tree for your user"/>);
        }
        return (
            <Tree
                defaultExpandAll={true}
                onSelect={handleTreeSelect}
                defaultSelectedTeam={componentState.selectedTeam && componentState.selectedTeam.key}
                treeData={teamStructure}
            />
        );
    }

    return (
        <div>
            <Typography.Title>
                Teams Assignment
            </Typography.Title>
            <Divider/>
            <Row>
                <Col>
                    {renderTeamTree()}
                </Col>
                <Col flex={1}>
                    {
                        Array.isArray(teamStructure) && teamStructure.length > 0 &&
                        <TeamMembers selectedTeam={componentState.selectedTeam}
                                     teamMembers={teamMembers[componentState.selectedTeam && componentState.selectedTeam.key] || []}
                                     onAddMembers={handleAddMembers} onRemoveMember={handleRemoveMember}/>
                    }
                </Col>
            </Row>
        </div>
    );
}
