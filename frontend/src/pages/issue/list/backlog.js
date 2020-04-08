import React from 'react';
import { object } from 'prop-types';
import FormIssueList from '../create';
import Item from './issue';
import MilestoneCreate from './milestone';

const Backlog = ({ project, milestone }) => (
  <div className="mb-8">
    <div className="flex items-center justify-between text-sm">
      <h2 className="font-semibold">{milestone.name}</h2>
      <div className="ml-3">
        <div className="text-gray-500">
          {(milestone.issues && milestone.issues.length) || 0} issues
        </div>
      </div>
      <div className="flex-1 text-right">
        {milestone.id === null && <MilestoneCreate project={project} />}
      </div>
    </div>
    <div className="text-gray-600 text-xs">
      {milestone.description && milestone.description} &nbsp;
    </div>
    <div className="mt-2 mb-1">
      {milestone.issues.length === 0 && (
        <div className="border border-dashed border-gray-400">
          <div className="text-center p-4 text-sm text-gray-600">Your backlog is empty</div>
        </div>
      )}
      {milestone.issues.length > 0 && (
        <div className="">
          {milestone.issues.map((issue) => (
            <Item issue={issue} key={issue.id} />
          ))}
        </div>
      )}
      <FormIssueList project={project} milestone={milestone} />
    </div>
  </div>
);

Backlog.propTypes = {
  project: object.isRequired,
  milestone: object.isRequired,
};

export default Backlog;
