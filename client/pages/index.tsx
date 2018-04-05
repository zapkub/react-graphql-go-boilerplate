import 'isomorphic-fetch'
import * as React from 'react'
import { render } from 'react-dom'
import ApolloClient, { gql } from 'apollo-boost'
import { ApolloProvider, Query } from 'react-apollo'

export default class Index extends React.Component<{}, {}> {
  client: ApolloClient<{}>
  constructor(props) {
    super(props)
    this.client = new ApolloClient({
      uri: 'http://localhost:3000/graphql'
    })
  }
  render() {
    const ApolloApp = AppComponent => (
      <ApolloProvider client={this.client}>
        <Query query={gql`query { hello } `} >
        {
          ({ data }) => (
            <div>{'result >>> ' + ( data.hello  || 'waiting.... ') }</div>
          )
        }
        </Query>
      </ApolloProvider>
    )
    return <ApolloApp />
  }
}
