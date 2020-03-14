import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import AppAuth from "./pages/auth/router";
import AppContainer from "./container";

import { AppProvider } from "./context";

const App = () => (
  <AppProvider>
    <Router>
      <Switch>
        <Route path="/auth" component={AppAuth} />
        <Route path="/" component={AppContainer} />
      </Switch>
    </Router>
  </AppProvider>
);

export default App;
