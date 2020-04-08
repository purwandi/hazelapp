import React from 'react';
import { Link } from 'react-router-dom';
import { func } from 'prop-types';

const AuthSignupLayut = ({ onSignUp, setState }) => (
  <div className="min-h-screen flex items-center justify-center py-12 px-4">
    <div className="max-w-sm w-full">
      <h1 className="text-3xl font-bold mb-3">Register new acccount</h1>
      <form onSubmit={onSignUp} method="POST">
        <div className="mb-2">
          <label htmlFor="fullname" className="block text-sm pb-1">
            Fullname
          </label>
          <input
            className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
            type="text"
            id="fullname"
            name="fullname"
            onChange={(e) => setState({ fullname: e.target.value })}
            placeholder="Please enter your fullname"
          />
        </div>
        <div className="mb-2">
          <label htmlFor="username" className="block text-sm pb-1">
            Username
          </label>
          <input
            className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
            type="text"
            id="username"
            name="username"
            onChange={(e) => setState({ username: e.target.value })}
            placeholder="Please pick a username for your login"
          />
        </div>
        <div className="mb-2">
          <label htmlFor="email" className="block text-sm pb-1">
            Email
          </label>
          <input
            className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
            type="text"
            id="email"
            name="email"
            onChange={(e) => setState({ email: e.target.value })}
            placeholder="Please write your valid email"
          />
        </div>
        <div className="mb-2">
          <div className="flex mb-2">
            <div className="w-full pr-5">
              <label htmlFor="password" className="block text-sm pb-1">
                Password
              </label>
              <input
                className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
                type="password"
                id="password"
                name="password"
                onChange={(e) => setState({ password: e.target.value })}
                placeholder="Type your password"
              />
            </div>
            <div className="w-full pl-5">
              <label htmlFor="confirm_password" className="block text-sm pb-1">
                Confirm
              </label>
              <input
                className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
                type="password"
                id="confirm_password"
                name="confirm_password"
                onChange={(e) => setState({ password_confirmation: e.target.value })}
                placeholder="Type your password confirmation"
              />
            </div>
          </div>
          <p className="text-gray-600 text-xs">
            Use 8 or more characters with a mix of letters, numbers & symbols
          </p>
        </div>

        <div className="flex justify-between items-center mt-5">
          <Link
            to="/auth/signin"
            className={`
              bg-indigo-100 hover:bg-indigo-200 text-sm py-2 px-6 text-indigo-600 rounded font-bold inline-block
            `}
          >
            Sign in instead
          </Link>
          <button
            type="submit"
            className="bg-indigo-500 hover:bg-indigo-700 text-sm text-white py-2 px-6 rounded font-bold"
          >
            Next
          </button>
        </div>
      </form>
    </div>
  </div>
);

AuthSignupLayut.propTypes = {
  onSignUp: func.isRequired,
  setState: func.isRequired,
};

export default AuthSignupLayut;
