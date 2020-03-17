import React from "react";
import { Switch, Route, useRouteMatch } from "react-router-dom";

import AppDiscussion from "../pages/discussion";
import AppProject from "../pages/projects/router";
import AppIssues from "../pages/issues/router";

const AppContainerRouter = props => {
  const { path } = useRouteMatch();
  return (
    <Switch>
      <Route path={`${path}project/:id/discussion`} component={AppDiscussion} />
      <Route path={`${path}project/:id/issues`} component={AppIssues} />
      <Route path={`${path}project`} component={AppProject} />
    </Switch>
  );
};

export default AppContainerRouter;
