package graphql

import "github.com/graphql-go/graphql"

func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    getRootQuery(),
		Mutation: getRootMutation(),
	})
}

// getRootQuery resolves the full query schema object
// used for graphql endpoint
func getRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"course":     getCourseQuery(),
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
