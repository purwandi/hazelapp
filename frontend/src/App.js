import React from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect, withRouter } from 'react-router-dom';
import { ApolloProvider } from 'react-apollo';

import AppAuth from './pages/auth/router';
import AppContainer from './container';

import { AppProvider, UseAppContextValue } from './context';
import { CreateApolloClient } from './graphql';

const AppRouter = withRouter(props => {
  const { match, history, location } = props;
  const [AppStateContext, dispatch] = UseAppContextValue();
  const ApolloClient = CreateApolloClient(AppStateContext, dispatch, { match, history, location });
  return (
    <ApolloProvider client={ApolloClient}>
      <Switch>
        <Route path="/auth" component={AppAuth} />
        <Route
          path="/"
          render={props => {
            if (AppStateContext.user.token === '') {
              return <Redirect to="/auth/signin" />;
            }
            return <AppContainer {...props} />;
          }}
        />
      </Switch>
    </ApolloProvider>
  );
});

const App = () => (
  <AppProvider>
    <Router>
      <AppRouter />
    </Router>
  </AppProvider>
);

export default App;
