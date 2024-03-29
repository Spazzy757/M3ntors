package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/spazzy757/m3ntors/courses/pkg/courses"
	"github.com/spazzy757/m3ntors/courses/pkg/middleware"
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
func (q *GraphQLSetup) getCourseQuery() *graphql.Field {
	return &graphql.Field{
		Type:        courseType,
		Description: "get a single course",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(string)
			ch := courses.NewCourseHandler(
				courses.WithDB(q.Cfg.DB),
			)
			course, err := ch.FindByID(p.Context, id)
			//TODO if nothing is returned return a "NOT FOUND" error
			// else just a generic something went wrong error
			return course, err
		},
	}
}

// getCourseListQuery resolver for listing courses
func getCourseListQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(courseType),
		Description: "list of courses",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//TODO return list of courses
			return nil, nil
		},
	}
}

// addCourseMutation mutation to create a new course
func (q *GraphQLSetup) addCourseMutation() *graphql.Field {
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
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			u := middleware.GetUserFromContext(p.Context)
			if u == nil {
				return nil, fmt.Errorf(
					"authentication: %s",
					"could not find user, are you loged in?",
				)
			}
			name, _ := p.Args["name"].(string)
			link, _ := p.Args["link"].(string)
			ch := courses.NewCourseHandler(
				courses.WithDB(q.Cfg.DB),
			)
			c, err := ch.Create(p.Context, &courses.Course{
				Name:     name,
				Link:     link,
				Reviewed: false,
				User:     u.ID,
			})
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

// updateCourseMutation mutation to update an existing course
func updateCourseMutation() *graphql.Field {
	return &graphql.Field{
		Type:        courseType, // the return type for this field
		Description: "Update existing course",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
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
