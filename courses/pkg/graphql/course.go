package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/spazzy757/m3ntors/courses/pkg/courses"
)

// Structure of a course for Graphql
var courseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Course",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"link": &graphql.Field{
			Type: graphql.String,
		},
		"reviewed": &graphql.Field{
			Type: graphql.Boolean,
		},
		"user": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// getCourseQuery resolver handle getting a single
// course by id
func getCourseQuery() *graphql.Field {
	return &graphql.Field{
		Type:        courseType,
		Description: "get a single course",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	}
}

// getCourseListQuery resolver for listing courses
func getCourseListQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(courseType),
		Description: "List of courses",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	}
}

// addCourseMutation mutation to create a new course
func addCourseMutation() *graphql.Field {
	return &graphql.Field{
		Type:        courseType,
		Description: "add a new course",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"link": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			name, _ := params.Args["name"].(string)
			link, _ := params.Args["link"].(string)
			// perform mutation operation here
			// TODO create a course and save to DB.
			newCourse := courses.Course{
				Name:     name,
				Link:     link,
				Reviewed: false,
			}
			return newCourse, nil
		},
	}
}

// updateCourseMutation mutation to update an existing course
func updateCourseMutation() *graphql.Field {
	return &graphql.Field{
		Type:        courseType, // the return type for this field
		Description: "Update existing course",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"link": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// TODO find and update course
			updatedCourse := courses.Course{}
			// Return affected beast
			return updatedCourse, nil
		},
	}
}
