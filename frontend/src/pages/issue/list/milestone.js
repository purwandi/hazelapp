import React from 'react'
import { useApolloClient, useMutation } from 'react-apollo'
import { useRouteMatch } from 'react-router'
import { object } from 'prop-types'
import { GET_ISSUES, CREATE_MILESTONE } from '../query'

const MilestoneCreate = ({ project }) => {
  const client = useApolloClient()
  const match = useRouteMatch()

  const store = client.readQuery({
    query: GET_ISSUES,
    variables: {
      owner: match.params.owner,
      name: match.params.name,
      first: 20
    }
  })

  const [createMilestone] = useMutation(CREATE_MILESTONE)

  const onCreateMilestone = (e) => {
    e.preventDefault()
    const name = `Milestone ${store.project.milestones.length + 1}`
    createMilestone({
      variables: {
        name,
        projectId: project.id
      },
      optimisticResponse: {
        __typename: 'Mutation',
        MilestoneCreate: {
          id: '',
          name,
          __typename: 'Milestone',
        }
      },
      update: (proxy, { data }) => {
        // Read the data from our cache for this query.
        const storeProxy = proxy.readQuery({
          query: GET_ISSUES,
          variables: {
            owner: match.params.owner,
            name: match.params.name,
            first: 20,
          }
        })

        storeProxy.project.milestones.push(data.MilestoneCreate)

        proxy.writeQuery({
          query: GET_ISSUES,
          data: storeProxy
        })
      }
    })
  }

  return (
    <button
      type="button"
      onClick={onCreateMilestone}
      className="px-2 py-1 bg-gray-300 rounded text-gray-700">
      Create sprint
    </button>
  )
}

MilestoneCreate.propTypes = {
  project: object.isRequired
}

export default MilestoneCreate
