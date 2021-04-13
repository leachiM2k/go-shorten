import React, {useEffect, useState} from "react";
import {Card, Divider, message, Spin, Switch, Typography} from 'antd';
import client from '../lib/client-fetch';
import {Link} from 'react-router-dom';

export default function HeimdallPage(props) {
    let [missingSshKey, setMissingSshKey] = useState(false);
    let [SshMe, setSshMe] = useState(false);
    let [loading, setLoading] = useState(true);
    let [TTL, setTTL] = useState(null);

    const makeDataRequest = async (method) => {
        setLoading(true);
        try {
            const result = await client[method]('/api/heimdall/key/active');
            // const result = { data: "2020-07-02T09:31:23.812096957Z" };
            if (result.data && typeof result.data === 'string') {
                const calculatedTTL = Math.floor((new Date(result.data) - new Date()) / 1000);
                if (calculatedTTL > 0) {
                    setTTL(calculatedTTL);
                    setSshMe(true);
                } else {
                    setTTL(0);
                    setSshMe(false);
                }
            } else {
                setTTL(0);
                setSshMe(false);
            }
        } catch (error) {
            setTTL(0);
            if (error.status === 404) {
                setMissingSshKey(true);
            } else if (error.status !== 410) {
                message.error('Request failed: ' + error.message);
            }
        }
        setLoading(false);
    }
    const enableSshMe = () => makeDataRequest('post');
    const disableSshMe = () => makeDataRequest('delete');

    useEffect(() => {
        makeDataRequest('get');
    }, []);

    const handleSshMe = checked => {
        if (checked) {
            return enableSshMe();
        } else {
            return disableSshMe();
        }
    };

    const renderSwitch = () => {
        if(TTL === null) {
            return <Spin/>;
        }

        if (missingSshKey) {
            return (
                <Card title="SSH Me">
                    <Typography.Paragraph>No SSH Key found for your user. Please go to your <Link to="/me">profile
                        page</Link> and add a SSH Key.</Typography.Paragraph>
                </Card>
            );
        }

        return (
            <Card title="SSH Me">
                <Typography.Paragraph>Provide SSH access rights for 6 hours. <strong>All changes are
                    logged.</strong></Typography.Paragraph>
                <Typography.Paragraph>
                    <Switch checkedChildren="on" unCheckedChildren="off" loading={loading} checked={SshMe}
                            onChange={handleSshMe}/> SSH Access Permitted
                </Typography.Paragraph>
            </Card>
        );
    }

    return (
        <div>
            <Typography.Title>
                Heimdall SSH Key Management
            </Typography.Title>
            <Divider/>
            {renderSwitch()}
        </div>
    );
}
