import React, {useContext} from 'react';
import {GoogleLogout} from 'react-google-login';
import {Button, Col, Row} from 'antd';
import {GoogleCircleFilled} from '@ant-design/icons';
import {GlobalContext} from '../context/GlobalProvider';

const Logout = ({ clientId, ...props }) => {
    const { state } = useContext(GlobalContext);
    const { user } = state;

    const onSuccess = () => {
        console.log('***** [Logout Success] ********');
    };

    const onFailure = res => {
        console.log('***** [Logout failed] res:', res);
    };

    const customButton = renderProps => (
        <Button
            type="link"
            icon={<GoogleCircleFilled/>}
            onClick={renderProps.onClick}
            disabled={renderProps.disabled}>
            Logout
        </Button>
    );

    return (
        <Row align="middle">
            <Col>Logged in as: {user.name}</Col>
            <Col>
                <GoogleLogout
                    clientId={clientId}
                    render={customButton}
                    onLogoutSuccess={onSuccess}
                    onFailure={onFailure}
                />
            </Col>
        </Row>
    );
};

export default Logout;
