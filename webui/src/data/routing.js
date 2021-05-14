import {NotFound} from '../global/NotFound';
import {lazy} from 'react';

const StartPage = lazy(() => import('../pages/StartPage'));
const StatisticsPage = lazy(() => import('../pages/StatisticsPage'));
const LoginPage = lazy(() => import('../pages/LoginPage'));

const routes = [
    { path: '/', exact: true, component: StartPage },
    { path: '/login', exact: true, component: LoginPage },
    { path: '/stats/:code', exact: true, component: StatisticsPage },
    { path: '*', component: NotFound },
];

export default routes;
