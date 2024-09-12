package models

import (
	"time"

	"github.com/google/uuid"
)

// ReviewCreate struct to describe how to create a Review.
type ReviewCreate struct {
	Title   string `db:"title" json:"title" validate:"required,lte=15"`
	Content string `db:"content" json:"content" validate:"required,lte=255"`
}

// Review struct to describe Review object.
type Review struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title     string    `db:"title" json:"title" validate:"required,lte=15"`
	Content   string    `db:"content" json:"content" validate:"required,lte=255"`
	Likes     int       `db:"likes" json:"likes"`
}
