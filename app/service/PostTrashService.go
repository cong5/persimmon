package service

import (
	"fmt"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"qiniupkg.com/x/errors.v7"
)

type PostTrashService struct{}

func (this *PostTrashService) GetTrashList(categoryId int, keywords string, limit int, page int, real bool) ([]models.Posts, error) {

	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	start := (page - 1) * limit
	postList := make([]models.Posts, 0)
	if err := db.MasterDB.Where("`deleted_at` IS NOT NULL").Cols("id").Limit(limit, start).Find(&postList); err != nil {
		revel.AppLog.Errorf("Get trash list failed : %s", err)
		return nil, err
	}

	idArr := make([]int, len(postList))
	for k, v := range postList {
		idArr[k] = v.Id
	}

	if len(idArr) <= 0 {
		return make([]models.Posts, 0), nil
	}

	postList, pErr := postService.GetPostByIdArr(idArr, real)
	if pErr != nil {
		revel.AppLog.Errorf("postService.GetPostByIdArr Error : %s", pErr)
		return nil, pErr
	}

	return postList, nil
}

func (this *PostTrashService) GetTrashListPaging(categoryId int, keywords string, limit int, page int, real bool) (*models.PagingContent, error) {
	dataList, err := this.GetTrashList(categoryId, keywords, limit, page, real)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountTrashPost()
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

func (this *PostTrashService) Restore(ids []int) (bool, error) {
	if len(ids) < 1 {
		return false, errors.New("params error.")
	}
	idString := utils.IntJoin(ids, ",")
	postsTable := this.Table("posts")
	sql := fmt.Sprintf("UPDATE %s SET deleted_at = NULL WHERE id in (%s)", postsTable, idString)

	if _, err := db.MasterDB.Exec(sql); err != nil {
		revel.AppLog.Errorf("Restore post failed: %s", err)
		return false, err
	}
	return true, nil
}

//real delete
func (this *PostTrashService) Destroy(ids []int) (bool, error) {
	if len(ids) < 1 {
		return false, errors.New("params error.")
	}
	if _, err := db.MasterDB.In("id", ids).Delete(&models.Posts{}); err != nil {
		revel.AppLog.Errorf("Destroy post failed: %s", err)
		return false, err
	}

	for _, id := range ids {
		cKey := utils.CacheKey("PostService", "InfoById", id)
		go cache.Delete(cKey)
	}

	return true, nil
}

func (this *PostTrashService) CountTrashPost() (int, error) {
	post := new(models.Posts)
	total, err := db.MasterDB.Where("deleted_at IS NULL OR `deleted_at` = '0001-01-01 00:00:00'").Count(post)
	if err != nil {
		revel.AppLog.Errorf("CountTrashPost Error: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostTrashService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
