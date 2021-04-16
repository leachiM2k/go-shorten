import React, {useContext} from 'react';
import {GlobalContext} from '../context/GlobalProvider';
import Login from '../components/Login';
import Logout from '../components/Logout';

const clientId = '***REMOVED***';

export default function CustomRightHeader() {
    const { state, setUser, setLoggedIn } = useContext(GlobalContext);
    const { loggedIn } = state;

    const handleLogin = user => {
        setUser(user);
        setLoggedIn(!!user);
    }

    if (loggedIn) {
        return (
            <Logout clientId={clientId}/>
        );

    } else {
        return (
            <Login clientId={clientId} loggedIn={loggedIn} onLogin={handleLogin}/>
        );

    }

}
