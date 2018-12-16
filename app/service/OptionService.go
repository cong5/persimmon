package service

import (
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"time"
)

type OptionService struct{}

func (this *OptionService) GetList(limit int, page int) ([]models.Options, error) {

	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	start := (page - 1) * limit
	optionList := make([]models.Options, 0)
	if err := db.MasterDB.Limit(limit, start).Find(&optionList); err != nil {
		//revel.AppLog.Errorf("Get option failed : %s", err)
		return nil, err
	}

	return optionList, nil
}

func (this *OptionService) GetListPaging(limit int, page int) (*models.PagingContent, error) {
	dataList, err := this.GetList(limit, page)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountOption()
	if cErr != nil {
		return nil, err
	}

	totalPage := utils.GetTotalPage(total, limit)
	pagingContent := &models.PagingContent{Data: dataList,
		Total:       total,
		TotalPage:   totalPage,
		CurrentPage: page}
	return pagingContent, nil
}

func (this *OptionService) CountOption() (int, error) {
	option := new(models.Options)
	total, err := db.MasterDB.Count(option)
	if err != nil {
		revel.AppLog.Errorf("Count option failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *OptionService) GetValueByName(optionName string, real bool) (string, error) {
	option := &models.Options{}
	_, err := db.MasterDB.Where("name = ?", optionName).Cols("value").Get(option)
	if err != nil {
		revel.AppLog.Errorf("Get value by name failed : %s", err)
		return "", err
	}

	return option.Value, nil
}

func (this *OptionService) GetAllOption(real bool) ([]models.Options, error) {
	options := make([]models.Options, 0)
	cKey := utils.CacheKey("OptionService", "AllOption")

	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &options); err != nil {
		dbErr := db.MasterDB.Where("status !='hidden'").Find(&options)
		if dbErr != nil {
			return nil, dbErr
		}
		go cache.Set(cKey, options, 30*time.Minute)
	}

	return options, nil
}

//============================ use backedn ============================

func (this *OptionService) GetOptionById(id int) (*models.Options, error) {
	option := &models.Options{Id: id}
	_, err := db.MasterDB.Get(option)
	if err != nil {
		//revel.AppLog.Errorf("Get option failed : %s", err)
		return nil, err
	}
	return option, nil
}

func (this *OptionService) Save(option models.Options) (int, error) {
	if _, err := db.MasterDB.InsertOne(option); err != nil {
		return 0, err
	}

	this.GetAllOption(true)
	return option.Id, nil
}

func (this *OptionService) Update(id int, option models.Options) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(option)
	if err != nil {
		//revel.AppLog.Errorf("Update option failed: %s", err)
		return false, err
	}

	this.GetAllOption(true)
	return true, nil
}

func (this *OptionService) UpdateByName(optionName string, optionValue string) (bool, error) {
	option := models.Options{Value: optionValue}
	_, err := db.MasterDB.Where("name=?", optionName).Cols("value").Update(option)
	if err != nil {
		//revel.AppLog.Errorf("Update by name option failed: %s", err)
		return false, err
	}

	this.GetAllOption(true)
	return true, nil
}

func (this *OptionService) Destroy(id int, option models.Options) (bool, error) {
	_, err := db.MasterDB.Id(id).Delete(option)
	if err != nil {
		//revel.AppLog.Errorf("Destroy option failed: %s", err)
		return false, err
	}

	this.GetAllOption(true)
	return true, nil
}
