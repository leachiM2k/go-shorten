import {Button} from 'antd';
import React, {useContext} from 'react';
import {GlobalContext} from '../context/GlobalProvider';

const Authentication = props => {
    const { state: { user }, setUser, setToken } = useContext(GlobalContext);

    const handleLogout = () => {
        if (user) {
            setUser(null);
            setToken(null);
        }
    }

    if (!user) {
        return null;
    }

    return (
        <div> Logged in via {user.p} as: {user.name}
            <Button onClick={handleLogout}>logout</Button>
        </div>
    );
};

export default Authentication;
