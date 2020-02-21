import React from "react";

const Header = props => {
  const { Actions } = props;
  return (
    <div className="pl-4 pr-4 pt-3 pb-3 border-bottom">
      <div className="d-flex flex-justify-between">
        <div className="flex-1">
          <h4>{props.title}</h4>
        </div>
        <div className="d-block">{Actions && <Actions />}</div>
      </div>
    </div>
  );
};

export default Header;
