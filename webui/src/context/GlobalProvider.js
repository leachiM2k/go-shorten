import React, {createContext, useEffect, useReducer} from "react";
import Reducer, {ReducerActions} from './Reducer';
import useStickyState from '../lib/use-sticky-state';
import client from '../lib/client-fetch';
import {message} from 'antd';

const initialState = {
    user: null,
    loading: false,
};

export const GlobalContext = createContext(initialState);

const GlobalProvider = ({ children }) => {
    const [stickyState, setStickyState] = useStickyState({}, 'go-shorten-state');
    const [state, dispatch] = useReducer(Reducer, { ...initialState, ...stickyState });

    useEffect(() => {
        setStickyState(state);
    }, [state, setStickyState]);

    useEffect(() => {
        (async function fetchAuthInfo() {
            if (!state.token) {
                return;
            }
            try {
                const result = await client.get('/auth/', {
                    headers: {
                        'Authorization': 'Bearer ' + state.token,
                    }
                });
                setUser(result.data);
            } catch (error) {
                message.error('Authentication failed. Please login again.');
                setUser(null);
                setToken(null);
            }
        })();
    }, [state.token]);

    const setUser = user => {
        dispatch({
            type: ReducerActions.setUser,
            payload: user
        })
    }

    const setToken = user => {
        dispatch({
            type: ReducerActions.setToken,
            payload: user
        })
    }

    return (
        <GlobalContext.Provider value={{
            state,
            dispatch,
            setUser,
            setToken,
        }}>
            {children}
        </GlobalContext.Provider>
    )
};

export default GlobalProvider;
