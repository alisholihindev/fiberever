package dokter

import (
	"fiberever/model"
	"fmt"

	"github.com/alisholihindev/go-lib"
)

func Get(id int) (model.Dokter, error) {
	var model model.Dokter
	if err := lib.DBConnection().First(&model, "id = ?", id).Error; err != nil {
		lib.CommonLogger().Error(err)
		return model, err
	}
	return model, nil
}

func Add(a model.Dokter) (model.Dokter, error) {
	if err := lib.DBConnection().Create(&a).Error; err != nil {
		lib.CommonLogger().Error(err)
		return a, err
	}
	return a, nil
}

func GetAll(q model.Dokter, p int, l int, o string) (*lib.Paginator, error) {
	var model []model.Dokter
	paginator := lib.Paging(&lib.PaginationParam{
		DB:      lib.DBConnection().Find(&model),
		Page:    p,
		Limit:   l,
		OrderBy: []string{o},
		ShowSQL: true,
	}, &model)
	return paginator, nil
}

func Update(id int, param *model.Dokter) (model.Dokter, error) {
	var a model.Dokter
	if err := lib.DBConnection().First(&a, "id = ? ", id).Updates(&param).Error; err != nil {
		lib.CommonLogger().Error(err)
		return a, err
	}

	return Get(id)
}

func Delete(id int) (bool, error) {
	fmt.Print(id)
	var a model.Dokter
	if err := lib.DBConnection().Delete(&a, id).Error; err != nil {
		lib.CommonLogger().Error(err)
		return false, err
	}

	return true, nil
}
