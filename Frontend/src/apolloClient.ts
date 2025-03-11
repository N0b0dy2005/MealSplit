import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from "@apollo/client/core";

const defaultOptions = {
    watchQuery: {
        fetchPolicy: 'no-cache',
        errorPolicy: 'ignore',
    },
    query: {
        fetchPolicy: 'no-cache',
        errorPolicy: 'all',
    },
};

const createApolloClient = () => {
    let baseUrl = "./";

    if (baseUrl[baseUrl.length - 1] === '/') {
        baseUrl = baseUrl.slice(0, -1);
    }

    const httpLink = new HttpLink({
        uri: baseUrl + '/api/query',
        credentials: "include",
    });

    const authLink = new ApolloLink((operation, forward) => {
        const token = window.global?._jwt;
        operation.setContext({
            headers: {
                Authorization: `Bearer ${token}`,
            }
        });
        return forward(operation);
    });

    return new ApolloClient({
        link: authLink.concat(httpLink),
        cache: new InMemoryCache(),
        defaultOptions,
    });
};

// Exportiere direkt eine Instanz des Apollo Clients
const client = createApolloClient();
export default client;