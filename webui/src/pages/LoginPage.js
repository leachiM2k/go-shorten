import React, {useContext} from "react";
import {GlobalContext} from '../context/GlobalProvider';
import {Redirect} from 'react-router-dom';

export default function LoginPage(props) {
    const {  setToken } = useContext(GlobalContext);

    const jwtToken = new URL(window.location.href).searchParams.get('jwt');
    setToken(jwtToken);

    return (
        <Redirect to="/" />
    )

}
