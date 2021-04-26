import React, {useCallback, useContext, useEffect, useState} from "react";
import {Card, Col, Divider, message, Row, Switch, Typography} from 'antd';
import {Link, useParams} from 'react-router-dom';
import {GlobalContext} from '../context/GlobalProvider';
import client from '../lib/client-fetch';

export default function StatisticsPage(props) {
    const { state: { user } } = useContext(GlobalContext);
    let { code } = useParams();
    const [loading, setLoading] = useState(false);
    const [shortStats, setShortStats] = useState(null);

    const fetchDataRaw = async () => {
        if (user === null || !user.token) {
            return;
        }

        setLoading(true);
        try {
            const result = await client.get('/api/shorten/' + code + "/stats", {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            setShortStats(result.data);
        } catch (error) {
            if (error.status === 404) {
                setShortStats({});
            } else {
                message.error('Request failed: ' + error.message);
            }
        }
        setLoading(false);
    };
    const fetchData = useCallback(fetchDataRaw, [user]);

    useEffect(() => {
        fetchData();
    }, [user, code, fetchData]);

    return (
        <Row>
            <Col offset={6} span={12}>
                <Typography.Title>
                    Access Logs for "{code}"
                </Typography.Title>
                <Divider/>
            </Col>
        </Row>
    );
}
