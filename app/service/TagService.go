package service

import (
	"errors"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"net/url"
	"time"
)

type TagService struct{}

func (this *TagService) GetTagById(id int, real bool) (models.Tags, error) {
	tag := models.Tags{}

	if id <= 0 {
		return tag, errors.New("Param error.")
	}

	cKey := utils.CacheKey("TagService", "InfoById", id)
	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &tag); err != nil {
		_, err := db.MasterDB.Where("id = ?", id).Get(&tag)
		if err != nil {
			return tag, err
		}

		go cache.Set(cKey, tag, 30*time.Minute)
	}

	return tag, nil
}

func (this *TagService) GetTagByIdArr(idArr []int, real bool) ([]models.Tags, error) {
	idLen := len(idArr)
	if idLen <= 0 {
		return nil, errors.New("参数不正确")
	}

	tagArr := make([]models.Tags, idLen)

	for k, postId := range idArr {
		post, err := this.GetTagById(postId, real)
		if err == nil {
			tagArr[k] = post
		}
	}

	return tagArr, nil
}

func (this *TagService) GetTagByName(tagName string) (*models.Tags, error) {
	tag := &models.Tags{}
	if _, err := db.MasterDB.Where("name = ?", tagName).Get(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (this *TagService) GetList(limit int, page int, real bool) ([]models.Tags, error) {

	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	start := (page - 1) * limit
	tagsList := make([]models.Tags, 0)
	err := db.MasterDB.Cols("id").Limit(limit, start).Find(&tagsList)
	if err != nil {
		revel.AppLog.Errorf("Get tag failed : %s", err)
		return nil, err
	}

	idArr := make([]int, len(tagsList))
	for k, v := range tagsList {
		idArr[k] = v.Id
	}

	tagsList, tErr := this.GetTagByIdArr(idArr, real)
	if tErr != nil {
		revel.AppLog.Errorf("GetTagByIdArr failed: %s", tErr)
		return nil, err
	}

	return tagsList, nil
}

func (this *TagService) GetListPaging(limit int, page int, real bool) (*models.PagingContent, error) {
	dataList, err := this.GetList(limit, page, real)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountTags()
	if cErr != nil {
		return nil, cErr
	}

	totalPage := utils.GetTotalPage(total, limit)
	pagingContent := &models.PagingContent{Data: dataList,
		Total:       total,
		TotalPage:   totalPage,
		CurrentPage: page}
	return pagingContent, nil
}

func (this *TagService) CountTags() (int, error) {
	tag := new(models.Tags)
	total, err := db.MasterDB.Count(tag)
	if err != nil {
		//revel.AppLog.Errorf("Count tag failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *TagService) GetListByPostId(postId int, real bool) ([]models.Tags, error) {
	tagsList := make([]models.Tags, 0)
	tagIdArr, err := postTagsService.GetTagIdsByPostId(postId)
	if err != nil {
		revel.AppLog.Errorf("PostTagsService.GetListByPostId Error : %s", err)
		return nil, err
	}

	tagsList, tErr := this.GetTagByIdArr(tagIdArr, real)
	if tErr != nil {
		return nil, err
	}

	return tagsList, nil
}

func (this *TagService) Save(postId int, tags []string) (bool, error) {
	tagsLen := len(tags)
	postTags := make([]models.PostsTags, tagsLen)
	for index := 0; index < tagsLen; index++ {
		//Exist?
		tagsName := tags[index]
		newTag := models.Tags{}
		if _, err := db.MasterDB.Where("name=?", tagsName).Get(&newTag); err != nil {
			revel.AppLog.Errorf("Get tag by tags_name failed : %s", err)
			return false, err
		}

		if newTag.Id <= 0 {
			//Not exist, Insert.
			newTag.Name = tagsName
			newTag.Slug = url.QueryEscape(tagsName)
			if _, insertErr := db.MasterDB.Insert(&newTag); insertErr != nil {
				revel.AppLog.Errorf("Get tag by tags_name failed : %s", insertErr)
				continue
			}
			this.GetTagById(newTag.Id, true)
		}

		//construct post tags map[]
		postTags[index].PostsId = postId
		postTags[index].TagsId = newTag.Id
	}

	//Insert post tags array.
	postTagsService.Save(postId, postTags)

	return true, nil
}

func (this *TagService) SaveOne(tag models.Tags) (int, error) {
	if _, err := db.MasterDB.InsertOne(tag); err != nil {
		//revel.AppLog.Errorf("Get tag by tags_name failed : %s", err)
		return 0, err
	}

	this.GetTagById(tag.Id, true)
	return tag.Id, nil
}

func (this *TagService) Update(id int, tag models.Tags) (bool, error) {
	if _, err := db.MasterDB.Id(id).Update(tag); err != nil {
		//revel.AppLog.Errorf("Update tag failed: %s", err)
		return false, err
	}

	this.GetTagById(id, true)
	return true, nil
}

func (this *TagService) Destroy(id int, tag models.Tags) (bool, error) {
	if _, err := db.MasterDB.Id(id).Delete(tag); err != nil {
		//revel.AppLog.Errorf("Destroy tag failed: %s", err)
		return false, err
	}

	cKey := utils.CacheKey("TagService", "InfoById", id)
	go cache.Delete(cKey)
	return true, nil
}

func (this *TagService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
