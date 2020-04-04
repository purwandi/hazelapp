import React from 'react';
import { array, object } from 'prop-types';
import Header from '../../../components/header';
import Backlog from './backlog';

const BacklogBoardLayout = ({ milestones, project }) => {
  return (
    <div className="h-screen flex flex-col">
      <Header name="Backlog" />
      <div className="flex-1 overflow-y-auto p-5">
        {milestones.map(milestone => (
          <Backlog project={project} milestone={milestone} key={milestone.id} />
        ))}
      </div>
    </div>
  );
};

BacklogBoardLayout.propTypes = {
  milestones: array.isRequired,
  project: object.isRequired
}

export default BacklogBoardLayout;
