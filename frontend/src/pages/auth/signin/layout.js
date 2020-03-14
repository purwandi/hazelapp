import React from "react";
import { Link } from "react-router-dom";

const AuthSigninLayout = props => {
  return (
    <div className="min-h-screen flex items-center justify-center py-12 px-4">
      <div className="max-w-sm w-full">
        <h1 className="text-3xl font-bold mb-3">Sign in to your account</h1>
        <form method="POST">
          <div className="mb-3">
            <label className="block text-sm pb-1">Username</label>
            <input
              className="py-2 px-3 w-full text-sm rounded border border-solid border-gray-400 focus:border-gray-500"
              type="text"
              name="username"
              placeholder="Please input your username"
            />
          </div>
          <div className="mb-3">
            <label className="block text-sm pb-1">Password</label>
            <input
              className="py-2 px-3 w-full text-sm rounded border border-solid border-gray-400 focus:border-gray-500"
              type="password"
              name="password"
              placeholder="Please input your password"
            />
          </div>
          <div className="flex justify-between items-center mt-5 p-1">
            <Link
              to="/auth/signup"
              className="bg-indigo-100 hover:bg-indigo-200 text-sm py-2 px-6 text-indigo-600 rounded font-bold inline-block"
            >
              Sign up instead
            </Link>
            <button className="bg-indigo-500 hover:bg-indigo-700 text-sm text-white py-2 px-6 rounded font-bold">
              Next
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AuthSigninLayout;
