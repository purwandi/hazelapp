import React from 'react';
import { Switch, Route, useRouteMatch } from 'react-router-dom';

import ProjectList from './index';
import ProjectCreate from './create';

const ProjectRouter = () => {
  const { path } = useRouteMatch();
  return (
    <Switch>
      <Route path={`${path}/`} component={ProjectList} exact />
      <Route path={`${path}/create`} component={ProjectCreate} exact />
    </Switch>
  );
};

export default ProjectRouter;
