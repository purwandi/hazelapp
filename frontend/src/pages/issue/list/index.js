import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

import AppComponentLoading from '../../../components/loading';
import Layout from './layout';
import { useRouteMatch } from 'react-router';

const GET_ISSUES = gql`
  query($owner: String!, $name: String!, $first: Int) {
    project(owner: $owner, name: $name) {
      milestones {
        id
        name
      }
      issues(input: { first: $first }) {
        nodes {
          id
          number
          title
          body
          milestone {
            id
          }
        }
      }
    }
  }
`;

const BacklogBoard = () => {
  const milestones = [{ id: null, name: 'backlog', issues: [] }];
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

  // Add backlog milestone
  if (data && data.project) {
    data.project.milestones.push({ id: 'backlog', name: 'Backlog', issues: [] });
  }

  // Map issue into each milestone
  if (data && data.project && data.project.issues.nodes) {
    data.project.issues.nodes.map(issue => {
      let mIndex;
      if (issue.milestone !== null) {
        if (data.project && milestones.find(item => issue.milestone.id !== item.id)) {
          data.milestones.push({ ...issue.milestone, issues: [] });
        }

        mIndex = data.project.milestones.findIndex(item => item.id === issue.milestone.id);
      } else {
        mIndex = data.project.milestones.findIndex(item => item.id === 'backlog');
      }

      data.project.milestones[mIndex].issues.push(issue);
    });
  }

  return <Layout milestones={data.project.milestones} />;
};

export default BacklogBoard;
