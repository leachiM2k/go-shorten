import React, {useContext} from 'react';
import {GoogleLogout} from 'react-google-login';
import {Button} from 'antd';
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
            icon={<GoogleCircleFilled/>}
            onClick={renderProps.onClick}
            disabled={renderProps.disabled}>
            Logout
        </Button>
    );

    return (
        <div style={{ color: 'white', display: 'inline-flex', alignItems: 'center' }}>
            <div style={{ marginRight: '20px' }}>Logged in as: {user.name}</div>
            <GoogleLogout
                clientId={clientId}
                render={customButton}
                onLogoutSuccess={onSuccess}
                onFailure={onFailure}
            />
        </div>
    );
};

export default Logout;
