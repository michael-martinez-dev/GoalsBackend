package unit

import (
	"github.com/mixedmachine/GoalsBackend/api/pkg/models"

	"testing"
	
)

func TestMakeSimpleGoal(t *testing.T) {
	userId := "123ABC"
	goal := models.NewUserGoal(userId)

	if goal.Id == "" {
		t.Errorf("Goal Id should not be empty")
	}
	if goal.UserId != userId {
		t.Errorf("Goal title is not correct, got: %s, want: %s.", goal.UserId, userId)
	}
	if goal.Completed != false {
		t.Errorf("Goal completed should be false")
	}
	if goal.Content != "" {
		t.Errorf("Goal content should be empty")
	}
	if goal.Extended != "" {
		t.Errorf("Goal extended should be empty")
	}
}

func TestMakeGoalWithContent(t *testing.T) {
	userId := "123ABC"
	content := "This is the content"
	goal := models.NewUserGoal(userId)
	goal.Content = content

	if goal.Content != content {
		t.Errorf("Goal content is not correct, got: %s, want: %s.", goal.Content, content)
	}
}

func TestMakeGoalWithExtended(t *testing.T) {
	userId := "123ABC"
	extended := "This is the extended"
	goal := models.NewUserGoal(userId)
	goal.Extended = extended

	if goal.Extended != extended {
		t.Errorf("Goal extended is not correct, got: %s, want: %s.", goal.Extended, extended)
	}
}
