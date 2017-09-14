package main

var Schema = `
	schema {
		query: Query
	}

	type Query {
		hello: String!
	}
`

type Resolver struct{}

func (r *Resolver) Hello() string {
	return "serverless-golang graphql"
}
