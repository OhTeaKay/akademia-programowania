package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	avg := float64(sum) / float64(len(grades))
	return int(math.Round(avg))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	if len(attendance) == 0 {
		return 0
	}
	present := 0
	for _, attended := range attendance {
		if attended {
			present++
		}
	}
	return math.Round((float64(present)/float64(len(attendance)))*100) / 100
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	avgGrade := AverageGrade(s.Grades)
	attendance := AttendancePercentage(s.Attendance)
	if attendance < 0.6 || avgGrade == 1 || s.Project == 1 {
		return 1
	} else if attendance < 0.8 {
		return int(math.Max(float64(avgGrade+s.Project)/2-1, 1))
	}
	return int(math.Round(float64(avgGrade+s.Project) / 2))
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	grades := make(map[string]uint8)
	for _, s := range students {
		grades[s.Name] = uint8(FinalGrade(s))
	}
	return grades
}
