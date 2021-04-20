import {Button, Col, DatePicker, Drawer, Form, Input, InputNumber, Row} from 'antd';
import React from 'react';
import moment from 'moment';
import './DrawerForm.css';

const DrawerForm = ({ title, actionText, onSave, onClose, initialValues, visible }) => {
    const [form] = Form.useForm();
    const handleSubmit = () => {
        form
            .validateFields()
            .catch(info => {
                console.log('Validate Failed:', info);
            })
            .then(values => onSave(values));
    }

    if (initialValues) {
        initialValues.startTime = initialValues.startTime && moment(initialValues.startTime);
        initialValues.expiresAt = initialValues.expiresAt && moment(initialValues.expiresAt);
    }

    return (
        <Drawer
            title={title}
            width={""}
            className="mobile-fullscreen-drawer"
            onClose={onClose}
            visible={visible}
            bodyStyle={{ paddingBottom: 80 }}
            footer={
                <div style={{ textAlign: 'right' }}>
                    <Button onClick={onClose} style={{ marginRight: 8 }}>
                        Cancel
                    </Button>
                    <Button onClick={handleSubmit} type="primary">{actionText}</Button>
                </div>
            }
        >
            <Form layout="vertical"
                  form={form}
                  initialValues={initialValues}
            >
                <Row gutter={16}>
                    <Col span={12}>
                        <Form.Item name="createdAt" hidden><Input/></Form.Item>
                        <Form.Item
                            name="link"
                            label="Full URL"
                            rules={[{ required: true, message: 'Please enter url' }]}
                        >
                            <Input
                                style={{ width: '100%' }}
                                addonBefore="http://"
                                placeholder="Please enter url"
                            />
                        </Form.Item>
                    </Col>
                    <Col span={12}>
                        <Form.Item
                            name="code"
                            label="Preferred short code"
                        >
                            <Input placeholder="Short code (e.g. foo12898989)"
                                   disabled={initialValues && initialValues.createdAt}/>
                        </Form.Item>

                    </Col>
                </Row>
                <Row gutter={16}>
                    <Col span={24}>
                        <Form.Item
                            name="description"
                            label="Description"
                            rules={[{ required: true, message: 'Please enter a description' }]}
                        >
                            <Input placeholder="Description of the URL"/>
                        </Form.Item>
                    </Col>
                </Row>
                <Row gutter={16}>
                    <Col span={12}>
                        <Form.Item
                            name="startTime"
                            label="Starts at"
                        >
                            <DatePicker
                                showTime={{ format: 'HH:mm' }}
                                style={{ width: '100%' }}
                                getPopupContainer={trigger => trigger.parentElement}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={12}>
                        <Form.Item
                            name="expiresAt"
                            label="Expires at"
                        >
                            <DatePicker
                                showTime={{ format: 'HH:mm' }}
                                style={{ width: '100%' }}
                                getPopupContainer={trigger => trigger.parentElement}
                            />
                        </Form.Item>
                    </Col>
                </Row>
                <Row gutter={16}>
                    <Col span={12}>
                        <Form.Item
                            name="maxCount"
                            label="Maximum allowed views (0 = unlimited)"
                        >
                            <InputNumber/>
                        </Form.Item>
                    </Col>
                    <Col span={12}>
                    </Col>
                </Row>
            </Form>
        </Drawer>
    );
};

export default DrawerForm;
