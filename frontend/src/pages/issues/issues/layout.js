import React from "react";
import { Link } from "react-router-dom";
import Octicon, { Plus } from "@primer/octicons-react";

import AppContainer from "../../../components/container";
import Header from "../../../components/header";

const Separator = props => (
  <div className="d-flex border-bottom pr-3 pl-3 pb-2 pt-2 bg-gray-light">
    <div className="pr-2 pl-2">
      <input type="checkbox" />
    </div>
    <div className="pr-2 pl-2 flex-1 h5">{props.title}</div>
    <div className="pr-2 pl-2">
      <Octicon icon={Plus} />
    </div>
  </div>
);

const Item = () => (
  <div className="d-flex border-bottom pr-3 pl-3 pb-2 pt-2">
    <div className="pr-2 pl-2">
      <input type="checkbox" />
    </div>
    <div className="pr-2 pl-2">XD-11</div>
    <div className="pr-2 pl-2 flex-1">Body</div>
    <div className="pr-2 pl-2">meta</div>
  </div>
);

const HeaderActions = () => (
  <Link to="/issues/create">
    <Octicon icon={Plus} size={22} />
  </Link>
);

const Issues = () => (
  <React.Fragment>
    <Header title="Issues" Actions={HeaderActions} />
    <AppContainer>
      <Separator title="Ready for testing" />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Separator title="In Review" />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Separator title="In Design" />
      <Item />
      <Item />
      <Item />
      <Separator title="Todo" />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
      <Item />
    </AppContainer>
  </React.Fragment>
);

export default Issues;
