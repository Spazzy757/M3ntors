package graphql

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/graphql-go/graphql"
	"github.com/spazzy757/m3ntors/courses/pkg/config"
	"github.com/stretchr/testify/require"
)

var courseQuery = `query course($courseid: String!) {
    course(id:$courseid) {
			id
			name
			link
			reviewed
			user
			createdAt
			updatedAt
    }
}`

var addCourseMutation = `mutation RootMutation(
		$user: String!, 
		$name: String!, 
		$link: String!
	) 
	{
		addCourse(user: $user, name: $name, link: $link) {
		  name
		}
}`

func TestCoursesSchema(t *testing.T) {
	assert := require.New(t)
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherEqual,
		),
	)

	defer db.Close()
	graphqlSetup := GetGraphQLSetup(
		WithConfig(&config.Config{
			DB: db,
		}),
	)

	assert.NoError(err)
	t.Run("course resolver returns a course", func(t *testing.T) {
		vars := make(map[string]interface{})
		vars["courseid"] = "123"
		query := "SELECT * FROM courses WHERE id = $1"
		timestamp := time.Now()
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"link",
			"reviewed",
			"user",
			"created_at",
			"updated_at",
		}).AddRow(
			123,
			"Foo",
			"bar.com",
			false,
			"baz123",
			timestamp,
			timestamp,
		)
		mock.ExpectQuery(query).WithArgs("123").WillReturnRows(rows)
		params := graphql.Params{
			Schema:         graphqlSetup.Schema,
			VariableValues: vars,
			RequestString:  courseQuery,
		}
		r := graphql.Do(params)
		res, err := json.Marshal(r)
		assert.NoError(err)
		err = mock.ExpectationsWereMet()
		assert.NoError(err)
		assert.NotContains(fmt.Sprintf("%s", res), "error")
		assert.Contains(fmt.Sprintf("%s", res), `"id":123`)
	})
	t.Run("course resolver returns not found", func(t *testing.T) {
		vars := make(map[string]interface{})
		vars["courseid"] = "123"
		query := "SELECT * FROM courses WHERE id = $1"
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"link",
			"reviewed",
			"user",
			"created_at",
			"updated_at",
		})
		mock.ExpectQuery(query).WithArgs("123").WillReturnRows(rows)
		params := graphql.Params{
			Schema:         graphqlSetup.Schema,
			VariableValues: vars,
			RequestString:  courseQuery,
		}
		r := graphql.Do(params)
		res, err := json.Marshal(r)
		assert.NoError(err)
		err = mock.ExpectationsWereMet()
		assert.NoError(err)
		assert.Contains(fmt.Sprintf("%s", res), "error")
		assert.Contains(fmt.Sprintf("%s", res), `"course":null`)
	})
	t.Run("addCourse mutation returns a course", func(t *testing.T) {
		vars := make(map[string]interface{})
		vars["name"] = "Foo"
		vars["link"] = "bar.com"
		vars["user"] = "user123"
		q := `INSERT INTO courses (name, link, reviewed, user_id) VALUES ($1, $2, $3, $4) RETURNING *;`
		timestamp := time.Now()
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"link",
			"reviewed",
			"user",
			"created_at",
			"updated_at",
		}).AddRow(
			123,
			"Foo",
			"bar.com",
			false,
			"baz123",
			timestamp,
			timestamp,
		)
		mock.ExpectBegin()
		mock.ExpectPrepare(q)
		mock.ExpectQuery(q).WithArgs("Foo", "bar.com", false, "user123").WillReturnRows(rows)
		mock.ExpectCommit()
		params := graphql.Params{
			Schema:         graphqlSetup.Schema,
			VariableValues: vars,
			RequestString:  addCourseMutation,
		}
		r := graphql.Do(params)
		res, err := json.Marshal(r)
		assert.NoError(err)
		err = mock.ExpectationsWereMet()
		assert.NoError(err)
		assert.NotContains(fmt.Sprintf("%s", res), "error")
		assert.Contains(fmt.Sprintf("%s", res), `"name":"Foo"`)
	})
}
