package mysql

import (
	"bluebell_backend/models"
    "database/sql"
	"fmt"
	"go.uber.org/zap"
	"strconv"
)

func CreateComment(comment *models.Comment) (err error) {
	sqlStr := `insert into comment(
	comment_id, content, post_id, author_id, parent_id, headImg, name,analyse)
	values(?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID,
		comment.AuthorID, comment.ParentID, comment.HeadImg, comment.Name, comment.Analyse)
	fmt.Println("插入时的ID",comment.CommentID)
	if err != nil {
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}
func GetCommentListByIDs(id string) (commentList []*models.Comment, err error) {
	sqlr:= `select name, headImg,  content, create_time, comment_id  from comment where post_id = ?`
	err = db.Select(&commentList, sqlr,id)
	fmt.Printf("Number of comments retrieved: %d\n", len(commentList))

    // 打印每个评论的内容
    for i, comment := range commentList {
        fmt.Printf("Comment %d: %v\n", i+1, comment)
    }
	if err == sql.ErrNoRows { // 查询为空
		zap.L().Warn("there is no comment in db")
		err = nil
	}
	fmt.Printf("Executing SQL: %s, with ID: %s\n", sqlr, id)
	fmt.Println("huoqushide id",id)
	fmt.Println("获取时的neirong++++++",commentList)
	for _, comment := range commentList {
		comment.CommentIDstr = strconv.FormatUint(comment.CommentID,10)
    }
	return commentList,nil
}
func GetGoodbyid(id string) (commentList []*models.Comment, err error) {
	sqlStr := `select name, headImg,  content, create_time, comment_id from comment where post_id = ? and analyse = 2`
	err = db.Select(&commentList, sqlStr, id)
	for _, comment := range commentList {
		comment.CommentIDstr = strconv.FormatUint(comment.CommentID,10)
   }

	fmt.Println("获取时的ID",&commentList)
	return commentList, nil
}
func GetBadbyid(id string) (commentList []*models.Comment, err error) {
	sqlStr := `select content, name, headImg, comment_id, create_time from comment where post_id = ? and analyse = 1`
	err = db.Select(&commentList, sqlStr, id)
	for _, comment := range commentList {
		comment.CommentIDstr = strconv.FormatUint(comment.CommentID,10)
	}
	fmt.Println("获取时的ID",&commentList)
	return commentList, nil
}


func DeleteCommentByID(id string)(err error) {
	fmt.Println("=========================",id)
	sqlStr := `delete from comment where comment_id = ? `
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		zap.L().Error("delete comment failed", zap.Error(err))
	}
	return
}
