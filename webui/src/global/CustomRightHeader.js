import React, {useContext} from 'react';
import {GlobalContext} from '../context/GlobalProvider';
import Login from '../components/Login';
import Logout from '../components/Logout';

const clientId = '***REMOVED***';

export default function CustomRightHeader() {
    const { state, setUser, setLoggedIn } = useContext(GlobalContext);
    const { loggedIn, user } = state;

    const handleLogin = user => {
        setUser(user);
        setLoggedIn(!!user);
    }

    if (loggedIn) {
        return (
            <div style={{ color: 'white', display: 'inline-flex' }}>
                <div>Logged in as: {user.name}</div>
                <Logout clientId={clientId}/>
            </div>
        );

    } else {
        return (
            <Login clientId={clientId} loggedIn={loggedIn} onLogin={handleLogin}/>
        );

    }

}
