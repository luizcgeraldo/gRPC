package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(name string, description string, CategoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, CategoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryId:  CategoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {

	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}

	for rows.Next() {
		var course Course
		err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryId:  course.CategoryId,
		})
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(CategoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var course Course
		err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryId:  course.CategoryId,
		})
	}
	return courses, nil
}
