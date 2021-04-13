import React, {createContext, useReducer} from "react";
import Reducer, {ReducerActions} from './Reducer';
import client from '../lib/client-fetch';
import {message} from 'antd';

const initialState = {
    teamsLoading: false,
    teams: null,
};

export const GlobalContext = createContext(initialState);

const GlobalProvider = ({children}) => {
    const [state, dispatch] = useReducer(Reducer, initialState);

    const apiGetTeamStructure = async () => {
        if (Array.isArray(state.teams) || state.teamsLoading) {
            return state.teams;
        }

        dispatch({
            type: ReducerActions.setTeamsLoading,
            payload: true
        });

        try {
            const result = await client.get('/api/teams');
            dispatch({
                type: ReducerActions.setTeams,
                payload: result.data || []
            });
        } catch (error) {
            message.error('Request failed: ' + error.message);
            dispatch({
                type: ReducerActions.setTeams,
                payload: []
            });

        }

        dispatch({
            type: ReducerActions.setTeamsLoading,
            payload: false
        });
    }

    const apiGetGithubOrgs = async () => {
        if (Array.isArray(state.githubOrgs) || state.githubOrgsLoading) {
            return state.githubOrgs;
        }

        dispatch({
            type: ReducerActions.setGithubOrgsLoading,
            payload: true
        });

        try {
            const result = await client.get('/api/github/orgs');
            dispatch({
                type: ReducerActions.setGithubOrgs,
                payload: result.data.sort() || []
            });
        } catch (error) {
            message.error('Request failed: ' + error.message);
            dispatch({
                type: ReducerActions.setGithubOrgs,
                payload: []
            });
        }

        dispatch({
            type: ReducerActions.setGithubOrgsLoading,
            payload: false
        });
    }

    const setCurrentGithubOrg = org => {
        dispatch({
            type: ReducerActions.setGithubOrgCurrent,
            payload: org
        });
    }

    return (
        <GlobalContext.Provider value={{
            state,
            dispatch,
            apiGetTeamStructure,
            apiGetGithubOrgs,
            setCurrentGithubOrg,
        }}>
            {children}
        </GlobalContext.Provider>
    )
};

export default GlobalProvider;
