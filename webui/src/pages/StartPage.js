import React, {useCallback, useContext, useEffect, useState} from "react";
import {GlobalContext} from '../context/GlobalProvider';
import {Button, Col, Form, List, message, Row, Typography} from 'antd';
import client from '../lib/client-fetch';
import ShortForm from '../components/ShortForm';

export default function StartPage(props) {
    const { state: { loggedIn, user } } = useContext(GlobalContext);
    const [formSaving, setFormSaving] = useState(false);
    const [loading, setLoading] = useState(false);
    const [form] = Form.useForm();
    const [allShorts, setAllShorts] = useState(null);
    const [editValues, setEditValues] = useState(null);

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
            await client.post('/api/shorten/', values, {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            message.success(`information persisted`);
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

    const handleEdit = code => async () => {
        try {
            const result = await client.get('/api/shorten/' + code, {
                headers: {
                    'Authorization': 'Bearer ' + user.token,
                }
            });
            setEditValues(result.data);
            form.resetFields();
        } catch (error) {
            if (error.status === 404) {
                setEditValues({});
            } else {
                message.error('Request failed: ' + error.message);
            }
        }
    }

    const handleSubmit = () => {
        form
            .validateFields()
            .catch(info => {
                console.log('Validate Failed:', info);
            })
            .then(values => save(values));
    }

    if (loggedIn) {
        return (
            <div>
                <Row gutter={10}>
                    <Col span={12}>
                        {allShorts && <List
                            loading={loading}
                            itemLayout="horizontal"
                            dataSource={allShorts}
                            renderItem={item => (
                                <List.Item actions={[
                                    <Button danger size="small" onClick={handleDelete(item.code)}>Delete</Button>,
                                    <Button size="small" onClick={handleEdit(item.code)}>Edit</Button>,
                                ]}>
                                    {item.createdAt}
                                    <br/>
                                    <Typography.Text>{item.description || item.link}</Typography.Text>
                                    <br/>
                                    <Typography.Link
                                        href={"https://shortener/" + item.code}>https://shortener/{item.code}</Typography.Link>
                                </List.Item>
                            )}
                        />}
                    </Col>
                    <Col span={12}>
                        <ShortForm form={form} initialValues={editValues} onFinish={handleSubmit} loading={formSaving}/>
                    </Col>
                </Row>

            </div>
        );
    } else {
        return (
            <div>
                Please login in the right upper corner
            </div>
        );
    }
}
