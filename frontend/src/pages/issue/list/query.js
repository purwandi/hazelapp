import gql from 'graphql-tag';

export default gql`
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
