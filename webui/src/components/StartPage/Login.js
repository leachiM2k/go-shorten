import {FacebookFilled, GithubFilled, GoogleCircleFilled, SlackCircleFilled} from '@ant-design/icons';
import {Button, Col, Row} from 'antd';
import React, {useContext} from 'react';
import {GlobalContext} from '../../context/GlobalProvider';

const Login = props => {
    const { state: { user } } = useContext(GlobalContext);

    if (user) {
        return null;
    }

    return (
        <Row gutter={[16,16]} justify="space-around">
            <Col xs={24} sm={12}>
                <Button
                    style={{width:'100%'}}
                    icon={<FacebookFilled/>}
                    onClick={() => window.location.href = '/auth/facebook/'}
                >
                    Login with Facebook
                </Button>
            </Col>
            <Col xs={24} sm={12}>
                <Button
                    style={{width:'100%'}}
                    icon={<GoogleCircleFilled/>}
                    onClick={() => window.location.href = '/auth/google/'}
                >
                    Login with Google
                </Button>
            </Col>
            <Col xs={24} sm={12}>
                <Button
                    style={{width:'100%'}}
                    icon={<GithubFilled/>}
                    onClick={() => window.location.href = '/auth/github/'}
                >
                    Login with GitHub
                </Button>
            </Col>
            <Col xs={24} sm={12}>
                <Button
                    style={{width:'100%'}}
                    icon={<SlackCircleFilled/>}
                    onClick={() => window.location.href = '/auth/slack/'}
                >
                    Login with Slack
                </Button>
            </Col>
        </Row>
    );
};

export default Login;
