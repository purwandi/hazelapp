import React from "react";
import { Link, useRouteMatch } from "react-router-dom";
import Octicon, { Repo, KebabVertical } from "@primer/octicons-react";

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
      <div className="pl-4 pr-4 pt-3 pb-2 border-bottom">
        <div className="d-flex flex-justify-between">
          <div className="">
            <h4>Project</h4>
          </div>
          <div className="d-block">
            <Link to={`${path}create`} className="btn btn-sm">
              New Project
            </Link>
          </div>
        </div>
      </div>
      <Item title="Marketing Website" description="Add new features" />
      <Item title="Projects" description="Implement project management for linear issues" />
      <Item title="Cycles" description="Initial version of cycles" />
      <Item title="Filters" description="Robus filtering on all views" />
      <Item title="Mobile App" description="First version of the mobile app for IOS and Android" />
      <Item title="Cycles" description="Initial version of cycles" />
      <Item title="Filters" description="Robus filtering on all views" />
      <Item title="Mobile App" description="First version of the mobile app for IOS and Android" />
      <Item title="Cycles" description="Initial version of cycles" />
      <Item title="Filters" description="Robus filtering on all views" />
      <Item title="Mobile App" description="First version of the mobile app for IOS and Android" />
    </React.Fragment>
  );
};

export default Layout;
