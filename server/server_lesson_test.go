package server

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"plasticine/db"
	"plasticine/models"
	"strconv"
	"strings"
	"testing"
	"time"
)

func (s *server) addLessonReq(lesson models.Lesson) {
	bytes, _ := json.Marshal(lesson)
	_, c := s.post(strings.NewReader(string(bytes)))
	_ = s.addLesson(c)
}

func (s *server) getLessonsReq(groupId int) (*httptest.ResponseRecorder, []models.Lesson) {
	rec, c := s.get()
	c.SetParamNames("groupId")
	c.SetParamValues(strconv.Itoa(groupId))
	_ = s.getLessons(c)
	var lessons []models.Lesson
	_ = json.Unmarshal([]byte(rec.Body.String()), &lessons)
	return rec, lessons
}

func (s *server) updateLessonReq(lessonId int, lesson models.Lesson) {
	bytes, _ := json.Marshal(lesson)
	_, c := s.put(strings.NewReader(string(bytes)))
	c.SetParamNames("lessonId")
	c.SetParamValues(strconv.Itoa(lessonId))
	_ = s.updateLesson(c)
}

func (s *server) deleteLessonReq(lessonId int) {
	_, c := s.delete()
	c.SetParamNames("lessonId")
	c.SetParamValues(strconv.Itoa(lessonId))
	_ = s.deleteLesson(c)
}

func TestGetLessons(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	rec, lessons := s.getLessonsReq(0)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(lessons))
}

func TestAddLesson(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	place := models.Place{Name: "place name"}
	s.addPlaceReq(place)
	rec, places := s.getPlacesReq()
	place = places[0]

	subject := models.Subject{Name: "subject name"}
	s.addSubjectReq(subject)
	rec, subjects := s.getSubjectsReq()
	subject = subjects[0]

	teacher := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Teacher,
	}
	s.addUserReq(teacher)
	rec, teachers := s.getTeachersReq()
	teacher = teachers[0]

	group := models.Group{Name: "group name"}
	s.addGroupsReq(group)
	rec, groups := s.getGroupsReq()
	group = groups[0]

	lesson := models.Lesson{
		Place:   place,
		Weekday: time.Monday.String(),
		Time:    time.Now(),
		Subject: subject,
		Teacher: teacher,
		GroupId: group.Id,
	}
	s.addLessonReq(lesson)

	rec, lessons := s.getLessonsReq(group.Id)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(lessons))
	assert.Equal(t, lesson.Place.Name, place.Name)
	assert.Equal(t, lesson.Subject.Name, subject.Name)
	assert.Equal(t, lesson.Teacher.FirstName, teacher.FirstName)
	assert.Equal(t, lesson.Teacher.SecondName, teacher.SecondName)
	assert.Equal(t, lesson.GroupId, group.Id)
}

func TestUpdateLesson(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	place := models.Place{Name: "place name"}
	s.addPlaceReq(place)
	_, places := s.getPlacesReq()
	place = places[0]

	subject := models.Subject{Name: "subject name"}
	s.addSubjectReq(subject)
	_, subjects := s.getSubjectsReq()
	subject = subjects[0]

	teacher := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Teacher,
	}
	s.addUserReq(teacher)
	_, teachers := s.getTeachersReq()
	teacher = teachers[0]

	group := models.Group{Name: "group name"}
	s.addGroupsReq(group)
	_, groups := s.getGroupsReq()
	group = groups[0]

	lesson := models.Lesson{
		Place:   place,
		Weekday: time.Monday.String(),
		Time:    time.Now(),
		Subject: subject,
		Teacher: teacher,
		GroupId: group.Id,
	}
	s.addLessonReq(lesson)
	_, lessons := s.getLessonsReq(group.Id)
	lesson = lessons[0]

	assert.Equal(t, lessons[0].Weekday, time.Monday.String())

	lesson.Weekday = time.Friday.String()
	s.updateLessonReq(lesson.Id, lesson)

	_, lessons = s.getLessonsReq(group.Id)

	assert.Equal(t, lesson.Weekday, time.Friday.String())
}

func TestDeleteLesson(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	place := models.Place{Name: "place name"}
	s.addPlaceReq(place)
	_, places := s.getPlacesReq()
	place = places[0]

	subject := models.Subject{Name: "subject name"}
	s.addSubjectReq(subject)
	_, subjects := s.getSubjectsReq()
	subject = subjects[0]

	teacher := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Teacher,
	}
	s.addUserReq(teacher)
	_, teachers := s.getTeachersReq()
	teacher = teachers[0]

	group := models.Group{Name: "group name"}
	s.addGroupsReq(group)
	_, groups := s.getGroupsReq()
	group = groups[0]

	lesson := models.Lesson{
		Place:   place,
		Weekday: time.Monday.String(),
		Time:    time.Now(),
		Subject: subject,
		Teacher: teacher,
		GroupId: group.Id,
	}
	s.addLessonReq(lesson)

	_, lessons := s.getLessonsReq(group.Id)
	assert.Equal(t, 1, len(lessons))

	s.deleteLessonReq(lessons[0].Id)

	_, lessons = s.getLessonsReq(group.Id)
	assert.Equal(t, 0, len(lessons))
}
