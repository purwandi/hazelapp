import React from 'react';
import { useSetState } from 'react-use';
import Layout from './layout';

const AuthSignup = (props) => {
  const [state, setState] = useSetState({
    fullname: '',
    username: '',
    email: '',
    password: '',
    password_confirmation: '',
  });

  const onSignUp = (e) => {
    const form = new URLSearchParams();
    form.set('fullname', state.fullname);
    form.set('email', state.email);
    form.set('username', state.username);
    form.set('password', state.password);
    form.set('password_confirmation', state.password_confirmation);

    fetch(`${process.env.REACT_APP_BACKEND_URL}/register`, {
      method: 'POST',
      mode: 'cors',
      cache: 'no-cache',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded',
        'x-requested-with': 'XMLHttpRequest',
      },
      body: form,
    })
      .then((response) => {
        if (!response.ok) throw response;
        return response.json();
      })
      .then(() => props.history.push('/auth/signin'))
      // eslint-disable-next-line no-console
      .catch(() => e.json().then(console.log));

    e.preventDefault();
  };

  return <Layout onSignUp={onSignUp} state={state} setState={setState} {...props} />;
};

export default AuthSignup;
