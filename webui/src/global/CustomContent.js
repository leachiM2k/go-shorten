import React, {Suspense} from 'react';
import {Route, Switch} from 'react-router-dom';
import routes from '../data/routing';
import RestrictToGroup from './RestrictToGroup';

// A special wrapper for <Route> that knows how to
// handle "sub"-routes by passing them in a `routes`
// prop to the component it renders.
function RouteWithSubRoutes({ component: C, path, routes, restrictToGroups, ...rest }) {
    return (
        <Route path={path}>
            <RestrictToGroup restrictToGroups={restrictToGroups}>
                <Suspense fallback={<div>Loading...</div>}>
                    <C routes={routes}/>
                </Suspense>
            </RestrictToGroup>
        </Route>
    );
}

export default function CustomContent() {
    return (
        <Switch>
            {routes.map((route, i) => (
                <RouteWithSubRoutes key={i} {...route} />
            ))}
        </Switch>
    )
}
