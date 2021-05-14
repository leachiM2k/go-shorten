import SocialButton from './SocialButton';
import {FacebookFilled, GoogleCircleFilled} from '@ant-design/icons';
import {Button, Col, Row} from 'antd';
import React, {useContext} from 'react';
import {GlobalContext} from '../context/GlobalProvider';

const Authentication = props => {
    const { state, setUser, setToken } = useContext(GlobalContext);

    const handleLogout = () => {
        if (state.user) {
            setUser(null);
            setToken(null);
        }
    }

    return (
        <Row justify="space-around">
            <Col>
                <SocialButton
                    icon={<FacebookFilled/>}
                    isLoggedIn={!!state.user}
                    triggerLogin={() => window.location.href='/auth/facebook/'}
                >
                    Login with Facebook
                </SocialButton>
            </Col>
            <Col>
                <SocialButton
                    icon={<GoogleCircleFilled/>}
                    isLoggedIn={!!state.user}
                    triggerLogin={() => window.location.href='/auth/google/'}
                >
                    Login with Google
                </SocialButton>
            </Col>

            {state.user &&
            <div> Logged in via {state.user.p} as: {state.user.name}
                <Button onClick={handleLogout}>logout</Button>
            </div>}

        </Row>
    );
};

export default Authentication;
