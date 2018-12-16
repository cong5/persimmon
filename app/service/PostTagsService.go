package service

import (
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
)

type PostTagsService struct{}

func (this *PostTagsService) Save(postId int, postsTags []models.PostsTags) (bool, error) {
	if _, dErr := postTagsService.DeleteByPostID(postId); dErr != nil {
		revel.AppLog.Errorf("postTagsService.DeleteByPostID Error: %s", dErr)
		return false, dErr
	}

	if _, err := db.MasterDB.Insert(&postsTags); err != nil {
		revel.AppLog.Errorf("Save post tag failed: %s", err)
		return false, err
	}

	return true, nil
}

func (this *PostTagsService) DeleteByPostID(postId int) (bool, error) {
	post := models.PostsTags{}
	_, err := db.MasterDB.Id(postId).Delete(post)
	if err != nil {
		revel.AppLog.Errorf("Destroy post tag error: %s", err)
		return false, err
	}
	return true, nil
}

func (this *PostTagsService) GetPostIdsByTagName(tagName string) ([]int, error) {
	tag, tErr := tagService.GetTagByName(tagName)
	if tErr != nil {
		return nil, tErr
	}

	postsTags := make([]models.PostsTags, 0)
	if err := db.MasterDB.Where("tags_id = ?", tag.Id).Find(&postsTags); err != nil {
		revel.AppLog.Errorf("GetPostIdsByTagId error: %s", err)
		return nil, err
	}

	idArr := make([]int, len(postsTags))
	for k, val := range postsTags {
		idArr[k] = val.PostsId
	}

	return idArr, nil
}

func (this *PostTagsService) GetTagIdsByPostId(postId int) ([]int, error) {
	postsTags := make([]models.PostsTags, 0)
	if err := db.MasterDB.Where("posts_id = ?", postId).Find(&postsTags); err != nil {
		revel.AppLog.Errorf("GetTagIdsByPostId error: %s", err)
		return nil, err
	}

	idArr := make([]int, len(postsTags))
	for k, val := range postsTags {
		idArr[k] = val.TagsId
	}

	return idArr, nil
}

func (this *PostTagsService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
