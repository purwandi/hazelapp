import React from 'react';
import gql from 'graphql-tag';
import { useSetState } from 'react-use';
import { useMutation } from '@apollo/react-hooks';

import Layout from './layout';

const CREATE_PROJECT = gql`
  mutation CreateProject($name: String!, $description: String) {
    ProjectCreate(input: { name: $name, description: $description }) {
      id
      name
      description
    }
  }
`;

const WorkspaceCreate = props => {
  const [state, setState] = useSetState({
    name: '',
    description: '',
  });

  const [createProject, { loading, data }] = useMutation(CREATE_PROJECT);
  const onCreateProject = e => {
    createProject({
      variables: {
        ownerID: '',
        name: state.name,
        description: state.description,
      },
    });
    e.preventDefault();
  };

  return <Layout onCreateProject={onCreateProject} state={state} setState={setState} {...props} />;
};

export default WorkspaceCreate;
