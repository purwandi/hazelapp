import React from 'react';
import { string, node } from 'prop-types';

const Header = ({ name, actions }) => (
  <header className="flex justify-between items-center py-3 border-b border-gray-200">
    <h1 className="tracking-wider font-medium px-5 flex-1">{name}</h1>
    <div className="px-5">{actions && actions}</div>
  </header>
);

Header.propTypes = {
  name: string.isRequired,
  actions: node,
};

Header.defaultProps = {
  actions: <span />,
};

export default Header;
