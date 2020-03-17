import React from "react";
import { useSetState } from "react-use";

import Layout from "./layout";

const AuthSignin = props => {
  const [state, setState] = useSetState({
    username: "",
    password: ""
  });

  const onSignIn = e => {
    const form = new URLSearchParams();
    form.set("username", state.username);
    form.set("password", state.password);

    fetch(`${process.env.REACT_APP_BACKEND_URL}/login`, {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/x-www-form-urlencoded",
        "x-requested-with": "XMLHttpRequest"
      },
      body: form
    })
      .then(r => r.json())
      .then(e => console.log(e));

    e.preventDefault();
  };

  return <Layout state={state} setState={setState} onSignIn={onSignIn} {...props} />;
};

export default AuthSignin;
