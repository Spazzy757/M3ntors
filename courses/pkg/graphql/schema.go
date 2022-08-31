package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/spazzy757/m3ntors/courses/pkg/config"
)

type GraphQLSetup struct {
	Schema graphql.Schema
	Cfg    *config.Config
}

type graphqlSetupOptions func(*GraphQLSetup)

func WithConfig(cfg *config.Config) graphqlSetupOptions {
	return func(g *GraphQLSetup) {
		g.Cfg = cfg
	}
}

func GetGraphQLSetup(opts ...graphqlSetupOptions) *GraphQLSetup {
	g := &GraphQLSetup{}
	for _, opt := range opts {
		opt(g)
	}
	// TODO error handling
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    g.getRootQuery(),
		Mutation: getRootMutation(),
	})
	g.Schema = schema
	return g
}

// getRootQuery resolves the full query schema object
// used for graphql endpoint
func (g *GraphQLSetup) getRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"course":     g.getCourseQuery(),
			"courseList": getCourseListQuery(),
		},
	})
}

// getRootMutation resolves the full mutation schema object
// used for the graphql endpoint
func getRootMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"addCourse":    addCourseMutation(),
			"updateCourse": updateCourseMutation(),
		},
	})
}
