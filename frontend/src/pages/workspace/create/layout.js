import React from 'react';
import { Link } from 'react-router-dom';

const Layout = props => {
  return (
    <div className="h-screen flex flex-col">
      <div className="flex-1 overflow-y-auto">
        <form className="w-3/5 mx-auto py-12" onSubmit={props.onCreateProject}>
          <h1 className="text-xl font-semibold">Create a project</h1>
          <p className="text-gray-700">Manage stories, task and workflows for a scrum team</p>
          <div className="mb-2 mt-4">
            <label for="name" className="block pb-1">
              Project name
            </label>
            <input
              type="text"
              id="name"
              name="name"
              className="block border border-gray-400 py-1 px-2 w-3/4 rounded shadow-inner"
              placeholder="Enter a project name"
              onChange={e => props.setState({ name: e.target.value })}
            />
            <small className="text-gray-600 text-sm">
              Great project names are short and memorable. Need inspiration?
            </small>
          </div>
          <div className="mb-2">
            <label for="key" className="block pb-1">
              Project key
            </label>
            <input
              type="text"
              id="key"
              name="key"
              className="block border border-gray-400 py-1 px-2 w-1/4 rounded shadow-inner"
              placeholder="Eg. AT"
            />
            <small className="text-gray-600 text-sm">Eg. AT (for a project named Atlantis)</small>
          </div>
          <div className="mb-2">
            <label className="block pb-1" for="description">
              Description
              <span className="text-gray-600 text-sm"> (optional)</span>
            </label>
            <textarea
              id="description"
              name="description"
              className="block border border-gray-400 py-1 px-2 w-full rounded shadow-inner"
              placeholder="Enter a project description"
              rows="3"
              onChange={e => props.setState({ description: e.target.value })}
            ></textarea>
          </div>

          <div className="mt-6">
            <button className="bg-indigo-500 text-white font-semibold rounded px-5 py-2">Create project</button>
            <Link to="/" className="inline-block px-5 py-2 ml-4">
              Cancel
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Layout;
