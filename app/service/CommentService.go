package service

import (
	"errors"
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"time"
)

type CommentService struct{}

func (this *CommentService) GetCommentById(id int, real bool) (models.Comments, error) {
	comments := models.Comments{}
	if id <= 0 {
		return comments, errors.New("param error")
	}

	cKey := utils.CacheKey("CommentService", "InfoById", id)

	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &comments); err != nil {
		_, err := db.MasterDB.Where("id = ?", id).Get(&comments)
		if err != nil {
			return comments, err
		}

		go cache.Set(cKey, comments, 30*time.Minute)
	}

	//post title
	post, _ := postService.GetPostById(comments.PostsId, false)
	comments.Title = post.Title
	comments.Slug = post.Slug

	return comments, nil
}

func (this *CommentService) GetCommentByIdArr(idArr []int, real bool) ([]models.Comments, error) {
	idLen := len(idArr)
	if idLen <= 0 {
		return nil, errors.New("参数不正确")
	}

	commentArr := make([]models.Comments, idLen)

	for k, commentId := range idArr {
		comment, err := this.GetCommentById(commentId, real)
		if err == nil {
			commentArr[k] = comment
		}
	}

	return commentArr, nil
}

func (this *CommentService) GetList(postId int, limit int, page int, real bool) ([]models.Comments, error) {
	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)
	postId = utils.IntDefault(postId > 0, postId, 0)
	start := (page - 1) * limit
	commentIdArr := make([]models.Comments, 0)
	dbSession := db.MasterDB.NewSession()

	if postId > 0 {
		dbSession.Where("posts_id = ?", postId)
	}

	err := dbSession.Cols("id").OrderBy("id DESC").Limit(limit, start).Find(&commentIdArr)
	if err != nil {
		revel.AppLog.Errorf("Get comment failed : %s", err)
		return nil, err
	}

	idArr := make([]int, len(commentIdArr))
	for k, v := range commentIdArr {
		idArr[k] = v.Id
	}

	if len(idArr) <= 0 {
		return make([]models.Comments, 0), nil
	}

	commentArr, cErr := this.GetCommentByIdArr(idArr, false)
	if cErr != nil {
		revel.AppLog.Errorf("GetCommentByIdArr Error : %s", cErr)
		return nil, cErr
	}

	return commentArr, nil
}

func (this *CommentService) GetListPaging(limit int, page int, real bool) (*models.PagingContent, error) {
	dataList, err := this.GetList(0, limit, page, real)
	if err != nil {
		return nil, err
	}

	total, cErr := this.CountComment(0)
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

func (this *CommentService) GetCommentByPostId(postId int, limit int, page int) (*models.PagingContent, error) {
	limit = utils.IntDefault(limit > 0, limit, 20)
	page = utils.IntDefault(page > 0, page, 1)

	comments := make([]models.Comments, 0)
	start := (page - 1) * limit
	err := db.MasterDB.Where("posts_id = ? AND status = 1", postId).OrderBy("id DESC").Limit(limit,
		start).Find(&comments)
	if err != nil {
		revel.AppLog.Errorf("Get comment by post id failed : %s", err)
		return nil, err
	}

	//对email进行Md5签名,内容@解析
	comments = utils.CommentRelolver(comments)

	total, cErr := this.CountComment(postId)
	if cErr != nil {
		return nil, cErr
	}

	totalPage := utils.GetTotalPage(total, limit)

	pagingContent := &models.PagingContent{Data: comments,
		Total:       total,
		TotalPage:   totalPage,
		CurrentPage: page}
	return pagingContent, nil
}

func (this *CommentService) Save(comment models.Comments) (int, error) {
	if _, err := db.MasterDB.InsertOne(&comment); err != nil {
		revel.AppLog.Errorf("Save comment failed : %s", err)
		return 0, err
	}
	return comment.Id, nil
}

func (this *CommentService) Update(id int, comment models.Comments) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(comment)
	if err != nil {
		revel.AppLog.Errorf("Update comment failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *CommentService) Destroy(id int, comment models.Comments) (bool, error) {
	_, err := db.MasterDB.Id(id).Delete(comment)
	if err != nil {
		revel.AppLog.Errorf("Destroy comment failed: %s", err)
		return false, err
	}
	return true, nil
}

func (this *CommentService) CountComment(postId int) (int, error) {
	comment := new(models.Comments)
	dbSession := db.MasterDB.NewSession()
	if postId > 0 {
		dbSession.Where("posts_id = ?", postId)
	}
	total, err := dbSession.Count(comment)
	if err != nil {
		revel.AppLog.Errorf("Count comment failed: %s", err)
		return 0, err
	}
	return int(total), nil
}

func (this *CommentService) Spam(id int, comment models.Comments) (bool, error) {
	_, err := db.MasterDB.Id(id).Update(comment)
	if err != nil {
		revel.AppLog.Errorf("update comment status: %s", err)
		return false, err
	}

	_, _ = this.GetCommentById(id, true)

	return true, nil
}

func (this *CommentService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
