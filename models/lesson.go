package models

import "time"

type Lesson struct {
	Place   Place     `json:"place"`
	Weekday string    `json:"weekday"`
	Time    time.Time `json:"time"`
	Subject Subject   `json:"subject"`
	Teacher User      `json:"teacher"`
	GroupId string    `json:"group_id"`
}
