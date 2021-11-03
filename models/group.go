package models

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GroupInfo struct {
	*Group
	Lessons []Lesson `json:"lessons"`
}
