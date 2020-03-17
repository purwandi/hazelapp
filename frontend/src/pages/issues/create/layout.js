import React, { useState } from "react";
import { Link } from "react-router-dom";
import Octicon, { ChevronLeft as ChevronLeftIcon, Gear as GearIcon } from "@primer/octicons-react";

import AppEditor from "../../../components/editor";
import AppComment from "../../../components/comment";

const IssueSidebar = () => (
  <div className="w-64 text-sm py-10 px-8 border-l border-gray-200">
    <div className="">
      <div className="flex justify-between mb-1 text-gray-700 hover:text-indigo-500 cursor-pointer">
        <div className="text-xs font-semibold">Status</div>
        <Octicon icon={GearIcon} />
      </div>
      <div className="text-xs text-gray-700">No status</div>
    </div>
    <div className="mt-3 pt-3">
      <div className="flex justify-between mb-1 text-gray-700 hover:text-indigo-500 cursor-pointer">
        <div className="text-xs font-semibold">Assignees</div>
        <Octicon icon={GearIcon} />
      </div>
      <div className="text-xs text-gray-700">
        No one - <span className="cursor-pointer hover:text-indigo-500">assign yourself</span>
      </div>
    </div>
    <div className="mt-3 pt-3">
      <div className="flex justify-between mb-1 text-gray-700 hover:text-indigo-500 cursor-pointer">
        <div className="text-xs font-semibold">Labels</div>
        <Octicon icon={GearIcon} />
      </div>
      <div className="text-xs text-gray-700">None yet</div>
    </div>
    <div className="mt-3 pt-3">
      <div className="flex justify-between mb-1 text-gray-700 hover:text-indigo-500 cursor-pointer">
        <div className="text-xs font-semibold">Reporter</div>
        <Octicon icon={GearIcon} />
      </div>
      <div className="text-xs text-gray-700">None yet</div>
    </div>
    <div className="mt-3 pt-3">
      <div className="flex justify-between mb-1 text-gray-700 hover:text-indigo-500 cursor-pointer">
        <div className="text-xs font-semibold">Priority</div>
        <Octicon icon={GearIcon} />
      </div>
      <div className="text-xs text-gray-700">None yet</div>
    </div>
    <div className="mt-3 pt-3">
      <div className="flex justify-between mb-1 text-gray-700 hover:text-indigo-500 cursor-pointer">
        <div className="text-xs font-semibold">Milestones</div>
        <Octicon icon={GearIcon} />
      </div>
      <div className="text-xs text-gray-700">None yet</div>
    </div>
    <div className="mt-3 pt-3 border-t border-gray-200">
      <div className="flex justify-between mb-1">
        <div className="text-xs font-semibold text-gray-700">Linked pull requests</div>
      </div>
      <p className="text-xs mb-2 text-gray-700">
        Successfully merging a pull request may close this issue.
      </p>
      <p className="text-xs text-gray-700">None yet</p>
    </div>
  </div>
);

const IssueHeader = () => (
  <header className="flex justify-between items-center py-2 px-5 border-b border-gray-200">
    <div className="inline-flex pr-5">
      <Octicon icon={ChevronLeftIcon} className="text-gray-500" size={28} />
    </div>
    <div className="flex-1 flex items-center">
      <span className="text-sm tracking-wide font-semibold">Awesome projects</span>
      <span className="mx-2 text-gray-300"> / </span>
      <span className="text-sm tracking-wide text-gray-700">New Issue</span>
    </div>
    <div className="flex justify-between items-center">
      <Link to="" className="text-sm px-4 py-1 inline-block rounded-full text-gray-700">
        Cancel
      </Link>
      <button className="text-sm ml-2 px-4 py-1 rounded bg-indigo-500 text-white">Save</button>
    </div>
  </header>
);

const Layout = () => {
  const [value, setValue] = useState("");
  return (
    <div className="h-screen flex flex-col">
      <IssueHeader />
      <div className="flex-1 flex justify-between overflow-y-auto">
        <div className="px-10 max-w-screen-md mx-auto">
          <AppEditor value={value} setValue={setValue} />
          <AppComment />
        </div>
        <IssueSidebar />
      </div>
    </div>
  );
};

export default Layout;
