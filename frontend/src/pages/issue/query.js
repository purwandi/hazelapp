import gql from 'graphql-tag';

export const GET_ISSUES = gql`
  query($owner: String!, $name: String!, $first: Int) {
    project(owner: $owner, name: $name) {
      id
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
          issueType
          milestone {
            id
          }
        }
      }
    }
  }
`;

export const CREATE_ISSUE = gql`
  mutation(
    $projectId: ID!
    $milestoneId: String
    $title: String!
    $body: String
    $issueType: String!
  ) {
    IssueCreate(
      input: {
        projectId: $projectId
        milestoneId: $milestoneId
        title: $title
        body: $body
        issueType: $issueType
      }
    ) {
      id
      number
      title
      body
      issueType
      milestone {
        id
      }
    }
  }
`;

export const CREATE_MILESTONE = gql`
  mutation($name: String!, $projectId: String!) {
    MilestoneCreate(input: { name: $name, projectId: $projectId }) {
      id
      name
    }
  }
`;
