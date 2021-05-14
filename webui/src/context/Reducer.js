export const ReducerActions = {
    setUser: 'SET_USER',
    setToken: 'SET_TOKEN',
    setLoading: 'SET_LOADING',
};

const Reducer = (state, action) => {
    switch (action.type) {
        case ReducerActions.setUser:
            return {
                ...state,
                user: action.payload
            };
        case ReducerActions.setToken:
            return {
                ...state,
                token: action.payload
            };
        case ReducerActions.setLoading:
            return {
                ...state,
                loading: action.payload
            };
        default:
            return state;
    }
};

export default Reducer;
