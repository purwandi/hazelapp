import React, { useReducer, useEffect, useContext } from 'react';
import { node } from 'prop-types';
import store from 'store';

export const LOCAL_STORAGE_KEY = 'authorization';
const dataStore = {
  user: {
    username: '',
    fullname: '',
    email: '',
    token: '',
  },
  loading: false,
};

const AppContext = React.createContext();
const useAppReducer = (reducer, defaultState) => {
  const hookVars = useReducer(reducer, defaultState, initState => store.get(LOCAL_STORAGE_KEY) || initState);

  useEffect(() => {
    store.set(LOCAL_STORAGE_KEY, hookVars[0]);
  });

  return hookVars;
};

const Reducer = (state, action) => {
  switch (action.type) {
    case 'AUTHENTICATE':
      return {
        ...state,
        user: {
          username: action.data.username,
          fullname: action.data.fullname,
          email: action.data.email,
          token: action.data.token,
        },
      };
    case 'LOGOUT':
      return {
        ...state,
        user: {
          username: '',
          fullname: '',
          email: '',
          token: '',
        },
      };
    default:
      break;
  }
};

export const AppProvider = props => {
  const AppStateContext = dataStore;
  return <AppContext.Provider value={useAppReducer(Reducer, AppStateContext)}>{props.children}</AppContext.Provider>;
};

AppProvider.propTypes = {
  children: node.isRequired,
};

export const UseAppContextValue = () => useContext(AppContext);
