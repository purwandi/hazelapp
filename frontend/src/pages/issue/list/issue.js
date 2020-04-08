import React from 'react';
import { object } from 'prop-types';
import { IssueType } from '../static';

const Item = ({ issue }) => (
  <div className="flex justify-between items-center border-gray-300 border-b first:border-t py-2 px-2 text-sm">
    <div className="flex-1 truncate flex items-center">
      <IssueType value={issue.issueType} />
      <span className="text-gray-800 font-lighter tracking-wide ml-2">{issue.title}</span>
    </div>
    <div className="">
      <div className="flex justify-end items-center">
        <div className="pl-2 text-gray-500">{`#${issue.number}`}</div>
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

Item.propTypes = {
  issue: object.isRequired,
};

export default Item;
