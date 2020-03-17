import React from "react";
import { useSetState } from "react-use";

import { UseAppContextValue } from "../../../context";
import Layout from "./layout";

const AuthSignin = props => {
  const [state, setState] = useSetState({
    username: "",
    password: ""
  });
  const [, dispatch] = UseAppContextValue();

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
      .then(r => {
        // set to react context
        const data = {
          username: r.username,
          fullname: r.fullname,
          email: r.email,
          token: r.accessToken
        };
        dispatch({ type: "AUTHENTICATE", data });

        // redirect after success login
        props.history.push("/");
      });

    e.preventDefault();
  };

  return <Layout state={state} setState={setState} onSignIn={onSignIn} {...props} />;
};

export default AuthSignin;
