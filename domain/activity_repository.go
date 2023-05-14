package domain

import "todolistapi/entity/model"

type ActivityRepository interface {
	GetSingleActivity(activityId uint) (model.Activity, error)
	GetAllActivity() ([]model.Activity, error)
	CreateActivity(activity model.Activity) (model.Activity, error)
	DeleteActivity(activityId uint) error
	UpdateActivity(activity model.Activity) (model.Activity, error)
}
