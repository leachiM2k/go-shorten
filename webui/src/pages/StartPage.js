import React, {lazy, useCallback, useContext, useEffect, useState} from "react";
import {GlobalContext} from '../context/GlobalProvider';
import {Button, Card, Col, Form, message, Row, Typography} from 'antd';
import client from '../lib/client-fetch';
import {PlusOutlined} from '@ant-design/icons';
import {ShortsList} from '../components/ShortsList';

const DrawerForm = lazy(() => import('../components/DrawerForm'))

export default function StartPage(props) {
    const { state: { loggedIn, user } } = useContext(GlobalContext);
    const [drawerMode, setDrawerMode] = useState(null);
    const [formSaving, setFormSaving] = useState(false);
    const [loading, setLoading] = useState(false);
    const [allShorts, setAllShorts] = useState(null);
    const [editValues, setEditValues] = useState(null);
    const [form] = Form.useForm();

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
        if (!values) {
            return;
        }
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

    const handleDelete = async (code) => {
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

    const handleEdit = async (code) => {
        form.resetFields();
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
        form.resetFields();
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

                    <DrawerForm title={drawerMode === "create" ? "Create a new short" : "Update a short"}
                                actionText={drawerMode === "create" ? "Shorten!" : "Update"}
                                onSave={save}
                                form={form}
                                onClose={onCloseDrawer}
                                loading={formSaving}
                                initialValues={drawerMode === "create" ? {} : editValues}
                                visible={drawerMode !== null}/>

                    {allShorts &&
                    <ShortsList loading={loading} allShorts={allShorts} onDelete={handleDelete} onEdit={handleEdit}/>}
                </Col>
            </Row>
        );
    } else {
        return (
            <Row>
                <Col offset={6} span={12}>
                    <Card title="Login required">
                        <Typography.Paragraph>Please login in the right upper corner</Typography.Paragraph>
                    </Card>
                </Col>
            </Row>
        );
    }
}
