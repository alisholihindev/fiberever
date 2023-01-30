package user

import (
	"fiberever/model"

	"github.com/alisholihindev/go-lib"
)

func GetAll(q model.User, p int, l int, o string) (*lib.Paginator, error) {
	var model []model.User
	paginator := lib.Paging(&lib.PaginationParam{
		DB:      lib.DBConnection().Find(&model),
		Page:    p,
		Limit:   l,
		OrderBy: []string{o},
		ShowSQL: true,
	}, &model)
	return paginator, nil
}

func Get(id int) (model.User, error) {
	var model model.User
	if err := lib.DBConnection().First(&model, "id = ?", id).Error; err != nil {
		lib.CommonLogger().Error(err)
		return model, err
	}
	return model, nil
}

func Add(a model.User) (model.User, error) {
	if err := lib.DBConnection().Create(&a).Error; err != nil {
		lib.CommonLogger().Error(err)
		return a, err
	}
	return a, nil
}
