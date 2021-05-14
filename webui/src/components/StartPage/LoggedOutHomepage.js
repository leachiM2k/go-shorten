import {Card, Col, Row, Typography} from 'antd';
import Features from './Features';
import React from 'react';
import Authentication from '../Authentication';

const LoggedOutHomepage = props => {
    return (
        <Row>
            <Col span={24} md={{ span: 18, offset: 3 }} lg={{ span: 12, offset: 6 }}>
                <Card>
                    <Typography.Title>Here is your new URL shortener</Typography.Title>
                    <Typography.Paragraph>No registration needed. To start using it, please login:</Typography.Paragraph>
                    <Authentication/>
                    <Typography.Paragraph>This free URL shortener offers following features:</Typography.Paragraph>
                    <Features/>
                </Card>
            </Col>
        </Row>
    );
}

export default LoggedOutHomepage;
