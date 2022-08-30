package courses

import "time"

type Course struct {
	ID        int       `json:"id"`
	Name      string    `json"name"`
	Link      string    `json:"link"`
	Reviewed  bool      `json:"reviewed"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
