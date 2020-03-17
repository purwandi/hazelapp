import React from "react";
import { Link } from "react-router-dom";

const AuthSigninLayout = props => {
  return (
    <div className="min-h-screen flex items-center justify-center px-4 text-center">
      <div className="max-w-sm w-full ">
        <form method="POST" className="border border-gray-400 shadow px-10 pt-10 pb-5">
          <p className="text-2xl font-medium mb-8">Sign in to your account</p>
          <div className="mb-3">
            <input
              className="py-2 px-3 w-full text-sm rounded border border-solid border-gray-400 focus:border-gray-500"
              type="text"
              name="username"
              placeholder="Please input your username"
            />
          </div>
          <div className="mb-3">
            <input
              className="py-2 px-3 w-full text-sm rounded border border-solid border-gray-400 focus:border-gray-500"
              type="password"
              name="password"
              placeholder="Please input your password"
            />
          </div>
          <div className="mt-5 flex flex-col">
            <button
              to="/"
              className="bg-indigo-500 hover:bg-indigo-700 text-sm text-white py-2 px-12 rounded font-medium tracking-wide"
            >
              Sign in
            </button>
            <div className="text-gray-700 py-6 text-sm">
              <a href="">Forgot Password? </a> - <a href="">Sign up for an account</a>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AuthSigninLayout;
