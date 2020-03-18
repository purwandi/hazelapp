import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Octicon, { Home as HomeIcon, Tools as ToolsIcon } from '@primer/octicons-react';

const NavItem = (props) => {
  const isActive = props.active ? 'bg-gray-300' : '';
  return (
    <Link
      to={props.to}
      className={`rounded px-3 py-1 ${isActive} hover:bg-gray-200 block flex justify-between items-center rounded text-gray-900 cursor-pointer`}
    >
      <Octicon icon={props.icon} className="w-3 h-3" />
      <span className="ml-2 flex-1">{props.label}</span>
    </Link>
  );
};

const NavProfile = () => {
  const [menu, setMenu] = useState(false);
  const active = menu === true ? 'block' : 'hidden';

  return (
    <div className="relative">
      <div className="flex justify-between items-start cursor-pointer" onClick={(e) => setMenu(true)}>
        <div className="w-10 h-10">
          <img
            src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
            className="rounded inline-block"
            alt="avatar"
          />
        </div>
        <div className="flex-1 flex flex-col pl-2">
          <span className="text-sm">Facebook</span>
          <span className="text-xs text-gray-700">@purwandi</span>
        </div>
      </div>
      <div className={`absolute top-0 left-0 bg-white w-64 z-10  ${active}`}>
        <div className="">
          <nav className="text-sm">
            <span>Profile</span>
          </nav>
        </div>
      </div>
    </div>
  );
};

const Navigation = () => (
  <>
    <NavProfile />
    <nav className="mt-6 tracking-wide text-sm">
      <NavItem icon={HomeIcon} to="/" label="Dashboard" active />
      <NavItem icon={ToolsIcon} label="Settings" />
    </nav>
  </>
);

export default Navigation;
