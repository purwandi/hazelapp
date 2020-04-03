import React from 'react';
import { useRouteMatch } from 'react-router';
import { useQuery } from '@apollo/react-hooks';

import AppComponentLoading from '../../../components/loading';
import Layout from './layout';
import GET_ISSUES from './query';

const BacklogBoard = () => {
  const milestones = [{ id: null, name: 'Backlog', issues: [] }];
  const { params } = useRouteMatch();
  const { loading, data } = useQuery(GET_ISSUES, {
    variables: {
      owner: params.owner,
      name: params.name,
      first: 20,
    },
  });

  if (loading) {
    return <AppComponentLoading />;
  }

  data.project.milestones.map(item => milestones.push({ ...item, issues: [] }));

  // Map issue into each milestone
  if (data && data.project && data.project.issues.nodes) {
    data.project.issues.nodes.map(issue => {
      let mIndex = 0;
      if (issue.milestone !== null) {
        mIndex = milestones.findIndex(item => item.id === issue.milestone.id);
      }
      milestones[mIndex].issues.push(issue);
    });
  }

  return <Layout milestones={milestones.reverse()} />;
};

export default BacklogBoard;
