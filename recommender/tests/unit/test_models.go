package unit

import (
	"github.com/mixedmachine/GoalsBackend/recommender/pkg/models"

	"testing"
)

func TestMakeSimpleGoal(t *testing.T) {
	goal := models.Goal{}

	if goal.Id != "" {
		t.Errorf("Goal Id should be empty")
	}
	if goal.UserId != "" {
		t.Errorf("Goal UserId should be empty")
	}
	if goal.Completed {
		t.Errorf("Goal completed should be false")
	}
	if goal.Content != "" {
		t.Errorf("Goal content should be empty")
	}
	if goal.Extended != "" {
		t.Errorf("Goal extended should be empty")
	}
}

func TestSetExtended(t *testing.T) {
	goal := models.Goal{}
	goal.SetExtended("extended")

	if goal.Extended != "extended" {
		t.Errorf("Goal extended should be 'extended'")
	}
}

func TestMakeSimpleMessage(t *testing.T) {
	message := models.Message{}

	if message.ReturnEndpoint != "" {
		t.Errorf("Message ReturnEndpoint should be empty")
	}
	if message.Goal.Id != "" {
		t.Errorf("Message Goal Id should be empty")
	}
	if message.Goal.UserId != "" {
		t.Errorf("Message Goal UserId should be empty")
	}
	if message.Goal.Completed {
		t.Errorf("Message Goal completed should be false")
	}
	if message.Goal.Content != "" {
		t.Errorf("Message Goal content should be empty")
	}
	if message.Goal.Extended != "" {
		t.Errorf("Message Goal extended should be empty")
	}
}
