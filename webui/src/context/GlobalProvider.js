import React, {createContext, useReducer} from "react";
import Reducer, {ReducerActions} from './Reducer';

const initialState = {
    loggedIn: false,
    user: null,
    loading: false,
    teams: null,
};

export const GlobalContext = createContext(initialState);

const GlobalProvider = ({ children }) => {
    const [state, dispatch] = useReducer(Reducer, initialState);

    const setUser = user => {
        dispatch({
            type: ReducerActions.setUser,
            payload: user
        })
    }

    const setLoggedIn = isLoggedIn => {
        dispatch({
            type: ReducerActions.setLoggedIn,
            payload: isLoggedIn
        })
    }

    return (
        <GlobalContext.Provider value={{
            state,
            dispatch,
            setUser,
            setLoggedIn,
        }}>
            {children}
        </GlobalContext.Provider>
    )
};

export default GlobalProvider;
