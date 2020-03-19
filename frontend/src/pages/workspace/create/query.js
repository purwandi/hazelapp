import gql from 'graphql-tag'

export const CreateProject = gql`
  mutation CreateProject($ownerID: ID, $name: String!, $description: String ) {
    CreateProject(ownerID: $ownerID, name: $name, description: $description) {
      id
      name
      description
    }
  }
`
