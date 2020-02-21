import React from "react";
import { Switch, Route, useRouteMatch } from "react-router-dom";

import IssueList from "./issues";
import IssueCreate from "./create";

const IssueRouter = () => {
  const { path } = useRouteMatch();

  return (
    <Switch>
      <Route path={`${path}/`} exact component={IssueList} />
      <Route path={`${path}/create`} exact component={IssueCreate} />
    </Switch>
  );
};

export default IssueRouter;
