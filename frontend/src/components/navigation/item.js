import React from 'react';
import { bool, func, string } from 'prop-types';
import { Link } from 'react-router-dom';
import Octicon from '@primer/octicons-react';

const NavItem = props => {
  const isActive = props.active ? 'bg-gray-300' : '';

  return (
    <Link
      to={props.to}
      className={`rounded px-3 text-gray-700 py-1 ${isActive} hover:bg-gray-200 block flex justify-between items-center rounded cursor-pointer`}
    >
      <Octicon icon={props.icon} className="w-4" />
      <span className="ml-2 flex-1">{props.label}</span>
    </Link>
  );
};

NavItem.propTypes = {
  active: bool,
  icon: func,
  label: string,
};

export default NavItem;
