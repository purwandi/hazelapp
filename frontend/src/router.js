import React from 'react'
import { BrowserRouter as Router, Switch, Route, Redirect, withRouter, useRouteMatch } from 'react-router-dom'
import { ApolloProvider } from 'react-apollo'

// Context Component
import { AppProvider, UseAppContextValue } from './context'
import { CreateApolloClient } from './graphql'

// Component Import
import Navigation from './components/navigation'

// Component Router Import
import AppAuth from './pages/auth/router'
import { WorkspaceIndex, WorkspaceCreate } from './pages/workspace'
import { BacklogIndex } from './pages/backlog'

const App = () => (
  <AppProvider>
    <Router>
      <AppRouter />
    </Router>
  </AppProvider>
)

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
            return <AppMainContainer {...props} />;
          }}
        />
      </Switch>
    </ApolloProvider>
  );
});

const AppMainContainer = props => {
  return (
    <div className="h-screen flex">
      <div className="bg-gray-100 w-56 border-r border-gray-400 flex flex-col">
        <Navigation {...props} />
      </div>
      <div className="flex-1 min-w-0 bg-white">
        <Switch>
          <Route path="/:login/:project" component={BacklogIndex} />
          <Route path="/:login/:project/backlog" component={BacklogIndex} />
          <Route path="/create" component={WorkspaceCreate} exact />
          <Route path="/" component={WorkspaceIndex} exact />
        </Switch>
      </div>
    </div>
  );
};

export default App;
