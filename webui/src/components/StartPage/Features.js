import React from "react";
import {Card, Col, Row} from 'antd';

const data = [
    {
        title: "Start time",
        content: "Enter the start time from which the links should become valid."
    },
    {
        title: "Expiration time",
        content: "Do you want the links to expire at a certain time? You can set it freely!"
    },
    {
        title: "Maximal visits limit",
        content: "After a defined number of requests, the links become invalid."
    },
    {
        title: "Detailed statistics",
        content: "`Go Shorten!` shows you the IP, the browser version and the time of access to your shortened links."
    },
];

const Features = props => {
    return (
        <Row gutter={16}>
            {data.map(item => (
                <Col xs={24} sm={24} md={12} style={{ display: 'flex' }}>
                    <Card style={{ flex: 1, display: 'flex', flexDirection:'column', marginBottom: '16px' }}
                          headStyle={{ background: '#006d77', color: '#edf6f9' }}
                          bodyStyle={{ background: '#def9fc', flex: 1 }} title={item.title}>
                        {item.content}
                    </Card>
                </Col>
            ))}
        </Row>
    );
}

export default Features;
