import { ApolloClient } from 'apollo-client';
import { ApolloLink } from 'apollo-link';
import { HttpLink } from 'apollo-link-http';
import { onError } from 'apollo-link-error';
import { setContext } from 'apollo-link-context';
import { InMemoryCache } from 'apollo-cache-inmemory';

export const CreateApolloClient = (context, dispatch, router) => {
  const authLink = () => {
    return setContext((_, { headers }) => {
      // get the authentication token from local storage if it exists
      const token = context.user.token || '';
      // return the headers to the context so httpLink can read them
      return {
        headers: {
          authorization: `Bearer ${token}`,
          ...headers,
        },
      };
    });
  };

  const httpLink = new HttpLink({
    uri: `${process.env.REACT_APP_BACKEND_URL}/graphql`,
    headers: {
      Accept: 'application/json',
    },
  });

  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors) {
      graphQLErrors.map(({ message, location, path }) => {
        console.log(`
        [GraphQL error]:
          Messages : ${message},
          Location: ${location},
          Path: ${path}
        `);
        return false;
      });
    }

    if (networkError) {
      const UnAuthorizedError = [400, 401];
      if (UnAuthorizedError.includes(networkError.statusCode)) {
        dispatch({ type: 'LOGOUT' });
        router.history.push('/auth');
      }
      // eslint-disable-next-line no-console
      console.log(`[Network error]: ${networkError}`);
    }
  });

  return new ApolloClient({
    link: ApolloLink.from([authLink(), errorLink, httpLink]),
    cache: new InMemoryCache(),
  });
};
