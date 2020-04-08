import React from 'react';
import gql from 'graphql-tag';
import { useSetState } from 'react-use';
import { useMutation } from '@apollo/react-hooks';

import Loading from '../../../components/loading';

import Layout from './layout';

const CREATE_PROJECT = gql`
  mutation($name: String!, $description: String) {
    ProjectCreate(input: { name: $name, description: $description }) {
      id
      name
      description
      owner {
        id
        fullname
        email
      }
    }
  }
`;

const WorkspaceCreate = (props) => {
  const [state, setState] = useSetState({
    name: '',
    description: '',
  });

  const [createProject, { loading }] = useMutation(CREATE_PROJECT);
  const onCreateProject = (e) => {
    createProject({
      variables: {
        ownerID: '',
        name: state.name,
        description: state.description,
      },
    });
    e.preventDefault();
  };

  if (loading) {
    return <Loading />;
  }

  return <Layout onCreateProject={onCreateProject} state={state} setState={setState} {...props} />;
};

export default WorkspaceCreate;
