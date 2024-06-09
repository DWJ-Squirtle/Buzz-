package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"bytes"
	"io"
	"net/http"
	"encoding/json"
)

// 评论

// CommentHandler 创建评论
func CommentHandler(c *gin.Context) {
	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		fmt.Println("ziduanpiprishibai",err)
		zap.L().Error("字段匹配失败",zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 生成评论ID
	commentID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	comment.CommentID = commentID
	comment.AuthorID = userID
	comment.Analyse = analy(comment.Content)
	// 创建评论
	if err := mysql.CreateComment(&comment); err != nil {
		zap.L().Error("mysql.CreateComment(&comment) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {
	id, ok := c.GetQuery("id")
	fmt.Println("id=========================",id)
	if !ok {
		ResponseError(c, CodeInvalidParams)
		return
	}
	posts, err := mysql.GetCommentListByIDs(id)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println("======================================",posts)
	ResponseSuccess(c, posts)
}
func DeleteComment(c *gin.Context) {
	id := c.Query("comment_id")
	fmt.Println("++++++++++++++++controllerid",id)
	if id == "" {
		ResponseError(c, CodeInvalidParams)
		return
	}
	if err := mysql.DeleteCommentByID(id); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
func analy(content string) uint8 {
    url := "http://127.0.0.1:5001/analyse/sentiment"

    // 构造请求体
    requestBody := map[string]string{"content": content}
    requestBodyBytes, err := json.Marshal(requestBody)
    if err != nil {
        fmt.Println("构造请求体时发生错误:", err)
        return 0
    }

    // 发送 POST 请求
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
    if err != nil {
        fmt.Println("发送时发生错误:", err)
        return 0
    }
    defer resp.Body.Close()

    // 读取响应体
    responseBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("读取响应时发生错误:", err)
        return 0
    }

    // 输出响应内容
    fmt.Println("响应内容:", string(responseBytes))
	var response struct {
    Code int `json:"code"`
    Data struct {
        Sentiments float64 `json:"sentiments"`
    } `json:"data"`
	}

	// 解码响应
	if err := json.Unmarshal(responseBytes, &response); err != nil {
    fmt.Println("解码响应时发生错误:", err)
    return 0
	}

	// 获取情感分析结果
	sentiments := response.Data.Sentiments

	// 根据情感分析结果返回相应的值
	if sentiments < 0.4 {
    	return 1
	} else{
		return 2
	}
}
func Commentgood(c *gin.Context){
	id, ok := c.GetQuery("id")
	fmt.Println("goodid=========================",id)

	if !ok {
		ResponseError(c, CodeInvalidParams)
		return
	}
	posts, err := mysql.GetGoodbyid(id)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println("post=========================",posts)

	ResponseSuccess(c, posts)
}
func Commentbad(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		ResponseError(c, CodeInvalidParams)
		return
	}
	posts, err := mysql.GetBadbyid(id)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, posts)


}
