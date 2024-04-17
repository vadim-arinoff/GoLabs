// student.go
package student

import (
	"fmt"
)

//Создать пакет student, предоставляющий интерфейс для управления успеваемостью студентов.

// Student представляет информацию о студенте.
type Student struct {
	ID       int
	Name     string
	Grades   map[string]float64 // Карта оценок по предметам
	Courses  []string          // Список курсов, на которых студент обучается
}

// NewStudent создает нового студента с заданным именем и ID.
func NewStudent(id int, name string) *Student {
	return &Student{
		ID:     id,
		Name:   name,
		Grades: make(map[string]float64),
	}
}

// AddGrade добавляет оценку для указанного предмета.
func (s *Student) AddGrade(course string, grade float64) {
	s.Grades[course] = grade
}

// GetAverageGrade возвращает средний балл студента.
func (s *Student) GetAverageGrade() float64 {
	total := 0.0
	count := 0
	for _, grade := range s.Grades {
		total += grade
		count++
	}
	if count == 0 {
		return 0.0
	}
	return total / float64(count)
}

// PrintStudentInfo выводит информацию о студенте.
func (s *Student) PrintStudentInfo() {
	fmt.Printf("Студент %s (ID: %d)\n", s.Name, s.ID)
	fmt.Println("Предметы и оценки:")
	for course, grade := range s.Grades {
		fmt.Printf("%s: %.2f\n", course, grade)
	}
	fmt.Printf("Средний балл: %.2f\n", s.GetAverageGrade())
}
