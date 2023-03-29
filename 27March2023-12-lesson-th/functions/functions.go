package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type StudentData interface{
	GetStudents() []Student
	GetStudentsById(id int) error
	AddStudent(student Student) error
	RemoveStudentById(id int) error
	UpdateStudentById(id int)
}

type Subjects struct{
	Id 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Marks 	int 	`json:"marks"`
}

type Student struct{
	Id 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Grade 	int 	`json:"grade"`
	Section string 	`json:"section"`
	Course 	string 	`json:"course"`
	Subjects []Subjects
}

type StudentLibrary struct{
	Student []Student
}

func (s *StudentLibrary) GetStudents() []Student {
	return s.Student
}

func (s *StudentLibrary) GetStudentsById(id int) error {
	for _, student := range s.Student {
		if student.Id == id {
			st := &student
			fmt.Println(st)
		}
	}
	return nil
}

func (s *StudentLibrary) AddStudent(student Student) error {
	s.Student = append(s.Student, student)
	return nil
}

// func (s *StudentLibrary) RemoveStudentById(id int) error{

// }

// func (s *StudentLibrary) UpdateStudentById(id int) {
// 	for _, student := range s.Student{
// 		if student.Id == id {
// 			student.Name = name,
// 		}
// 	}
// }

func Helper(){

	fmt.Println("Welcome!")

	data, err := ioutil.ReadFile("student.json")
	if err != nil {
		fmt.Println("Error reading JSON: ", err)
		return
	}

	var StudentList []Student
	err = json.Unmarshal(data, &StudentList)
	if err != nil {
		fmt.Println("Error Unmarshling the JSON: ", err)
		return
	}

	library := &StudentLibrary{Student: StudentList}

	fmt.Println(" - method GetStudents() - ")
	fmt.Println(library.GetStudents())

	fmt.Println(" - method GetStudentsById() - ")
	fmt.Println(library.GetStudentsById(2))

	fmt.Println(" - method AddStudent() - ")
	library.AddStudent(Student{
		Id: 5,	
		Name: "Amirbek",	
		Grade: 12,
		Section: "C",
		Course: "Science",
		Subjects: []Subjects{
			{
				Id: 1,
				Name: "Physics",
				Marks: 78,
			},
			{
				Id: 2,
				Name: "Software",
				Marks: 76,
			},
			{
				Id: 3,
				Name: "Math",
				Marks: 85,
			},
		},
	})
	fmt.Println(library.GetStudents())
}