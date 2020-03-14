import React from "react";
import { Link } from "react-router-dom";
import Octicon, {
  Inbox as InboxIcon,
  Octoface as OctofaceIcon,
  Repo as RepoIcon,
  Tools as ToolsIcon,
  SignOut as SignOutIcon,
  Organization as OrganizationIcon,
  CommentDiscussion as CommentDiscussionIcon,
  IssueOpened as IssueOpenedIcon,
  Project as KanbanIcon,
  Graph as GraphIcon,
  Gear as GearIcon
} from "@primer/octicons-react";

const Navigation = () => {
  return (
    <React.Fragment>
      <nav className="tracking-wide text-xs">
        <Link className="rounded flex justify-between items-center px-3 py-1 bg-gray-200 rounded text-gray-900">
          <Octicon icon={OctofaceIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Purwandi</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 hover:bg-gray-300 text-gray-900">
          <Octicon icon={RepoIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">My Project</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 hover:bg-gray-300 text-gray-700">
          <Octicon icon={InboxIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">My Issues</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 hover:bg-gray-300 text-gray-700">
          <Octicon icon={ToolsIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Settings</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 hover:bg-gray-300 text-gray-700">
          <Octicon icon={SignOutIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Sign Out</span>
        </Link>
      </nav>
      <nav className="mt-4 tracking-wide text-xs">
        <h2 className="text-gray-700 font-semibold mb-2">Workspace</h2>
        <Link className="rounded flex justify-between items-center px-3 py-1 rounded text-gray-900">
          <Octicon icon={OrganizationIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Foobar Team</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={OrganizationIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Changelog Team</span>
        </Link>
      </nav>
      <nav className="mt-4 tracking-wide text-xs">
        <h2 className="text-gray-700 font-semibold mb-2">Favorites Project</h2>
        <Link className="rounded flex justify-between items-center px-3 py-1 rounded text-gray-900">
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
          <Octicon icon={GraphIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Insight</span>
        </Link>
        <Link className="rounded flex justify-between items-center px-3 py-1 text-gray-700">
          <Octicon icon={GearIcon} className="w-3 h-3"></Octicon>
          <span className="flex-1 ml-3">Project Settings</span>
        </Link>
      </nav>
    </React.Fragment>
  );
};

export default Navigation;
