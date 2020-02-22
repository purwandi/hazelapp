import React from "react";
import { Link } from "react-router-dom";
import styled from "styled-components";
import TextareaAutosize from "react-autosize-textarea";
import Octicon, { Gear } from "@primer/octicons-react";

import Header from "../../../components/header";

const AsideContainer = styled.div`
  width: 250px;
  height: calc(100vh - 57px);
`;

const Layout = () => {
  return (
    <React.Fragment>
      <Header title="Create Issue" />
      <div className="d-flex">
        <div className="flex-1">
          <div className="pl-4 pr-4 pt-3">
            <div className="form-group">
              <input type="text" class="form-control" placeholder="Issue title" />
            </div>
            <div className="form-group">
              <TextareaAutosize className="form-control" rows={6} placeholder="Leave a comment" />
            </div>
            <div className="form-actions border-top pt-3">
              <button type="submit" className="btn btn-primary" disabled>
                Save changes
              </button>
              <Link to="/issues" class="btn btn-invisible">
                Cancel
              </Link>
              <small className="float-left">Styling with Markdown is supported</small>
            </div>
          </div>
        </div>
        <div className="d-block">
          <AsideContainer className="p-3">
            <div className="form-group">
              <a href="/" className="d-block">
                <div className="float-right">
                  <Octicon icon={Gear} />
                </div>
                <label>Status</label>
              </a>
              <p>Closed</p>
            </div>
            <div className="form-group">
              <a href="/" className="d-block">
                <div className="float-right">
                  <Octicon icon={Gear} />
                </div>
                <label>Labels</label>
              </a>
              <p>Closed</p>
            </div>
            <div className="form-group">
              <a href="/" className="d-block">
                <div className="float-right">
                  <Octicon icon={Gear} />
                </div>
                <label>Assignees</label>
              </a>
              <p>Closed</p>
            </div>
          </AsideContainer>
        </div>
      </div>
    </React.Fragment>
  );
};

export default Layout;
