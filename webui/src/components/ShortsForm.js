import {Col, DatePicker, Form, Input, InputNumber, Row} from 'antd';
import React from 'react';
import moment from 'moment';
import './DrawerForm.css';
import {GlobalOutlined} from '@ant-design/icons';

const DrawerForm = ({ form, initialValues }) => {
    if (initialValues) {
        initialValues.startTime = initialValues.startTime && moment(initialValues.startTime);
        initialValues.expiresAt = initialValues.expiresAt && moment(initialValues.expiresAt);
    }

    return (
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
                        rules={[
                            { required: true, message: 'Please enter url' },
                            { pattern: "^http(s)?://", message: 'Should start with http:// or https://' },
                        ]}
                    >
                        <Input
                            style={{ width: '100%' }}
                            addonBefore={<GlobalOutlined />}
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
    );
};

export default DrawerForm;
