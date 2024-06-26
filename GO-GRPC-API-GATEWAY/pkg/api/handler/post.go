package handler

import (
	"net/http"
	"strconv"

	interfaces "github.com/akhi9550/api-gateway/pkg/client/interface"
	"github.com/akhi9550/api-gateway/pkg/helper"
	"github.com/akhi9550/api-gateway/pkg/logging"
	"github.com/akhi9550/api-gateway/pkg/utils/models"
	"github.com/akhi9550/api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	GRPC_Client interfaces.PostClient
	PostCachig  *helper.RedisPostCaching
}

func NewPostHandler(postClient interfaces.PostClient, postCaching *helper.RedisPostCaching) *PostHandler {
	return &PostHandler{
		GRPC_Client: postClient,
		PostCachig:  postCaching,
	}
}

// @Summary			Create a Post
// @Description		User Create a Post
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param caption formData string true "New caption for the post"
// @Param posttype formData string true "New type of the post"
// @Param user formData array true "Users associated with the post. Provide multiple user IDs"
// @Param           photo formData file true "Photo of the post"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/post   [POST]
func (p *PostHandler) CreatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	caption := c.PostForm("caption")
	typeid := c.PostForm("posttype")
	user := c.PostFormArray("user")
	users, err := helper.ConvertStringToArray(user)
	if err != nil {
		return
	}
	for _, i := range user {
		if i == "1" {
			return
		}
	}
	req := models.PostRequest{
		Caption: caption,
		TypeId:  typeid,
	}
	file, err := c.FormFile("photo")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "No file provided", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	data, err := p.GRPC_Client.CreatePost(userID.(int), req, file, users)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully Posted", data, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary		Show PostType
// @Description	 Show posttypes form user side
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/post/type     [GET]
func (p *PostHandler) ShowPostTypeUser(c *gin.Context) {
	data, err := p.GRPC_Client.ShowPostType()
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't Remove Post", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	sucess := response.ClientResponse(http.StatusOK, "Successfully Get Posttype", data, nil)
	c.JSON(http.StatusOK, sucess)
}

// @Summary			Show User Posts
// @Description		Retrieve  User Posts
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/post   [GET]
func (p *PostHandler) GetUserPost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	data, err := p.GRPC_Client.GetUserPost(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Get a Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Show Specific Post
// @Description		Retrieve Specific Post With Its ID
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			post_id	 query	string	true	"post id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/post/posts   [GET]
func (p *PostHandler) GetPost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "PostID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.PostCachig.GetPost(userID.(int), int(PostID))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Get a Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Update a Post
// @Description		Update  Specific Post
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param post_id formData string true "ID of the post to be updated"
// @Param caption formData string true "New caption for the post"
// @Param posttype formData string true "New type of the post"
// @Param user formData array true "Users associated with the post. Provide multiple user IDs"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/post   [put]
func (p *PostHandler) UpdatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postid := c.PostForm("post_id")
	caption := c.PostForm("caption")
	typeid := c.PostForm("posttype")
	user := c.PostFormArray("user")
	users, err := helper.ConvertStringToArray(user)
	if err != nil {
		return
	}
	for _, i := range user {
		if i == "1" {
			return
		}
	}

	postID, _ := strconv.Atoi(postid)
	req := models.UpdatePostReq{
		PostID:  uint(postID),
		Caption: caption,
		TypeID:  typeid,
	}

	data, err := p.GRPC_Client.UpdatePost(userID.(int), req, users)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully Updated", data, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary			Delete a Post
// @Description		Delete a Specific Post
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			post_id	 query	string	true	"Post id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/post   [DELETE]
func (p *PostHandler) DeletePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.DeletePost(userID.(int), PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Deleted Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Show  Specific User Posts
// @Description		Show  Posts With Specific UserID
// @Tags			Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			user_id	 query	string	true	"User id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/post/getposts   [GET]
func (p *PostHandler) GetAllPost(c *gin.Context) {
	userID := c.Query("user_id")
	UserID, err := strconv.Atoi(userID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "UserID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.GRPC_Client.GetAllPost(UserID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Get a Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Archive Specific User Post
// @Description		Archive Specific User Post
// @Tags			Archive Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			post_id	 query	string	true	"post id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/archive   [POST]
func (p *PostHandler) ArchivePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.ArchivePost(userID.(int), PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Archive Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			UnArchive Specific User Post
// @Description		UnArchive Specific User Post
// @Tags			Archive Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			post_id	 query	string	true	"post id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/archive/unarchive   [POST]
func (p *PostHandler) UnArchivePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.UnArchivePost(userID.(int), PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully UnArchive Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Get All Archive User Posts
// @Description		Get All Archive User Posts
// @Tags			Archive Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/archive   [GET]
func (p *PostHandler) GetAllArchivePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	data, err := p.PostCachig.GetAllArchivePost(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Get a Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Like a Post
// @Description		Like a Post
// @Tags			Like Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			post_id	 query	string	true	"post id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/like   [PUT]
func (p *PostHandler) LikePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.GRPC_Client.LikePost(userID.(int), PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Successfully Liked Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			UnLike a Post
// @Description		UnLike a Post
// @Tags			Like Post
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			post_id	 query	string	true	"post id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/like/unlike   [PUT]
func (p *PostHandler) UnLikePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.UnLikePost(userID.(int), PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Successfully Liked Post", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully UnLiked Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Comment Specific Post
// @Description	Comment Specific Post
// @Tags			Comment
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			req	body		models.PostCommentReq	true	"Commnet details"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/comment   [POST]
func (p *PostHandler) PostComment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req models.PostCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	data, err := p.GRPC_Client.PostComment(userID.(int), req)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Commented Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Show All Comments on Post
// @Description	Show All Comments on Post
// @Tags			Comment
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			post_id	 query	string	true	"Post id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/comment   [GET]
func (p *PostHandler) GetAllPostComments(c *gin.Context) {
	postID := c.Query("post_id")
	PostID, err := strconv.Atoi(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.GRPC_Client.GetAllPostComments(PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Received All Comments in the Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Delete specific user Comment on Post
// @Description		Delete specific user Comment on Post
// @Tags			Comment
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			comment_id	 query	string	true	"Comment id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/comment  [DELETE]
func (p *PostHandler) DeleteComment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	comment := c.Query("comment_id")
	commentID, err := strconv.Atoi(comment)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "CommentID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.DeleteComment(userID.(int), commentID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Deleted Comment in the Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Reply specific user Comment on Post
// @Description		Reply specific user Comment on Post
// @Tags			Comment
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			req	 body		models.ReplyCommentReq	true	"Admin login details"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/comment/reply  [POST]
func (p *PostHandler) ReplyComment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req models.ReplyCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	data, err := p.GRPC_Client.ReplyComment(userID.(int), req)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Received Comment in the Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Show All Comments and Replies on Specific Post
// @Description	Show All Comments and Replies on Specific Post
// @Tags			Comment
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			post_id	 query	string	true	"Post id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/comment/showcomments   [GET]
func (p *PostHandler) ShowAllPostComments(c *gin.Context) {
	PostID := c.Query("post_id")
	postID, err := strconv.Atoi(PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "PostID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.GRPC_Client.ShowAllPostComments(postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User successfully Recived All Comments", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			User Can Saved Specific Post
// @Description		User Can Saved Specific Post
// @Tags			Save Post
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			post_id	 query	string	true	"Post id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/saved    [POST]
func (p *PostHandler) SavedPost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	PostID := c.Query("post_id")
	postID, err := strconv.Atoi(PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.SavedPost(userID.(int), postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Saved Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			User Can UnSaved Specific Post
// @Description		User Can UnSaved Specific Post
// @Tags			Save Post
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			post_id	 query	string	true	"Post id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/saved/unsaved  [POST]
func (p *PostHandler) UnSavedPost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	PostID := c.Query("post_id")
	postID, err := strconv.Atoi(PostID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.UnSavedPost(userID.(int), postID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully UnSaved Post", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Get User Saved Posts
// @Description		Get User Saved Posts
// @Tags			Save Post
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/saved    [GET]
func (p *PostHandler) GetSavedPost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	data, err := p.PostCachig.GetSavedPost(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Received Saved Post", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Create  Story
// @Description		Create User Story
// @Tags			Story
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param           photo formData	 file true "Image file to upload" collectionFormat "multi"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/story     [POST]
func (p *PostHandler) CreateStory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	file, err := c.FormFile("photo")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "No file provided", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	data, err := p.GRPC_Client.CreateStory(userID.(int), file)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully Posted Stroy", data, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary			Get Story
// @Description		Get Specific User Story
// @Tags			Story
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			user_id	 query	string	true	"User id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/story     [GET]
func (p *PostHandler) GetStory(c *gin.Context) {
	viewUser, _ := c.Get("user_id")
	UserID := c.Query("user_id")
	userID, err := strconv.Atoi(UserID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Post_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.GRPC_Client.GetStory(userID, viewUser.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User successfully Recived All Stroy", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Delete Story
// @Description		Delete User Story
// @Tags			Story
// @Accept 			json
// @Produce 		json
// @Security 		Bearer
// @Param			story_id	 query	string	true	"Story id"
// @Success			200		{object}	response.Response{}
// @Failure			500		{object}	response.Response{}
// @Router			/story     [DELETE]
func (p *PostHandler) DeleteStory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	StoryID := c.Query("story_id")
	storyID, err := strconv.Atoi(StoryID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "StoryID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.DeleteStory(userID.(int), storyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User successfully Deleted Stroy", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Like Story
// @Description		Liked User Story
// @Tags			Story
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			story_id	 query	string	true	"Story id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/story/like     [PUT]
func (p *PostHandler) LikeStory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	StoryID := c.Query("story_id")
	storyID, err := strconv.Atoi(StoryID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "StoryID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.LikeStory(userID.(int), storyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User successfully Liked Stroy", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			UnLike Story
// @Description		UnLiked User Story
// @Tags			Story
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			story_id	 query	string	true	"Story id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/story/unlike   [PUT]
func (p *PostHandler) UnLikeStory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	StoryID := c.Query("story_id")
	storyID, err := strconv.Atoi(StoryID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "StoryID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.UnLikeStory(userID.(int), storyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully UnLiked Stroy", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary			Get User Story Details
// @Description		Get User Story Details
// @Tags			Story
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			story_id	 query	string	true	"Story id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/story/details   [GET]
func (p *PostHandler) StoryDetails(c *gin.Context) {
	userID, _ := c.Get("user_id")
	StoryID := c.Query("story_id")
	storyID, err := strconv.Atoi(StoryID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "StoryID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	data, err := p.PostCachig.StoryDetails(userID.(int), storyID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Retrive Stroy Details", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary 	Report Post
// @Description Report User to Post
// @Tags 		Reports
// @Accept 		json
// @Produce 	json
// @Security 	Bearer
// @Param 		body 	body models.ReportPostRequest	 true 	"Post Report"
// @Success 	200 {object} response.Response{}
// @Failure 	500 {object} response.Response{}
// @Router 		/report/post     [POST]
func (p *PostHandler) ReportPost(c *gin.Context) {
	ReportedID, _ := c.Get("user_id")
	var req models.ReportPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	err := p.GRPC_Client.ReportPost(ReportedID.(int), req)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Reported", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary 	HomePage
// @Description HomePage
// @Tags 		Home
// @Accept 		json
// @Produce 	json
// @Security 	Bearer
// @Success 	200 {object} response.Response{}
// @Failure 	500 {object} response.Response{}
// @Router 		/post/home     [GET]
func (p *PostHandler) Home(c *gin.Context) {
	userID, _ := c.Get("user_id")
	data, err := p.PostCachig.Home(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Couldn't Get HomePage", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Get HomePage", data, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Create PostType
// @Description	 Admin can create a posttype
// @Tags			Admin Posttype Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			Posttype	body		models.PostType	true	"details"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/post/type     [POST]
func (p *PostHandler) CreatePostType(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "CreatePostType")
	logEntry.Info("Processing CreatePostType")
	var Posttype models.PostType
	if err := c.ShouldBindJSON(&Posttype); err != nil {
		logEntry.WithError(err).Error("Error binding request body")
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}
	err := p.GRPC_Client.CreatePostType(Posttype)
	if err != nil {
		logEntry.WithError(err).Error("Error During CreatePostType RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't Create PostType", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logEntry.Info("Successfully Created Posttype")
	sucess := response.ClientResponse(http.StatusOK, "Successfully Created Posttype", nil, nil)
	c.JSON(http.StatusOK, sucess)
}

// @Summary		Show PostType
// @Description	 Show posttypes form admin side
// @Tags			Admin Posttype Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/post/type     [GET]
func (p *PostHandler) ShowPostType(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "CreatePostType")
	logEntry.Info("Processing CreatePostType")
	data, err := p.GRPC_Client.ShowPostType()
	if err != nil {
		logEntry.WithError(err).Error("Error During GetPostType RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't Get PostType", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logEntry.Info("Successfully Get Posttype")
	sucess := response.ClientResponse(http.StatusOK, "Successfully Get Posttype", data, nil)
	c.JSON(http.StatusOK, sucess)
}

// @Summary		Delete PostType
// @Description	 Admin can create a posttype
// @Tags			Admin Posttype Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param			posttype_id	 query	string	true	"Posttype id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/post/type     [DELETE]
func (p *PostHandler) DeletePostType(c *gin.Context) {
	logEntry := logging.GetLogger().WithField("context", "CreatePostType")
	logEntry.Info("Processing CreatePostType")
	PostTypeID := c.Query("posttype_id")
	postTypeID, err := strconv.Atoi(PostTypeID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "PostTypeID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = p.GRPC_Client.DeletePostType(postTypeID)
	if err != nil {
		logEntry.WithError(err).Error("Error During DeletePostType RPC call")
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't Delete PostType", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	logEntry.Info("Successfully Deleted Posttype")
	sucess := response.ClientResponse(http.StatusOK, "Successfully Deleted Posttype", nil, nil)
	c.JSON(http.StatusOK, sucess)
}
