import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

import ComponentLoading from '../../../components/loading';
import Layout from './layout';

const GET_PROJECTS = gql`
  query($first: Int, $last: Int, $after: String, $before: String) {
    viewer {
      username
      projects(input: { first: $first, last: $last, after: $after, before: $before }) {
        nodes {
          id
          name
          description
        }
      }
    }
  }
`;

const Workspace = (props) => {
  const { loading, data } = useQuery(GET_PROJECTS, {
    variables: {
      first: 20,
    },
  });

  if (loading) {
    return <ComponentLoading />;
  }

  if (data && data.viewer) {
    return <Layout login={data.viewer.username} projects={data.viewer.projects.nodes} {...props} />;
  }
};

export default Workspace;
