package courses

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spazzy757/m3ntors/courses/pkg/middleware"
	"github.com/stretchr/testify/require"
)

func TestCoursesHandler(t *testing.T) {
	assert := require.New(t)
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherEqual,
		),
	)
	defer db.Close()
	assert.NoError(err)
	ctx := context.TODO()
	ctx = context.WithValue(ctx, middleware.UserKey, "foobar")
	h := NewCourseHandler(
		WithContext(ctx),
		WithDB(db),
	)
	timestamp := time.Now()
	columns := []string{
		"id",
		"name",
		"link",
		"reviewed",
		"user",
		"created_at",
		"updated_at",
	}
	rows := sqlmock.NewRows(columns).AddRow(
		123,
		"Foo",
		"bar.com",
		false,
		"baz123",
		timestamp,
		timestamp,
	)

	t.Run("create a course", func(t *testing.T) {
		vars := make(map[string]interface{})
		vars["name"] = "Foo"
		vars["link"] = "bar.com"
		vars["user"] = "user123"
		q := `INSERT INTO courses (name, link, reviewed, user_id) VALUES ($1, $2, $3, $4) RETURNING *;`
		mock.ExpectQuery(q).WithArgs(
			"Foo",
			"bar.com",
			false,
			"user123",
		).WillReturnRows(rows)
		h.Create(&Course{
			Name:     "Foo",
			Link:     "bar.com",
			Reviewed: false,
			User:     "user123",
		})
		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}
