import React, { useState } from "react";
import { Link } from "react-router-dom";
import Octicon, {
  GithubAction as GithubActionIcon,
  Plus as PlusIcon
} from "@primer/octicons-react";

const NavItem = props => {
  const isActive = props.active ? "bg-gray-300" : "";
  return (
    <Link
      to={props.to}
      className={`rounded px-3 py-1 ${isActive} hover:bg-gray-200 block rounded text-gray-900 flex items-center cursor-pointer`}
    >
      <Octicon icon={props.icon} className="w-3 h-3 block"></Octicon>
      <div className="flex-1 ml-3 tracking-wide">{props.label}</div>
      {/* <div className="">
        <span className="rounder bg-gray-400 text-gray-900 rounded px-2 text-xs">9</span>
      </div> */}
    </Link>
  );
};

const NavProfile = () => {
  const [menu, setMenu] = useState(false);
  const active = menu == true ? "block" : "hidden";

  return (
    <div className="relative">
      <div className="flex justify-between items-start cursor-pointer" onClick={e => setMenu(true)}>
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

const Navigation = () => {
  return (
    <React.Fragment>
      <NavProfile />
      {/* <nav className="tracking-wide text-sm mt-4">
        <NavItem icon={GithubActionIcon} label="My Workspace" active />
      </nav> */}
      <nav className="mt-6 tracking-wide text-sm">
        <h2 className="text-gray-700 font-semibold mb-2 px-3">Favourites</h2>
        <NavItem icon={GithubActionIcon} label="React" active />
        <NavItem icon={GithubActionIcon} label="React Native" />
        <NavItem icon={PlusIcon} label="Add a project" />
        {/* <Link className="rounded flex justify-between items-center px-3 py-1 rounded text-gray-900">
          <Octicon icon={CommentDiscussionIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Discussion</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={IssueOpenedIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Backlog</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={KanbanIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Board</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={CodeIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Code</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={GitPullRequestIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Pull Request</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={GraphIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Insight</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={GearIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Project Settings</span>
        </Link> */}
      </nav>
    </React.Fragment>
  );
};

export default Navigation;
