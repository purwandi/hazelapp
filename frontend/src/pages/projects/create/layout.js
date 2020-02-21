import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";

const FormContainer = styled.div`
  max-width: 600px;
  margin: 0 auto;
`;

const Layout = () => {
  return (
    <React.Fragment>
      <div className="pl-4 pr-4 pt-3 pb-2 border-bottom border-gray">
        <div className="d-flex flex-justify-between">
          <div className="">
            <p>CREATE PROJECT</p>
          </div>
        </div>
      </div>

      <FormContainer className="pl-4 pr-4 pt-4">
        <h1 className="h2 mb-1">Create Project</h1>
        <p className="text-gray">Create new project to manage your</p>
        <form>
          <div className="form-group">
            <label className="d-block">Project Name</label>
            <input type="text" className="form-control" />
          </div>
          <div className="form-group">
            <label className="d-block">Project Description</label>
            <textarea type="text" rows="4" className="form-control"></textarea>
          </div>
          <div className="form-actions float-left pt-4">
            <button type="submit" className="float-left btn btn-primary">
              Save changes
            </button>
            <Link to="/projects" className="float-left btn btn-invisible">
              Cancel
            </Link>
          </div>
        </form>
      </FormContainer>
    </React.Fragment>
  );
};

export default Layout;
