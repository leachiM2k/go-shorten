import {Button, Col, List, Popconfirm, Row, Typography} from 'antd';
import {Link} from 'react-router-dom';
import React from 'react';

export const ShortsList = ({ loading, allShorts, onEdit, onDelete }) => {
    const shortenerPrefix = new URL('/', window.location.href).toString();

    const handleEdit = code => event => {
        event.preventDefault();
        onEdit(code);
    };
    const handleDelete = code => () => onDelete(code);

    const listItemActions = code => {
        return (
            [
                <Link to={"/stats/" + code}>Stats</Link>,
                <Link to={"/edit/" + code} onClick={handleEdit(code)}>Edit</Link>,
                <Popconfirm title="Are you sure？" okText="Yes" cancelText="No"
                            onConfirm={handleDelete(code)}>
                    <Button danger size="small">Delete</Button>
                </Popconfirm>,
            ]
        );
    }

    const renderItem = item => {
        return (
            <List.Item actions={listItemActions(item.code)}>
                <Row style={{ flex: '1' }} justify="space-between">
                    <Col>
                        <Typography.Text>{item.description || item.link}</Typography.Text>
                        <br/>
                        <Typography.Link copyable
                                         href={shortenerPrefix + item.code}>{shortenerPrefix}{item.code}</Typography.Link>
                    </Col>
                    <Col>
                        {new Date(item.createdAt).toLocaleDateString()}
                        <br/>
                        {item.count} visits
                    </Col>
                </Row>
            </List.Item>
        );
    }

    return (
        <List
            loading={loading}
            itemLayout="horizontal"
            dataSource={allShorts}
            renderItem={renderItem}
        />
    )
}
