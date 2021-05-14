import React, {useCallback, useContext, useEffect, useState} from "react";
import {useLocation} from "react-router-dom";
import {Col, Row, Spin, Typography} from 'antd';
import client from '../lib/client-fetch';
import {GlobalContext} from '../context/GlobalProvider';

export function NotFound() {
    const location = useLocation();
    const { state: { user, token } } = useContext(GlobalContext);
    const [loading, setLoading] = useState(true);

    const fetchDataRaw = async () => {
        if (!token) {
            setLoading(false);
            return;
        }

        setLoading(true);
        try {
            const result = await client.get('/api/shorten/handle/' + location.pathname, {
                headers: {
                    'Authorization': 'Bearer ' + token,
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
            <Col span={24} md={{ span: 18, offset: 3 }} lg={{ span: 12, offset: 6 }}>
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
