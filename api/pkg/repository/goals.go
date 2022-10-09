package repository

import (
	"log"

	"github.com/mixedmachine/GoalsBackend/api/pkg/db"
	"github.com/mixedmachine/GoalsBackend/api/pkg/models"
	"github.com/mixedmachine/GoalsBackend/api/pkg/util"
	"gorm.io/gorm"
)

type GoalsRepository interface {
	Create(goal *models.Goal) error
	Update(goal *models.Goal) error
	RetrieveById(user, id string) (goal *models.Goal, err error)
	RetrieveAll(user string) (goals []*models.Goal, err error)
	Delete(user, id string) error
}

type goalsRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn db.SqlConnection) GoalsRepository {
	return &goalsRepository{
		db: conn.DB(),
	}
}

func (r *goalsRepository) Create(goal *models.Goal) error {
	result := r.db.Create(goal)
	if result.Error != nil {
		return util.ErrUnableToCreate
	}
	log.Printf("Created goal with primary key: %v\n", goal.Id)
	return nil
}

func (r *goalsRepository) Update(goal *models.Goal) error {
	res := r.db.Save(goal)
	if res.Error != nil {
		return util.ErrUnableToUpdate
	}
	return nil
}

func (r *goalsRepository) RetrieveById(user, id string) (goal *models.Goal, err error) {
	res := r.db.Where("user_id = ? AND id = ?", user, id).First(&goal)
	if res.Error != nil {
		return nil, util.ErrNotFound
	}
	return goal, nil
}

func (r *goalsRepository) RetrieveAll(user string) (goals []*models.Goal, err error) {
	r.db.Where("user_id = ?", user).Find(&goals)
	return goals, nil
}

func (r *goalsRepository) Delete(user, id string) error {
	res := r.db.Where("user_id = ? AND id = ?", user, id).Delete(&models.Goal{})
	if res.Error != nil {
		return util.ErrUnableToUpdate
	}
	return nil
}
