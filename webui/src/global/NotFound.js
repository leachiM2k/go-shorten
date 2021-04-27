import React from "react";
import {useLocation} from "react-router-dom";
import {Col, Row, Typography} from 'antd';

export function NotFound() {
    const location = useLocation();

    return (
        <Row>
            <Col offset={6} span={12}>
                <div>
                    <Typography.Title>
                        The path <code>{location.pathname}</code> is unknown
                    </Typography.Title>
                    <Typography.Paragraph>
                        The page you're trying to reach does not exist.
                    </Typography.Paragraph>
                </div>
            </Col>
        </Row>
    );
}
