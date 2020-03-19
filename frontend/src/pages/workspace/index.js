import React from "react";
import { Switch, Route, useRouteMatch } from "react-router-dom";

import WorkspaceCreate from "./create";
import WorkspaceIndex from "./board";

const WorkspaceRouter = () => {
  const { path } = useRouteMatch();
  console.log(path);
  return (
    <Switch>
      <Route path={`${path}create`} exact component={WorkspaceCreate} />
      <Route path={`${path}`} exact component={WorkspaceIndex} />
    </Switch>
  );
};

export default WorkspaceRouter;
