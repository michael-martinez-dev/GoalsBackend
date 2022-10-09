package models

import (
	"time"

	"github.com/mixedmachine/GoalsBackend/api/pkg/util"
)

type Goal struct {
	Id        string    `json:"id" gorm:"primary_key"`
	UserId    string    `json:"user_id" gorm:"column:user_id;primary_key"`
	Completed bool      `json:"completed" gorm:"column:completed"`
	Content   string    `json:"content" gorm:"column:content"`
	Extended  string    `json:"extended" gorm:"column:extended"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;<-:create"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func NewUserGoal(userId string) *Goal {
	currentTime := time.Now()
	return &Goal{
		Id:        util.GenerateId(),
		UserId:    userId,
		Completed: false,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}
