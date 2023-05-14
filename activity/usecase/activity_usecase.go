package usecase

import (
	"fmt"
	"todolistapi/domain"
	"todolistapi/entity/model"

	"gorm.io/gorm"
)

type activityUsecase struct {
	ActivityRepository domain.ActivityRepository
}

func NewActivityUsecase(activityRepository domain.ActivityRepository) domain.ActivityUsecase {
	return &activityUsecase{
		ActivityRepository: activityRepository,
	}
}

func (u *activityUsecase) GetAllActivity() ([]model.Activity, error) {
	return u.ActivityRepository.GetAllActivity()
}

func (u *activityUsecase) GetOneActivity(activityID uint) (model.Activity, error) {
	return u.ActivityRepository.GetSingleActivity(activityID)
}

func (u *activityUsecase) CreateActivity(activity model.Activity) (model.Activity, error) {
	return u.ActivityRepository.CreateActivity(activity)
}

func (u *activityUsecase) DeleteActivity(activityID uint) error {
	_, err := u.ActivityRepository.GetSingleActivity(activityID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("record not found")
		}
		return err
	}

	err = u.ActivityRepository.DeleteActivity(activityID)
	if err != nil {
		return err
	}

	return nil
}

func (u *activityUsecase) UpdateActivity(activityID uint, title string) (model.Activity, error) {
	activity, err := u.ActivityRepository.GetSingleActivity(activityID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Activity{}, fmt.Errorf("record not found")
		}
		return model.Activity{}, err
	}

	activity.Title = title

	return u.ActivityRepository.UpdateActivity(activity)
}
