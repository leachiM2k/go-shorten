import {Button, Form, Input} from 'antd';
import React from 'react';
import * as PropTypes from 'prop-types';

const ShortForm = props => {

    return (
        <Form
            layout="horizontal"
            labelCol={{ span: 4 }}
            wrapperCol={{ span: 14 }}
            form={props.form}
            initialValues={props.initialValues}
            onFinish={props.onFinish}
        >
            <Form.Item label="Link" name="link">
                <Input/>
            </Form.Item>

            <Form.Item label="Preferred code" name="code">
                <Input/>
            </Form.Item>

            <Form.Item label="Description" name="description">
                <Input/>
            </Form.Item>

            <Form.Item label="Maximum Count of Views" name="maxCount">
                <Input/>
            </Form.Item>

            <Form.Item label="Start time" name="startTime">
                <Input/>
            </Form.Item>

            <Form.Item label="Expiration time" name="expiresAt">
                <Input/>
            </Form.Item>

            <Form.Item label="" wrapperCol={{ offset: 8 }}>
                <Button type="primary" htmlType="submit" loading={props.loading}>
                    Create
                </Button>
            </Form.Item>
        </Form>
    );
};

ShortForm.propTypes = {
    initialValues: PropTypes.any,
    onSave: PropTypes.func,
    loading: PropTypes.bool
};

export default ShortForm;
