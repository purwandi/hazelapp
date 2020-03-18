import React from 'react';
import Octicon, { Plus as PlusIcon } from '@primer/octicons-react';

const Layout = () => (
  <div className="px-5">
    <header className="flex justify-between items-center border-b border-gray-200 py-2">
      <h1 className="tracking-wide text-sm">My Project</h1>
      <button
        type="button"
        className="bg-indigo-500 text-white flex justify-between items-center rounded-sm px-2 py-1"
      >
        <Octicon icon={PlusIcon} className="h-3 w-3" />
        <span className="text-xs pl-2">New Project</span>
      </button>
    </header>
  </div>
);
export default Layout;
