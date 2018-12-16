package service

import (
	"errors"
	"fmt"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"html/template"
	"time"
)

type PostService struct{}

func (this *PostService) GetPostById(id int, real bool) (models.Posts, error) {
	post := models.Posts{}

	if id <= 0 {
		return post, errors.New("param error")
	}

	cKey := utils.CacheKey("PostService", "InfoById", id)
	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &post); err != nil {
		_, err := db.MasterDB.Where("id = ?", id).Get(&post)
		if err != nil {
			return post, err
		}

		go cache.Set(cKey, post, 30*time.Minute)
	}

	//related tags
	tags, err := tagService.GetListByPostId(post.Id, real)
	if err != nil {
		revel.AppLog.Errorf("GetListByPostId failed : %s", err)
		return post, err
	}
	if len(tags) > 0 {
		post.Tags = tags
	} else {
		post.Tags = make([]models.Tags, 0)
	}

	post.Views = postService.GetViews(post, false)

	//related Categories
	category, catErr := categoryService.GetCategoryById(post.CategoryId, real)
	if catErr != nil {
		revel.AppLog.Errorf("GetCategoryById failed : %s", catErr)
		return post, catErr
	}
	post.Categories = category

	return post, nil
}

func (this *PostService) GetPostBySlug(slug string, real bool) (models.Posts, error) {
	post := models.Posts{}
	if slug == "" {
		return post, errors.New("Slug is empty.")
	}
	_, err := db.MasterDB.Where("slug = ?", slug).Cols("id").Get(&post)
	if err != nil {
		return post, err
	}

	post, pErr := this.GetPostById(post.Id, real)
	if pErr != nil {
		return post, err
	}

	return post, nil
}

func (this *PostService) GetSlugList(limit int, page int) ([]models.Posts, error) {
	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	start := (page - 1) * limit
	postArr := make([]models.Posts, 0)
	err := db.MasterDB.Desc("id").Cols("slug", "created_at").Limit(limit, start).Find(&postArr)
	if err != nil {
		return nil, err
	}

	return postArr, nil
}

func (this *PostService) GetPostIdArr(categoryId int, keywords string, limit int, page int) ([]int, error) {
	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	session := db.MasterDB.NewSession()
	start := (page - 1) * limit
	postList := make([]models.Posts, 0)
	if categoryId > 0 {
		session.And("category_id = ?", categoryId)
	}
	if len(keywords) > 0 {
		keywordsStr := "%" + keywords + "%"
		session.And("title like ?", keywordsStr)
	}
	err := session.Desc("id").Cols("id").Limit(limit, start).Find(&postList)
	if err != nil {
		return nil, err
	}

	idArr := make([]int, len(postList))
	for k, val := range postList {
		idArr[k] = val.Id
	}

	return idArr, nil
}

func (this *PostService) GetPostByIdArr(idArr []int, real bool) ([]models.Posts, error) {
	idLen := len(idArr)
	if idLen <= 0 {
		return nil, errors.New("Param error.")
	}

	postArr := make([]models.Posts, idLen)

	for k, postId := range idArr {
		post, err := this.GetPostById(postId, real)
		if err == nil {
			postArr[k] = post
		}
	}

	return postArr, nil
}

func (this *PostService) CountPost(categoryId int, keywords string) (int, error) {
	post := new(models.Posts)
	dbSession := db.MasterDB.NewSession()

	if categoryId > 0 {
		dbSession.And("category_id = ?", categoryId)
	}

	if keywords != "" {
		likeTitle := "%" + keywords + "%"
		dbSession.And("title like ?", likeTitle)
	}

	total, err := dbSession.Count(post)
	if err != nil {
		//revel.AppLog.Errorf("Count post failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostService) ShowPage(prefix string, page int, totalPage int) template.HTML {
	next := "<a rel='next' class='pager-btn pager-next' href='/%s/%d'>更早 &rarr;</a>"
	previous := "<a rel='prev' class='pager-btn pager-previous' href='/%s/%d'>&larr; 最近</a>"
	html := ""

	if totalPage == 1 {
		return template.HTML("")
	}

	if page < 1 {
		html = fmt.Sprintf(next, prefix, 1)
	} else if page == 1 {
		html = fmt.Sprintf(next, prefix, page+1)
	} else if page >= totalPage {
		newTotalPage := totalPage - 1
		previousNum := utils.IntDefault(newTotalPage > 0, newTotalPage, 1)
		html = fmt.Sprintf(previous, prefix, previousNum)
	} else {
		newTotalPage := totalPage - 1
		previousNum := utils.IntDefault(newTotalPage > 0, newTotalPage, 1)
		html = fmt.Sprintf(previous+next, prefix, previousNum, prefix, page+1)
	}

	return template.HTML(html)
}

func (this *PostService) GetListPaging(categoryId int, keywords string, limit int, page int, real bool) *models.PagingContent {
	dataList, _ := this.SearchList(categoryId, keywords, limit, page, real)
	total, _ := this.CountPost(0, "")
	totalPage := utils.GetTotalPage(total, limit)
	pagingContent := &models.PagingContent{Data: dataList,
		Total:       total,
		TotalPage:   totalPage,
		CurrentPage: page}
	return pagingContent
}

func (this *PostService) SearchList(categoryId int, keywords string, limit int, page int, real bool) ([]models.Posts, error) {
	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	session := db.MasterDB.NewSession()
	start := (page - 1) * limit
	postList := make([]models.Posts, 0)

	session.Where("deleted_at IS NULL")
	if categoryId > 0 {
		session.And("category_id = ?", categoryId)
	}
	if len(keywords) > 0 {
		keywordsStr := "%" + keywords + "%"
		session.And("title like ?", keywordsStr)
	}
	err := session.Desc("id").Cols("id").Limit(limit, start).Find(&postList)
	if err != nil {
		return nil, err
	}

	idArr := make([]int, len(postList))
	for k, v := range postList {
		idArr[k] = v.Id
	}

	postList, pErr := this.GetPostByIdArr(idArr, real)
	if pErr != nil {
		return nil, pErr
	}

	return postList, nil
}

func (this *PostService) GetPostTitleById(id int, real bool) (string, error) {
	post := models.Posts{}

	if id <= 0 {
		return "", errors.New("param error")
	}

	cKey := utils.CacheKey("PostService", "InfoById", id)
	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &post); err != nil {
		_, err := db.MasterDB.Where("id = ?", id).Get(&post)
		if err != nil {
			return "", err
		}

		go cache.Set(cKey, post, 30*time.Minute)
	}

	return post.Title, nil
}

//================================ use backend ===============================

func (this *PostService) Save(post models.Posts) (int, error) {
	if _, err := db.MasterDB.InsertOne(post); err != nil {
		revel.AppLog.Errorf("Save post failed : %s", err)
		return 0, err
	}
	return post.Id, nil
}

func (this *PostService) Update(id int, post models.Posts) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(post)
	if err != nil {
		revel.AppLog.Errorf("Update post failed: %s", err)
		return false, err
	}

	_, _ = this.GetPostById(id, true)
	return true, nil
}

func (this *PostService) Trash(ids []int) (bool, error) {
	post := models.Posts{DeletedAt: time.Now().Unix()}
	_, err := db.MasterDB.In("id", ids).Update(&post)
	if err != nil {
		revel.AppLog.Errorf("Destroy post failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *PostService) SumViews() (int, error) {
	post := new(models.Posts)
	total, err := db.MasterDB.Sum(post, "views")
	if err != nil {
		revel.AppLog.Errorf("Count post failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *PostService) IncrView(postId int, num int) {
	post := models.Posts{Views: num}
	_, err := db.MasterDB.Where("id = ?", postId).Update(&post)
	if err != nil {
		revel.AppLog.Errorf("Incr Views failed: %s", err)
	}
}

func (this *PostService) GetViews(post models.Posts, incr bool) int {
	cKey := fmt.Sprintf("persimmon:PostViews:%d", post.Id)
	val, err := redisObj.Do("EXISTS", cKey)

	num64, se := redis.Int64(val, err)
	if num64 <= 0 || se != nil {
		revel.AppLog.Errorf("EXISTS Error %s", err)
		_, e := redisObj.Do("SET", cKey, post.Views)
		if e != nil {
			revel.AppLog.Errorf("SET Error %s", e.Error())
			return post.Views
		}
	}

	if incr == false {
		val, err = redisObj.Do("GET", cKey)
		num64, se = redis.Int64(val, err)
		if se != nil {
			return 0
		}
		return int(num64)
	}

	val, err = redisObj.Do("INCR", cKey)
	num64, se = redis.Int64(val, err)
	if err != nil {
		revel.AppLog.Errorf("INCR Error %s", err.Error())
		return post.Views
	}

	var viewNum = int(num64)
	if num64%50 == 0 {
		go this.IncrView(post.Id, viewNum)
	}

	return viewNum
}

func (this *PostService) FormatPosts(postArr []models.Posts) []models.Posts {
	for k, v := range postArr {
		postArr[k].Views = postService.GetViews(v, false)
	}

	return postArr
}

func (this *PostService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
