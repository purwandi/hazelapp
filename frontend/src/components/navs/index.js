import React from "react";
import styled from "styled-components";

const NavContainer = styled.div`
  // height: calc(100vh - 66px);
  // top: 66px;
  height: calc(100vh);
  min-width: 200px;
  color: rgb(47, 54, 61);
  background-color: rgb(250, 251, 252);
  position: sticky;
`;

const Navs = () => (
  <div className="d-block">
    <NavContainer />
  </div>
);

export default Navs;
