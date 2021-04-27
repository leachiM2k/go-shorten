import React, {useCallback, useContext, useEffect, useState} from "react";
import {GlobalContext} from '../context/GlobalProvider';
import {Button, Col, List, message, Popconfirm, Row, Typography} from 'antd';
import client from '../lib/client-fetch';
import {PlusOutlined} from '@ant-design/icons';
import DrawerForm from '../components/DrawerForm';
import {Link} from 'react-router-dom';

export default function StartPage(props) {
    const { state: { loggedIn, user } } = useContext(GlobalContext);
    const [drawerMode, setDrawerMode] = useState(null);
    const [formSaving, setFormSaving] = useState(false);
    const [loading, setLoading] = useState(false);
    const [allShorts, setAllShorts] = useState(null);
    const [editValues, setEditValues] = useState(null);

    const shortenerPrefix = window.location.host + '/';

    const fetchDataRaw = async () => {
        if (user === null || !user.token) {
            return;
        }

        setLoading(true);
        try {
            const result = await client.get('/api/shorten/', {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            setAllShorts(result.data);
        } catch (error) {
            if (error.status === 404) {
                setAllShorts({});
            } else {
                message.error('Request failed: ' + error.message);
            }
        }
        setLoading(false);
    };
    const fetchData = useCallback(fetchDataRaw, [user]);

    useEffect(() => {
        fetchData();
    }, [user, fetchData]);

    const save = async (values) => {
        setFormSaving(true);
        try {
            const options = {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            }
            if (values.createdAt) {
                await client.put('/api/shorten/' + values.code, values, options);
            } else {
                await client.post('/api/shorten/', values, options);
            }
            message.success(`information persisted`);
            setDrawerMode(null);
            await fetchData();
        } catch (error) {
            message.error('Request failed: ' + error.message, 15);
        }
        setFormSaving(false);
    }

    const handleDelete = code => async () => {
        try {
            await client.delete('/api/shorten/' + code, {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            fetchDataRaw()
        } catch (error) {
            message.error('Request failed: ' + error.message);
        }

    }

    const handleEdit = code => async (event) => {
        event.preventDefault();
        try {
            const result = await client.get('/api/shorten/' + code, {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            setEditValues(result.data);
            setDrawerMode('edit');
        } catch (error) {
            if (error.status === 404) {
                setEditValues({});
            } else {
                message.error('Request failed: ' + error.message);
            }
        }
    }

    const showDrawer = () => {
        setDrawerMode('create');
    };

    const onCloseDrawer = () => {
        setEditValues(null);
        setDrawerMode(null);
    };

    if (loggedIn) {
        return (
            <Row>
                <Col offset={6} span={12}>
                    <div style={{ textAlign: 'right', margin: '1em 0' }}>
                        <Button type="primary" onClick={showDrawer}>
                            <PlusOutlined/> Create
                        </Button>
                    </div>

                    {drawerMode !== null &&
                    <DrawerForm title={drawerMode === "create" ? "Create a new short" : "Update a short"}
                                actionText={drawerMode === "create" ? "Shorten!" : "Update"}
                                onSave={save}
                                onClose={onCloseDrawer}
                                loading={formSaving}
                                initialValues={editValues}
                                visible={true}/>}

                    {allShorts && <List
                        loading={loading}
                        itemLayout="horizontal"
                        dataSource={allShorts}
                        renderItem={item => (
                            <List.Item actions={[
                                <Link to={"/stats/" + item.code}>Stats</Link>,
                                <Link to={"/edit/" + item.code} onClick={handleEdit(item.code)}>Edit</Link>,
                                <Popconfirm title="Are you sure？" okText="Yes" cancelText="No"
                                            onConfirm={handleDelete(item.code)}>
                                    <Button danger size="small">Delete</Button>
                                </Popconfirm>,
                            ]}>
                                <Row style={{ flex: '1' }} justify="space-between">
                                    <Col>
                                        <Typography.Text>{item.description || item.link}</Typography.Text>
                                        <br/>
                                        <Typography.Link
                                            href={shortenerPrefix + item.code}>https://{shortenerPrefix}{item.code}</Typography.Link>
                                    </Col>
                                    <Col>
                                        {new Date(item.createdAt).toLocaleDateString()}
                                        <br/>
                                        {item.count} visits
                                    </Col>
                                </Row>
                            </List.Item>
                        )}
                    />}
                </Col>
            </Row>
        );
    } else {
        return (
            <Row>
                <Col offset={6} span={12}>
                    <p>Please login in the right upper corner</p>
                </Col>
            </Row>
        );
    }
}
