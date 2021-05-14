import {FacebookFilled, GithubFilled, GoogleCircleFilled, SlackCircleFilled} from '@ant-design/icons';
import {Button, Col, Row} from 'antd';
import React, {useContext} from 'react';
import {GlobalContext} from '../context/GlobalProvider';

const Login = props => {
    const { state: { user } } = useContext(GlobalContext);

    if (user) {
        return null;
    }

    return (
        <Row justify="space-around">
            <Col>
                <Button
                    icon={<FacebookFilled/>}
                    onClick={() => window.location.href = '/auth/facebook/'}
                >
                    Login with Facebook
                </Button>
            </Col>
            <Col>
                <Button
                    icon={<GoogleCircleFilled/>}
                    onClick={() => window.location.href = '/auth/google/'}
                >
                    Login with Google
                </Button>
            </Col>
            <Col>
                <Button
                    icon={<GithubFilled/>}
                    onClick={() => window.location.href = '/auth/github/'}
                >
                    Login with GitHub
                </Button>
            </Col>
            <Col>
                <Button
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
