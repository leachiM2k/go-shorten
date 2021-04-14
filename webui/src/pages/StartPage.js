import React, {useContext} from "react";
import {GlobalContext} from '../context/GlobalProvider';

export default function StartPage(props) {
    const { state: { loggedIn } } = useContext(GlobalContext);

    if(loggedIn) {
        return (
            <div>
                Feel free to have a look. But: <strong>Don't try things, if you don't know exaclty what they are for.</strong>
            </div>
        );
    } else {
        return (
            <div>
                Please login in the right upper corner
            </div>
        );
    }
}
