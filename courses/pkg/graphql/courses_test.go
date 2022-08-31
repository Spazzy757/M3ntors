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
}
