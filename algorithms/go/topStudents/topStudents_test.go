package topstudents

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopStudents(t *testing.T) {
	positive_feedback := []string{"smart", "brilliant", "studious"}
	negative_feedback := []string{"not"}
	report := []string{"this student is studious", "the student is smart"}
	student_id := []int{1, 2}
	k := 2
	assert.Equal(t, []int{1, 2}, topStudents(positive_feedback, negative_feedback, report, student_id, k))
}
