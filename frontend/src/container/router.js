import React from "react";
import { Switch, Route } from "react-router-dom";

import AppDiscussion from "../pages/discussion";
import AppProject from "../pages/projects/router";
import AppIssues from "../pages/issues/router";

const AppContainerRouter = () => {
  return (
    <Switch>
      <Route path="/project/:id/discussion" component={AppDiscussion} exact />
      <Route path="/project/:id/issues" component={AppIssues} />
      <Route path="/project" component={AppProject} />
    </Switch>
  );
};

export default AppContainerRouter;
