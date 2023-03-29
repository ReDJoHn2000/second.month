package json

import "gitlab.com/i.abdukhoshimov/golang_bootcamp/new-project/storage/repo"

// Define Get all Students
func (s *repo.School) GetStudents() []Student {
}

// Get Student by its id
func (s *repo.School) GetStudentByID(id int) (*Student, error) {
}

// Inserting the New Student
func (s *repo.School) AddStudent(student Student) error {
}

// Updating the Students info
func (s *repo.School) Update(student repo.Student) error {
}

// Deleting the Student from the list
func (s *repo.School) RemoveStudentByID(id int) error {
}
