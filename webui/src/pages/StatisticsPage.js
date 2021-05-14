import React, {useCallback, useContext, useEffect, useState} from 'react';
import {Card, Col, message, Row, Table} from 'antd';
import {useParams} from 'react-router-dom';
import {GlobalContext} from '../context/GlobalProvider';
import client from '../lib/client-fetch';

const columns = [
    {
        title: 'IP',
        dataIndex: 'clientIP',
        key: 'clientIP',
        sorter: true,
    },
    {
        title: 'User-Agent',
        dataIndex: 'userAgent',
        key: 'userAgent',
        responsive: ['md'],
        ellipsis: true,
        sorter: true,
    },
    {
        title: 'Date Time',
        dataIndex: 'timestamp',
        key: 'timestamp',
        render: value => new Date(value).toLocaleString(),
        responsive: ['md'],
        sorter: true,
    },
];

export default function StatisticsPage(props) {
    const { state: { user, token } } = useContext(GlobalContext);
    let { code } = useParams();
    const [loading, setLoading] = useState(false);
    const [shortStats, setShortStats] = useState(null);

    const fetchDataRaw = async () => {
        if (!token) {
            return;
        }

        setLoading(true);
        try {
            const result = await client.get('/api/shorten/' + code + "/stats", {
                headers: {
                    'Authorization': 'Bearer ' + token,
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

    function renderData() {
        return (
            <Table
                columns={columns}
                loading={loading}
                dataSource={shortStats}/>
        )
    }

    return (
        <Row>
            <Col span={24} md={{ span: 18, offset: 3 }} lg={{ span: 12, offset: 6 }}>
                <Card title={"Access Logs for \"" + code + "\""}>
                    {renderData()}
                </Card>
            </Col>
        </Row>
    );
}
