import {NotFound} from '../global/NotFound';
import {lazy} from 'react';

const TeamsPage = lazy(() => import('../pages/TeamsPage'));
const HeimdallPage = lazy(() => import('../pages/HeimdallPage'));
const MyAccountPage = lazy(() => import('../pages/MyAccountPage'));
const StartPage = lazy(() => import('../pages/StartPage'));

const routes = [
    { path: '/teams', component: TeamsPage, restrictToGroups: ['admin'] },
    { path: '/heimdall', component: HeimdallPage, restrictToGroups: ['role_agile_software_developer'] },
    { path: '/me', component: MyAccountPage, restrictToGroups: [] },
    { path: '/', component: StartPage },
    { path: '*', component: NotFound },
];

export default routes;
