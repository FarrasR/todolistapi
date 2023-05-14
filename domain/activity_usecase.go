package domain

import "todolistapi/entity/model"

type ActivityUsecase interface {
	GetAllActivity() ([]model.Activity, error)
	GetOneActivity(activityID uint) (model.Activity, error)
	CreateActivity(activity model.Activity) (model.Activity, error)
	DeleteActivity(activityID uint) error
	UpdateActivity(activityID uint, title string) (model.Activity, error)
}
