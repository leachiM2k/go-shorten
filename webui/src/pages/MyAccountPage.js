import React, {useEffect, useState} from "react";
import {Button, Divider, Form, Input, message, Spin, Typography} from 'antd';
import client from '../lib/client-fetch';

export default function MyAccountPage(props) {
    const [formSaving, setFormSaving] = useState(false);
    const [user, setUser] = useState(null);
    const [sshKey, setSshKey] = useState(null);
    const [sshLoading, setSshLoading] = useState(false);
    const [loading, setLoading] = useState(false);
    const [form] = Form.useForm();

    const fetchData = async () => {
        setLoading(true);
        try {
            const result = await client.get('/api/user');
            setUser(result.data);
        } catch (error) {
            if (error.status === 404) {
                setUser({});
            } else {
                message.error('Request failed: ' + error.message);
            }
        }
        setLoading(false);
    }

    const fetchKeyData = async () => {
        setSshLoading(true);
        try {
            const result = await client.get('/api/heimdall/key');
            setSshKey(result.data.key);
        } catch (error) {
            if (error.status === 404) {
                setSshKey("");
            } else {
                message.error('Request failed: ' + error.message);
            }
        }
        setSshLoading(false);
    }

    useEffect(() => { fetchData() }, []);
    useEffect(() => { fetchKeyData() }, []);

    const save = async ({ sshKey, ...values }) => {
        setFormSaving(true);
        try {
            await client.put('/api/user', values);
            await client.post('/api/heimdall/key', { key: sshKey });
            message.success(`User information updated`);
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

    const currentUser = { };

    return (
        <div>
            <Typography.Title>
                My Profile
            </Typography.Title>
            <Divider/>

            {(loading || sshLoading) ? <Spin/> :
                user && sshKey !== null &&
                <Form
                    layout="horizontal"
                    labelCol={{ span: 4 }}
                    wrapperCol={{ span: 14 }}
                    initialValues={{
                        ...user,
                        sshKey
                    }}
                    form={form}
                    onFinish={handleSubmit}
                >
                    <Form.Item label="Personal Number">
                        <span className="ant-form-text">{currentUser.pnum}</span>
                    </Form.Item>

                    <Form.Item label="Name">
                        <span className="ant-form-text">{currentUser.name}</span>
                    </Form.Item>

                    <Form.Item label="Team Name" name="Team">
                        <span className="ant-form-text">{user.Team && user.Team.Name}</span>
                    </Form.Item>

                    <Form.Item label="Slack ID" name="SlackId">
                        <Input/>
                    </Form.Item>

                    <Form.Item label="Github User Name" name="GithubId">
                        <Input/>
                    </Form.Item>

                    <Form.Item label="AWS Username" name="AWSId">
                        <Input/>
                    </Form.Item>

                    <Form.Item label="SSH Public Key" name="sshKey" rules={[
                        {
                            pattern: "^ssh-rsa AAAA[0-9A-Za-z+/]+[=]{0,3}",
                            message: "The key should start with 'ssh-rsa AAAA...'"
                        }
                    ]}>
                        <Input.TextArea autoSize={{ minRows: 2, maxRows: 10 }}
                                        placeholder="ssh-rsa key. Used for heimdall."/>
                    </Form.Item>

                    <Form.Item>
                        <Button type="primary" htmlType="submit" loading={formSaving}>
                            Update
                        </Button>
                    </Form.Item>
                </Form>
            }
        </div>
    );
}
