import React from "react";

import AppContainer from "../../../components/container";
import Header from "../../../components/header";

const Layout = () => (
  <React.Fragment>
    <Header title="Create Issue" />
    <div className="d-flex">
      <AppContainer className="flex-1">
        <p>DDd</p>
      </AppContainer>
      <div className="">Content</div>
    </div>
  </React.Fragment>
);

export default Layout;
