import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";
import Octicon, {
  IssueOpened as IssueOpenedIcon,
  Inbox as InboxIcon,
  Octoface as OctofaceIcon,
  Repo as RepoIcon,
  Tools as ToolsIcon,
  SignOut as SignOutIcon,
  Project as KanbanIcon,
  Milestone as MilestoneIcon,
  Gear as GearIcon,
  Saved as SavedIcon,
  CommentDiscussion as CommentDiscussionIcon
} from "@primer/octicons-react";

const NavContainer = styled.div`
  // height: calc(100vh - 66px);
  // top: 66px;
  height: calc(100vh);
  min-width: 200px;
  color: rgb(47, 54, 61);
  background-color: rgb(250, 251, 252);
  position: sticky;
`;

const OcticonWrapper = styled(Octicon)`
  width: 20px;
`;

const Navs = () => (
  <div className="d-block">
    <NavContainer>
      <div className="pl-3 pr-3 pt-2 pb-2">
        <div>
          <Link to="/" className="d-block p-1 text-gray">
            <OcticonWrapper icon={OctofaceIcon} className="mr-1" />
            <span>Purwandi</span>
          </Link>
          <Link to="/projects" className="d-block p-1 text-gray">
            <OcticonWrapper icon={RepoIcon} className="mr-1" />
            <span>My Projects</span>
          </Link>
          <Link to="/issues" className="d-block p-1 text-gray">
            <OcticonWrapper icon={InboxIcon} className="mr-1" />
            <span>My Issues</span>
          </Link>
          <Link to="/" className="d-block p-1 text-gray">
            <OcticonWrapper icon={ToolsIcon} className="mr-1" />
            <span>Settings</span>
          </Link>
          <Link to="/" className="d-block p-1 text-gray">
            <OcticonWrapper icon={SignOutIcon} className="mr-1" />
            <span>Sign Out</span>
          </Link>
        </div>
      </div>
      <div className="pl-3 pr-3 pt-2 pb-1">
        <h3 className="f5 pb-1">Marketing Website</h3>
        <div>
          <Link to="" className="d-block p-1 text-gray">
            <OcticonWrapper icon={CommentDiscussionIcon} className="mr-1" />
            <span>Discussion</span>
          </Link>
          <Link to="" className="d-block p-1 text-gray">
            <OcticonWrapper icon={IssueOpenedIcon} className="mr-1" />
            <span>Backlog</span>
          </Link>
          <Link to="" className="d-block p-1 text-gray">
            <OcticonWrapper icon={KanbanIcon} className="mr-1" />
            <span>Board</span>
          </Link>
          <Link to="" className="d-block p-1 text-gray">
            <OcticonWrapper icon={SavedIcon} className="mr-1" />
            <span>Labels</span>
          </Link>
          <Link to="" className="d-block p-1 text-gray">
            <OcticonWrapper icon={MilestoneIcon} className="mr-1" />
            <span>Milestones</span>
          </Link>
          <Link to="" className="d-block p-1 text-gray">
            <OcticonWrapper icon={GearIcon} className="mr-1" />
            <span>Setting</span>
          </Link>
        </div>
      </div>
    </NavContainer>
  </div>
);

export default Navs;
