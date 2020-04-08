import React from 'react';
import { useRouteMatch } from 'react-router';
import { useQuery } from '@apollo/react-hooks';

import AppComponentLoading from '../../../components/loading';
import Layout from './layout';
import { GET_ISSUES } from '../query';

const BacklogBoard = () => {
  const milestones = [];
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

  // Add backlog milestone if the issues doesn't hva milestone
  data.project.milestones.map((item) => milestones.push({ ...item, issues: [] }));
  milestones.push({ id: null, name: 'Backlog', issues: [] });

  // Map issue into each milestone
  if (data && data.project && data.project.issues.nodes) {
    data.project.issues.nodes.map((issue) => {
      const milestoneID = issue.milestone !== null ? issue.milestone.id : null;
      const mIndex = milestones.findIndex((item) => item.id === milestoneID);
      milestones[mIndex].issues.push(issue);
    });
  }

  return <Layout project={{ id: data.project.id }} milestones={milestones} />;
};

export default BacklogBoard;
