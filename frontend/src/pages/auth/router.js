import React from "react";
import { Switch, Route, useRouteMatch } from "react-router-dom";

import AuthSignin from "./signin";
import AuthSignup from "./signup";

const AuthRouter = () => {
  let { path } = useRouteMatch();
  return (
    <Switch>
      <Route path={`${path}/signin`} component={AuthSignin} exact />
      <Route path={`${path}/signup`} component={AuthSignup} exact />
    </Switch>
  );
};

export default AuthRouter;
