import React, {useCallback, useContext, useEffect, useState} from "react";
import {useLocation} from "react-router-dom";
import {Col, Row, Spin, Typography} from 'antd';
import client from '../lib/client-fetch';
import {GlobalContext} from '../context/GlobalProvider';

export function NotFound() {
    const location = useLocation();
    const { state: { user } } = useContext(GlobalContext);
    const [loading, setLoading] = useState(true);

    const fetchDataRaw = async () => {
        if (user === null || !user.token) {
            setLoading(false);
            return;
        }

        setLoading(true);
        try {
            const result = await client.get('/api/shorten/handle/' + location.pathname, {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            window.location.href = result.data;
        } catch (error) {
        }
        setLoading(false);
    };
    const fetchData = useCallback(fetchDataRaw, [user]);

    useEffect(() => {
        fetchData();
    }, [user, fetchData, location.pathname]);

    return (
        <Row>
            <Col offset={6} span={12}>
                {loading && <Spin/>}
                {!loading && <div>
                    <Typography.Title>
                        The path <code>{location.pathname}</code> is unknown
                    </Typography.Title>
                    <Typography.Paragraph>
                        The page you're trying to reach does not exist.
                    </Typography.Paragraph>
                </div>}
            </Col>
        </Row>
    );
}
