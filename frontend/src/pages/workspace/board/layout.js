import React from 'react';
import { Link } from 'react-router-dom';
import Octicon, { Plus as PlusIcon } from '@primer/octicons-react';
import { array, string } from 'prop-types';
import Header from '../../../components/header';

const Item = ({ login, name, description }) => (
  <Link
    to={`/${login}/${name}`}
    className="p-2 xl:w-1/4 lg:w-1/3 md:w-1/2 sm:w-1/2 cursor-pointer "
  >
    <div className="border border-gray-300 rounded hover:shadow">
      <div className="p-3 border-b-2 border-gray-300">
        <h2 className="font-semibold text-sm mb-2 tracking-wide">{name}</h2>
        <div className="text-gray-700 text-xs h-10 overflow-y-hidden">{description}</div>
      </div>
      <div className="p-3 text-xs flex justify-between">
        <div className="">
          <span className="font-semibold">8 </span>
          <span className="text-gray-700">issues</span>
        </div>
        <div className="">
          <span className="font-semibold">1 </span>
          <span className="text-gray-700">pull request</span>
        </div>
        <div className="">
          <span className="font-semibold">2 </span>
          <span className="text-gray-700">issues</span>
        </div>
      </div>
    </div>
  </Link>
);

Item.propTypes = {
  login: string.isRequired,
  name: string.isRequired,
  description: string.isRequired,
};

const Actions = () => (
  <Link to="/create" className="bg-indigo-500 text-white px-2 py-1 rounded text-sm">
    <Octicon icon={PlusIcon} />
  </Link>
);

const Layout = ({ projects, login }) => (
  <div className="h-screen flex flex-col">
    <Header name="Workspace" actions={<Actions />} />
    <div className="flex-1 flex content-start flex-wrap p-4 overflow-y-auto">
      {projects.map((item) => (
        <Item
          login={login}
          id={item.id}
          name={item.name}
          description={item.description}
          key={item.id}
        />
      ))}
    </div>
  </div>
);

Layout.propTypes = {
  projects: array.isRequired,
  login: string.isRequired,
};

export default Layout;
