import {NotFound} from '../global/NotFound';
import {lazy} from 'react';

const StartPage = lazy(() => import('../pages/StartPage'));
const StatisticsPage = lazy(() => import('../pages/StatisticsPage'));

const routes = [
    { path: '/', exact: true, component: StartPage },
    { path: '/stats/:code', exact: true, component: StatisticsPage },
    { path: '*', component: NotFound },
];

export default routes;
