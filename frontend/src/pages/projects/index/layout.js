import React from "react";
import { Link, useRouteMatch } from "react-router-dom";
import Octicon, { Repo, KebabVertical, Plus as PlusIcon } from "@primer/octicons-react";

const Item = props => (
  <div class="border-bottom border-gray-light pl-4 pr-5 pb-3 pt-3 d-flex flex-column flex-md-row flex-md-items-center">
    <div class="pr-0 pr-md-3 mb-3 mb-md-0 d-flex flex-justify-center flex-md-justify-start">
      <Octicon icon={Repo} size={30} className="text-gray" />
    </div>
    <div class="text-center text-md-left flex-1">
      <h4 className="text-gray-dark">{props.title}</h4>
      <p className="text-gray">{props.description}</p>
    </div>
    <div class="ml-md-3 d-flex flex-justify-end">
      <a href="/" className="text-gray">
        <Octicon icon={KebabVertical} />
      </a>
    </div>
  </div>
);

const Layout = () => {
  const { path } = useRouteMatch();
  return (
    <React.Fragment>
      <div className="px-5">
        <header className="flex justify-between items-center border-b border-gray-200 py-2">
          <h1 className="tracking-wide text-sm">My Project</h1>
          <button className="bg-indigo-500 text-white flex justify-between items-center rounded-sm px-2 py-1">
            <Octicon icon={PlusIcon} className="h-3 w-3"></Octicon>
            <span className="text-xs pl-2">New Project</span>
          </button>
        </header>
      </div>
      <div className=""></div>
    </React.Fragment>
  );
};

export default Layout;
