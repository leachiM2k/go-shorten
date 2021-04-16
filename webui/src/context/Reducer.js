export const ReducerActions = {
    setUser: 'SET_USER',
    setTeams: 'SET_TEAMS',
    setTeamsLoading: 'SET_TEAMS_LOADING',
    setGithubOrgs: 'SET_GITHUB_ORGS',
    setGithubOrgsLoading: 'SET_GITHUB_ORGS_LOADING',
    setGithubOrgCurrent: 'SET_GITHUB_ORG_CURRENT',
    setLoggedIn: 'SET_LOGGEDIN',
    setLoading: 'SET_LOADING',
};

const Reducer = (state, action) => {
    switch (action.type) {
        case ReducerActions.setUser:
            return {
                ...state,
                user: action.payload
            };
        case ReducerActions.setLoggedIn:
            return {
                ...state,
                loggedIn: action.payload
            };
        case ReducerActions.setLoading:
            return {
                ...state,
                loading: action.payload
            };
        case ReducerActions.setTeams:
            return {
                ...state,
                teams: action.payload
            };
        case ReducerActions.setTeamsLoading:
            return {
                ...state,
                teamsLoading: action.payload
            };
        case ReducerActions.setGithubOrgs:
            return {
                ...state,
                githubOrgs: action.payload
            };
        case ReducerActions.setGithubOrgsLoading:
            return {
                ...state,
                githubOrgsLoading: action.payload
            };
        case ReducerActions.setGithubOrgCurrent:
            return {
                ...state,
                githubOrgCurrent: action.payload
            };
        default:
            return state;
    }
};

export default Reducer;
