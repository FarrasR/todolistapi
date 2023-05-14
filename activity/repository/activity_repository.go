package repository

import (
	"todolistapi/domain"
	"todolistapi/entity/model"

	"gorm.io/gorm"
)

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) domain.ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

func (r *activityRepository) GetSingleActivity(activityId uint) (model.Activity, error) {
	var activity model.Activity

	result := r.db.Where(model.Activity{ActivityID: activityId}).First(&activity)

	if result.Error != nil {
		return model.Activity{}, result.Error
	}
	return activity, nil
}

func (r *activityRepository) GetAllActivity() ([]model.Activity, error) {
	var activities []model.Activity

	result := r.db.Find(&activities)

	if result.Error != nil {
		return []model.Activity{}, result.Error
	}
	return activities, nil
}

func (r *activityRepository) CreateActivity(activity model.Activity) (model.Activity, error) {

	result := r.db.Create(&activity)

	if result.Error != nil {
		return model.Activity{}, result.Error
	}

	return activity, nil
}

func (r *activityRepository) DeleteActivity(activityId uint) error {
	result := r.db.Delete(&model.Activity{}, activityId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *activityRepository) UpdateActivity(activity model.Activity) (model.Activity, error) {
	result := r.db.Save(&activity)
	if result.Error != nil {
		return model.Activity{}, result.Error
	}
	return activity, nil
}
