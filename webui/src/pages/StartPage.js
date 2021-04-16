import React, {useContext, useEffect, useState} from "react";
import {GlobalContext} from '../context/GlobalProvider';
import {Button, Col, Form, Input, List, message, Row, Skeleton} from 'antd';
import client from '../lib/client-fetch';

export default function StartPage(props) {
    const { state: { loggedIn, user }, apiAddShort } = useContext(GlobalContext);
    const [formSaving, setFormSaving] = useState(false);
    const [loading, setLoading] = useState(false);
    const [form] = Form.useForm();
    const [allShorts, setAllShorts] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
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
        }
        fetchData();
        }, [user]);

    const save = async ({ sshKey, ...values }) => {
        setFormSaving(true);
        try {
            apiAddShort(values);
            message.success(`information persisted`);
        } catch (error) {
            message.error('Request failed: ' + error.message, 15);
        }
        setFormSaving(false);
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
                <Row>
                    <Col span={12}>
                        {allShorts && <List
                            loading={loading}
                            itemLayout="horizontal"
                            dataSource={allShorts}
                            renderItem={item => (
                                <List.Item
                                    actions={["Foo", "Bar"]}
                                >
                                    <Skeleton avatar title={false} loading={item.loading} active>
                                        <List.Item.Meta
                                            title={<a href={item.link}>{item.link}</a>}
                                            description={item.description}
                                        />
                                        <div>content</div>
                                    </Skeleton>
                                </List.Item>
                            )}
                        />}
                    </Col>
                    <Col span={12}>

                        <Form
                            layout="horizontal"
                            labelCol={{ span: 4 }}
                            wrapperCol={{ span: 14 }}
                            form={form}
                            onFinish={handleSubmit}
                        >
                            <Form.Item label="link" name="link">
                                <Input/>
                            </Form.Item>

                            <Form.Item label="preferred code" name="code">
                                <Input/>
                            </Form.Item>

                            <Form.Item label="description" name="description">
                                <Input/>
                            </Form.Item>

                            <Form.Item label="Maximum Count of Views" name="maxcount">
                                <Input/>
                            </Form.Item>

                            <Form.Item label="Start time" name="startTime">
                                <Input/>
                            </Form.Item>

                            <Form.Item label="Expiration time" name="expiresAt">
                                <Input/>
                            </Form.Item>

                            <Form.Item label="" wrapperCol={{ offset: 8 }}>
                                <Button type="primary" htmlType="submit" loading={formSaving}>
                                    Create
                                </Button>
                            </Form.Item>
                        </Form>

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
