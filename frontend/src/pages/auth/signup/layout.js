import React from "react";
import { Link } from "react-router-dom";

const AuthSignupLayut = props => {
  return (
    <div className="min-h-screen flex items-center justify-center py-12 px-4">
      <div className="max-w-sm w-full">
        <h1 className="text-3xl font-bold mb-3">Register new acccount</h1>
        <form>
          <div className="mb-2">
            <label className="block text-sm pb-1">Fullname</label>
            <input
              className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
              type="text"
              name="fullname"
              placeholder="Please enter your fullname"
            />
          </div>
          <div className="mb-2">
            <label className="block text-sm pb-1">Username</label>
            <input
              className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
              type="text"
              name="username"
              placeholder="Please pick a username for your login"
            />
          </div>
          <div className="mb-2">
            <div className="flex mb-2">
              <div className="w-full pr-5">
                <label className="block text-sm pb-1">Password</label>
                <input
                  className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
                  type="password"
                  name="password"
                  placeholder="Type your password"
                />
              </div>
              <div className="w-full pl-5">
                <label className="block text-sm pb-1">Confirm</label>
                <input
                  className="block px-3 py-2 w-full text-sm border border-gray-400 rounded"
                  type="password"
                  name="confirm"
                  placeholder="Type your password confirmation"
                />
              </div>
            </div>
            <p className="text-gray-600 text-xs">
              Use 8 or more characters with a mix of letters, numbers & symbols
            </p>
          </div>

          <div className="flex justify-between items-center mt-5 p-1">
            <Link
              to="/auth/signin"
              className="bg-indigo-100 hover:bg-indigo-200 text-sm py-2 px-6 text-indigo-600 rounded font-bold inline-block"
            >
              Sign in instead
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

export default AuthSignupLayut;
