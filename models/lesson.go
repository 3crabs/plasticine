package models

import "time"

type Lesson struct {
	Id      int       `json:"id"`
	Place   Place     `json:"place"`
	Weekday string    `json:"weekday"`
	Time    time.Time `json:"time"`
	Subject Subject   `json:"subject"`
	Teacher User      `json:"teacher"`
	GroupId int       `json:"group_id"`
}
