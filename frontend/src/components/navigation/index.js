import React, { useState } from 'react';
import { useRouteMatch, useLocation } from 'react-router-dom';
import {
  Home as HomeIcon,
  Tools as ToolsIcon,
  SignOut as SignOutIcon,
  Note as NoteIcon,
  Comment as CommentIcon,
  CommentDiscussion as CommentDiscussionIcon,
  Question as QuestionIcon,
  Organization as OrganizationIcon,
  Gear as GearIcon,
  IssueReopened as IssueReopenedIcon,
  Code as CodeIcon,
  GitPullRequest as GitPullRequestIcon,
} from '@primer/octicons-react';
import { UseAppContextValue } from '../../context';

import NavItem from './item';

const NavProfile = () => {
  const [state] = UseAppContextValue();
  const [menu, setMenu] = useState(false);
  const active = menu === true ? 'block' : 'hidden';

  return (
    <div className="relative p-3">
      <div className="flex justify-between items-start cursor-pointer" onClick={e => setMenu(!menu)}>
        <div className="w-10 h-10">
          <img
            src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
            className="rounded inline-block"
            alt="avatar"
          />
        </div>
        <div className="flex-1 flex flex-col pl-2">
          <span className="text-sm">{state.user.fullname}</span>
          <span className="text-xs text-gray-700">{`@${state.user.username}`}</span>
        </div>
      </div>
      <div className={`absolute left-0 bg-white w-64 z-10 shadow rounded px-4 ${active}`}>
        <div className="border border-gray-200 ">
          <nav className="text-sm"></nav>
        </div>
      </div>
    </div>
  );
};

const Navigation = () => {
  const [, dispatch] = UseAppContextValue();
  const { pathname } = useLocation();
  const match = useRouteMatch('/:owner/:name');

  console.log(match);

  return (
    <>
      <NavProfile />
      <div className="overflow-auto flex-1 p-3 ">
        <nav className="tracking-wide text-sm">
          <NavItem icon={HomeIcon} to="/" label="Dashboard" />
          <NavItem icon={CommentDiscussionIcon} to="/" label="Threads" />
          <NavItem icon={ToolsIcon} to="/" label="Settings" />
        </nav>
        {match && (
          <nav className="mt-5 tracking-wide text-sm">
            <h2 className="text-gray-700 font-semibold mb-2 px-3 truncate" title={match.params.name}>
              {match.params.name}
            </h2>
            <NavItem icon={CommentIcon} to={`${match.url}`} label="Discussion" />
            <NavItem icon={NoteIcon} to={`${match.url}/backlog`} label="Backlog" />
            <NavItem icon={IssueReopenedIcon} to={`${match.url}/active-sprints`} label="Active Sprint" />
            <NavItem icon={CodeIcon} to={`${match.url}/code`} label="Code" />
            <NavItem icon={GitPullRequestIcon} to={`${match.url}/pulls`} label="Pull Request" />
            <NavItem icon={OrganizationIcon} to={`${match.url}/members`} label="Member" />
            <NavItem icon={GearIcon} to={`${match.url}/settings`} label="Project Settings" />
          </nav>
        )}
      </div>
      <div className="p-3">
        <nav className="mb-2 tracking-wide text-sm">
          <NavItem icon={QuestionIcon} to="/" label="Help" />
          <NavItem icon={SignOutIcon} to="/" label="Sign Out" onClick={e => dispatch({ type: 'LOGOUT' })} />
        </nav>
      </div>
    </>
  );
};

export default Navigation;
