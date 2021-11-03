package models

import "time"

type Lesson struct {
	Weekday string    `json:"weekday"`
	Time    time.Time `json:"time"`
	Subject Subject   `json:"subject"`
	Teacher User      `json:"teacher"`
	GroupId string    `json:"group_id"`
}
