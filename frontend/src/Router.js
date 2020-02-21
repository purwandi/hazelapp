import React from "react";
import { Switch, Route } from "react-router-dom";
import styled from "styled-components";

import AppDashboard from "./pages/dashboard";
import AppProject from "./pages/projects/router";
import AppIssue from "./pages/issues/router";

const AppRoot = styled.div`
  width: 100%;
`;

const AppRouter = () => (
  <AppRoot>
    <Switch>
      <Route path="/" component={AppDashboard} exact />
      <Route path="/projects" component={AppProject} />
      <Route path="/issues" component={AppIssue} />
    </Switch>
  </AppRoot>
);

export default AppRouter;
