import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Octicon, {
  Home as HomeIcon,
  Tools as ToolsIcon,
  SignOut as SignOutIcon,
  IssueOpened as IssueOpenedIcon,
  Note as NoteIcon,
  CommentDiscussion as CommentDiscussionIcon,
  Question as QuestionIcon,
  Organization as OrganizationIcon,
  Gear as GearIcon,
} from '@primer/octicons-react';
import { UseAppContextValue } from '../../context';

const NavItem = props => {
  const isActive = props.active ? 'bg-gray-300' : '';
  return (
    <Link
      to={props.to}
      className={`rounded px-3 py-1 ${isActive} hover:bg-gray-200 block flex justify-between items-center rounded text-gray-900 cursor-pointer`}
      {...props}
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
    <div className="relative p-3">
      <div className="flex justify-between items-start cursor-pointer" onClick={e => setMenu(true)}>
        <div className="w-10 h-10">
          <img
            src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
            className="rounded inline-block"
            alt="avatar"
          />
        </div>
        <div className="flex-1 flex flex-col pl-2">
          <span className="text-sm">Purwandi</span>
          <span className="text-xs text-gray-700">@purwandi</span>
        </div>
      </div>
      <div className={`absolute left-0 bg-white w-64 z-10 shadow rounded px-4 hidden`}>
        <div className="border border-gray-200 ">
          <nav className="text-sm">
            <a href=""></a>
          </nav>
        </div>
      </div>
    </div>
  );
};

const Navigation = () => {
  const [, dispatch] = UseAppContextValue();
  return (
    <>
      <NavProfile />
      <div className="overflow-auto flex-1 p-3 ">
        <nav className="tracking-wide text-sm">
          <NavItem icon={HomeIcon} to="/" label="Dashboard" active={true} />
          <NavItem icon={ToolsIcon} label="Settings" />
          {/* <NavItem icon={SignOutIcon} label="Sign Out" onClick={e => dispatch({ type: 'LOGOUT' })} /> */}
        </nav>
        <nav className="mt-5 tracking-wide text-sm">
          <h2 className="text-gray-700 font-semibold mb-2 px-3">Symfony</h2>
          <NavItem icon={CommentDiscussionIcon} to="/" label="Discussion" />
          <NavItem icon={NoteIcon} label="Backlog" />
          <NavItem icon={IssueOpenedIcon} to="/" label="Issue" />
          <NavItem icon={OrganizationIcon} to="/" label="Member" />
          <NavItem icon={GearIcon} label="Project Settings" />
        </nav>
      </div>
      <div className="p-3">
        <nav className="mb-3 tracking-wide text-sm">
          <NavItem icon={QuestionIcon} label="Help" />
          <NavItem icon={SignOutIcon} label="Sign Out" onClick={e => dispatch({ type: 'LOGOUT' })} />
        </nav>
      </div>
    </>
  );
};

export default Navigation;
