import React from "react";

const Header = props => {
  return (
    <header className="flex justify-between items-center py-3 border-b border-gray-200">
      <h1 className="tracking-wider font-medium px-5 flex-1">{props.name}</h1>
      <div className="px-5">{props.actions && props.actions}</div>
    </header>
  );
};

export default Header;
