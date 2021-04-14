import {NotFound} from '../global/NotFound';
import {lazy} from 'react';

const StartPage = lazy(() => import('../pages/StartPage'));

const routes = [
    { path: '/', component: StartPage },
    { path: '*', component: NotFound },
];

export default routes;
