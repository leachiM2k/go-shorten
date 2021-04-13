import React from "react";
import {useLocation} from "react-router-dom";

export function NotFound() {
    const location = useLocation();

    return (
        <div>
            <h3>
                The path <code>{location.pathname}</code> has no route
            </h3>
            <p>
                The page you're trying to reach does not exist.
            </p>
        </div>
    );
}
