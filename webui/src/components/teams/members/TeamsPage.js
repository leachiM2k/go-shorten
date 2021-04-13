import React, {useState} from 'react';
import {Button, Col, Empty, List, Row, Typography} from 'antd';
import UserSelect from '../../UserSelect';

const TeamMembers = props => {
    const [selected, setSelected] = useState(props.teamMembers || []);

    const handleSelect = (item) => {
        setSelected([...selected, item]);
    }

    const handleAddMembers = () => {
        props.onAddMembers && props.onAddMembers(selected);
        setSelected([]);
    }

    const handleRemoveMember = (item) => {
        props.onRemoveMember && props.onRemoveMember(item);
    }

    if (!props.selectedTeam) {
        return <div>
            <p>Please select a team</p>
            <Empty/>
        </div>;
    }
    return (
        <div>
            <Typography.Title level={2}>Members of "{props.selectedTeam.title}"</Typography.Title>
            <List
                bordered
                header={<Row>
                    <Col flex={1}><UserSelect selected={selected} onSelect={handleSelect}/></Col>
                    <Col><Button type="primary" onClick={handleAddMembers}>Add members</Button></Col>
                </Row>}
                dataSource={props.teamMembers}
                renderItem={item => (
                    <List.Item actions={[<Button type="danger"
                                                 size="small"
                                                 onClick={() => handleRemoveMember(item)}>remove</Button>]}>{item.name} ({item.p})</List.Item>
                )}
            />
        </div>
    );
};

export default TeamMembers;
