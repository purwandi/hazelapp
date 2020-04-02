import React from 'react';
import Header from '../../../components/header';
import Octicon, { PrimitiveDot as PrimitiveDotIcon } from '@primer/octicons-react';

const Item = props => (
  <div className="flex justify-between items-center border-b border-gray-200 py-1 text-sm">
    <div className="flex-1 truncate">
      <span className="text-gray-800 font-lighter tracking-wide">{props.title}</span>
    </div>
    <div className="">
      <div className="flex justify-end items-center">
        <div className="pl-2 text-gray-600">{`#${props.number}`}</div>
        {/* <div className="flex justify-between">
          <div className="px-2 ml-1 inline-block border border-gray-200 text-red-700 rounded-lg">
            <Octicon icon={PrimitiveDotIcon} />
            <span className="ml-1">Bug</span>
          </div>
          <div className="px-2 ml-1 inline-block border border-gray-200 text-yellow-700 rounded-full">
            <Octicon icon={PrimitiveDotIcon} />
            <span className="ml-1">Bug</span>
          </div>
        </div>
        <div className="px-2 text-gray-600 text-xs">Apr 29</div>
        <div className="">
          <img
            src="https://avatars.slack-edge.com/2020-02-13/949588721668_1ea008d5d04743ef5131_72.png"
            className="rounded-full inline-block w-6"
            alt="avatar"
          />
        </div> */}
      </div>
    </div>
  </div>
);

const Backlog = props => (
  <div className="mb-8">
    <div className="flex items-center text-sm">
      <h2 className="font-semibold">{props.title}</h2>
      <div className="ml-3">
        <div className="text-gray-500">{(props.issues && props.issues.length) || 0} issues</div>
      </div>
    </div>
    <div className="text-gray-600 text-xs">Implement the new weather alert system and make over 50,000+ customers</div>
    <div className="mt-2 mb-1">
      {props.issues && props.issues.map(issue => <Item key={issue.id} title={issue.title} number={issue.number} />)}
    </div>
  </div>
);

const BacklogBoardLayout = props => {
  return (
    <div className="h-screen flex flex-col">
      <Header name="Backlog" />
      <div className="flex-1 overflow-y-auto p-5">
        {props.milestones.map(milestone => (
          <Backlog title={milestone.name} key={milestone.id} issues={milestone.issues} />
        ))}
      </div>
    </div>
  );
};

export default BacklogBoardLayout;
