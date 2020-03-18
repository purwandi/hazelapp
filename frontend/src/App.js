import React from 'react';
import {
  BrowserRouter as Router, Switch, Route, Redirect,
} from 'react-router-dom';

import AppAuth from './pages/auth/router';
import AppContainer from './container';

import { AppProvider, UseAppContextValue } from './context';

const AppRouter = () => {
  const [AppStateContext] = UseAppContextValue();
  return (
    <Router>
      <Switch>
        <Route path="/auth" component={AppAuth} />
        <Route
          path="/"
          render={(props) => {
            if (AppStateContext.user.token === '') {
              return <Redirect to="/auth/signin" />;
            }
            return <AppContainer {...props} />;
          }}
        />
      </Switch>
    </Router>
  );
};

const App = () => (
  <AppProvider>
    <AppRouter />
  </AppProvider>
);

export default App;
