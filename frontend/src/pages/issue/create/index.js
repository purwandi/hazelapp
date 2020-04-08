import React from 'react';
import { object } from 'prop-types';
import { useSetState } from 'react-use';
import { useMutation } from '@apollo/react-hooks';
import { useRouteMatch } from 'react-router';
import { GET_ISSUES, CREATE_ISSUE } from '../query';
import Layout from './layout';

const IssueCreate = ({ milestone, project }) => {
  const [state, setState] = useSetState({
    title: '',
    type: 'story',
  });

  const match = useRouteMatch();
  const [issueCreate] = useMutation(CREATE_ISSUE);
  const onFormSubmit = (e) => {
    e.preventDefault();

    // perform issue create mutation
    issueCreate({
      variables: {
        projectId: project.id,
        milestoneId: milestone.id,
        title: state.title,
        issueType: state.type,
      },
      optimisticResponse: {
        __typename: 'Mutation',
        IssueCreate: {
          __typename: 'Issue',
          milestone: {
            id: milestone.id,
            name: milestone.name,
            __typename: 'Milestone',
          },
          id: '',
          number: '',
          body: '',
          project,
          title: state.title,
          issueType: state.type,
        },
      },
      update: (proxy, { data }) => {
        // Read the data from our cache for this query.
        const store = proxy.readQuery({
          query: GET_ISSUES,
          variables: {
            owner: match.params.owner,
            name: match.params.name,
            first: 20,
          },
        });

        // push the issues data nodes
        store.project.issues.nodes.push(data.IssueCreate);

        // write our data back to the cache
        proxy.writeQuery({
          query: GET_ISSUES,
          variables: {
            owner: match.params.owner,
            name: match.params.name,
            first: 20,
          },
          data: store,
        });
      },
    });

    // reset form state
    setState({
      title: '',
      type: 'story',
    });
  };

  const onChangeType = (value) => {
    setState({ type: value });
  };

  return (
    <Layout
      state={state}
      setState={setState}
      onFormSubmit={onFormSubmit}
      onChangeType={onChangeType}
    />
  );
};

IssueCreate.propTypes = {
  project: object.isRequired,
  milestone: object.isRequired,
};

export default IssueCreate;
