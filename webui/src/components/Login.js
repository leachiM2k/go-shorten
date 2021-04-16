import React from 'react';
import {GoogleLogin} from 'react-google-login';
import refreshTokenSetup from '../lib/refresh-token';
import {Button} from 'antd';
import {GoogleCircleFilled} from '@ant-design/icons';

const Login = ({ clientId, loggedIn, onLogin, ...props }) => {
    const onSuccess = res => {
        if (typeof onLogin === 'function') {
            onLogin({ token: res.tokenId, ...res.profileObj });
        }

        // Initializing the setup
        refreshTokenSetup(res);
    };

    const onFailure = res => {
        console.log('***** [Login failed] res:', res);
    };

    const customButton = renderProps => (
        <Button
            icon={<GoogleCircleFilled/>}
            onClick={renderProps.onClick}
            disabled={renderProps.disabled}>
            Login with Google
        </Button>
    )

    return (
        <div>
            <GoogleLogin clientId={clientId}
                         render={customButton}
                         onSuccess={onSuccess}
                         onFailure={onFailure}
                         cookiePolicy={'single_host_origin'}
                         isSignedIn={true}
            />
        </div>
    );
};

export default Login;
