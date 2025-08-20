package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"Name"`
	Email string  `json:"Email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.5
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("GPA must be between 0-4")
	}
	return nil
}

func main() {
	//var st Student = Student{ID: "1", Name: "Fasia", Email: "minaiad_f@silpakorn.edu", Year: 3, GPA: 3.89}
	students := []Student{
		{ID: "1", Name: "Fasia", Email: "minaiad_f@silpakorn.edu", Year: 3, GPA: 3.89},
		{ID: "2", Name: "Prae", Email: "preatm@gmail.com", Year: 3, GPA: 3.95},
	}
	newStudent := Student{ID: "3", Name: "Yok", Email: "yok@gmail.com", Year: 3, GPA: 3.99}
	students = append(students, newStudent)
	for i, student := range students {
		fmt.Printf("%d Honor = %v\n", i, student.IsHonor())
		fmt.Printf("Validation = %v\n", student.Validate())
	}
	fmt.Println()
}
