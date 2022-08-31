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

// WithContext used to set the context of the course handler
func WithContext(ctx context.Context) courseHandlerOptions {
	return func(ch *CourseHandler) {
		ch.ctx = ctx
	}
}

// WithDB used to set the database connection on the
// course handler
func WithDB(db *sql.DB) courseHandlerOptions {
	return func(ch *CourseHandler) {
		ch.db = db
	}
}

// NewCourseHandler creates a coourse handler instance
func NewCourseHandler(opts ...courseHandlerOptions) *CourseHandler {
	ch := &CourseHandler{}
	for _, opt := range opts {
		opt(ch)
	}
	return ch
}

// FindByID takes an ID and queries db for a single course
func (ch *CourseHandler) FindByID(id string) (*Course, error) {
	course := new(Course)
	query := "SELECT * FROM courses WHERE id = $1"
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
	if err != nil {
		return nil, err
	}
	return course, nil
}

// Create takes a course and inserts it into the database and retruns
// the full course resource
func (ch *CourseHandler) Create(c *Course) (*Course, error) {
	tx, err := ch.db.Begin()
	if err != nil {
		return nil, err
	}
	q := `INSERT INTO courses (name, link, reviewed, user_id) VALUES ($1, $2, $3, $4) RETURNING *;`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(c.Name, c.Link, c.Reviewed, c.User).Scan(
		&c.ID,
		&c.Name,
		&c.Link,
		&c.Reviewed,
		&c.User,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return c, nil
}
