go run github.com/99designs/gqlgen generate

# generiere von graphql.ts für frontend
go run ../Tools/GraphTypeGen/main.go \
  -schema="./graph/schema.graphqls" \
  -output="../Frontend/src/graphql.ts" \
  -header="import { gql } from '@apollo/client/core';\nimport apolloClient from './apolloClient.ts';\n\n" \
  -client="apolloClient" \
  -error="return"