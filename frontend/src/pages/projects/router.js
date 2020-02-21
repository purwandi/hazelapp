import React from "react";
import { Switch, Route, useRouteMatch } from "react-router-dom";

import ProjectList from "./project";
import ProjectCreate from "./create";

const ProjectRouter = () => {
  let { path } = useRouteMatch();
  return (
    <Switch>
      <Route path={`${path}/`} component={ProjectList} exact />
      <Route path={`${path}/create`} component={ProjectCreate} exact />
    </Switch>
  );
};

export default ProjectRouter;
