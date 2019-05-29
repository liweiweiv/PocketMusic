package model

import "time"

type List_songs struct {
	Id        uint      `json:"id"`
	Lid       uint      `json:"lid"`
	Mid       uint      `json:"mid"`
	Status    uint      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
