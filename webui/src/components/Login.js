import React from 'react';
import {GoogleLogin} from 'react-google-login';
import refreshTokenSetup from '../lib/refresh-token';

const Login = ({ clientId, loggedIn, onLogin, ...props }) => {
    const onSuccess = res => {
        console.log('*************************** res', res);
        console.log('***** [Login Success] currentUser:', res.profileObj);

        if(typeof onLogin === 'function') {
            onLogin(res.profileObj);
        }

        // Initializing the setup
        refreshTokenSetup(res);
    };

    const onFailure = res => {
        console.log('***** [Login failed] res:', res);
    };

    return (
        <div>
            <GoogleLogin clientId={clientId}
                         buttonText="Login"
                         onSuccess={onSuccess}
                         onFailure={onFailure}
                         cookiePolicy={'single_host_origin'}
                         style={{ marginTop: '100px' }}
                         isSignedIn={true}
            />
        </div>
    );
};

export default Login;
