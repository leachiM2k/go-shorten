import React from 'react';
import {GoogleLogout} from 'react-google-login';

const Logout = ({ clientId, ...props }) => {
    const onSuccess = () => {
        console.log('***** [Logout Success] ********');
    };

    const onFailure = res => {
        console.log('***** [Logout failed] res:', res);
    };

    return (
        <div>
            <GoogleLogout
                clientId={clientId}
                buttonText="Logout"
                onLogoutSuccess={onSuccess}
                onFailure={onFailure}
            />
        </div>
    );
};

export default Logout;
