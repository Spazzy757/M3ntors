package courses

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Course struct {
	ID        int       `json:"id"`
	Name      string    `json"name"`
	Link      string    `json:"link"`
	Reviewed  bool      `json:"reviewed"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CourseHandler struct {
	ctx context.Context
	db  *sql.DB
}

type courseHandlerOptions func(*CourseHandler)

func WithContext(ctx context.Context) courseHandlerOptions {
	return func(ch *CourseHandler) {
		ch.ctx = ctx
	}
}

func WithDB(db *sql.DB) courseHandlerOptions {
	return func(ch *CourseHandler) {
		ch.db = db
	}
}

func NewCourseHandler(opts ...courseHandlerOptions) *CourseHandler {
	ch := &CourseHandler{}
	for _, opt := range opts {
		opt(ch)
	}
	return ch
}

func (ch *CourseHandler) FindByID(id string) (*Course, error) {
	course := new(Course)
	query := "SELECT * FROM courses WHERE id = ?"
	row := ch.db.QueryRowContext(ch.ctx, query, id)
	err := row.Scan(
		&course.ID,
		&course.Name,
		&course.Link,
		&course.Reviewed,
		&course.User,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
	return course, err
}
