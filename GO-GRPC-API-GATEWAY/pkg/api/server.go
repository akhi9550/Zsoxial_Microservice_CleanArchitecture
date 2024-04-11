package server

import (
	"log"

	"github.com/akhi9550/api-gateway/pkg/api/handler"
	"github.com/akhi9550/api-gateway/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler *handler.AuthHandler, postHandler *handler.PostHandler) *ServerHTTP {
	r := gin.New()

	r.Use(gin.Logger())

	r.POST("/login", authHandler.AdminLogin)

	// r.Use(middleware.AdminAuthMiddleware())
	r.GET("admin/users", middleware.AdminAuthMiddleware(), authHandler.ShowAllUsers)
	r.PUT("admin/user/block", middleware.AdminAuthMiddleware(), authHandler.BlockUser)
	r.PUT("admin/user/unblock", middleware.AdminAuthMiddleware(), authHandler.UnBlockUser)
	r.GET("/admin/report/user", middleware.AdminAuthMiddleware(), authHandler.ShowUserReports)
	r.GET("/admin/report/post", middleware.AdminAuthMiddleware(), authHandler.ShowPostReports)
	r.GET("/admin/posts", middleware.AdminAuthMiddleware(), authHandler.GetAllPosts)
	r.POST("/admin/post/remove", middleware.AdminAuthMiddleware(), authHandler.RemovePost)

	r.POST("user/signup", authHandler.UserSignup)
	r.POST("user/login", authHandler.Userlogin)

	r.POST("user/send-otp", authHandler.SendOtp)
	r.POST("user/verify-otp", authHandler.VerifyOtp)

	r.POST("user/forgot-password", authHandler.ForgotPassword)
	r.POST("user/forgot-password-verify", authHandler.ForgotPasswordVerifyAndChange)

	r.Use(middleware.UserAuthMiddleware())
	{
		user := r.Group("/user")
		{
			user.GET("", authHandler.UserDetails)
			user.PUT("", authHandler.UpdateUserDetails)
			user.PUT("/changepassword", authHandler.ChangePassword)
			user.GET("/search", authHandler.SearchUser)
		}

		report := r.Group("/report")
		{
			report.POST("/user", authHandler.ReportUser)
			report.POST("/post", postHandler.ReportPost)
		}

		follow := r.Group("/follow")
		{
			follow.POST("/request", authHandler.FollowREQ)
			follow.GET("/requests", authHandler.ShowFollowREQ)
			follow.POST("/accept", authHandler.AcceptFollowREQ)
			follow.POST("/unfollow", authHandler.UnFollow)
			follow.GET("/following", authHandler.Following)
			follow.GET("/followers", authHandler.Follower)
		}

		post := r.Group("/post")
		{
			post.POST("", postHandler.CreatePost)
			post.GET("", postHandler.GetPost)
			post.PUT("", postHandler.UpdatePost)
			post.DELETE("", postHandler.DeletePost)
			post.GET("/getposts", postHandler.GetAllPost)
		}

		savepost := r.Group("/saved")
		{
			savepost.POST("", postHandler.SavedPost)
			savepost.GET("", postHandler.GetSavedPost)
			savepost.POST("/unsaved", postHandler.UnSavedPost)

		}

		archive := r.Group("/archive")
		{
			archive.POST("", postHandler.ArchivePost)
			archive.GET("", postHandler.GetAllArchivePost)
			archive.POST("/unarchive", postHandler.UnArchivePost)

		}

		like := r.Group("/like")
		{
			like.PUT("", postHandler.LikePost)
			like.PUT("/unlike", postHandler.UnLinkPost)
		}

		comment := r.Group("/comment")
		{
			comment.POST("", postHandler.PostComment)
			comment.DELETE("", postHandler.DeleteComment)
			comment.POST("/reply", postHandler.ReplyComment)
			comment.GET("", postHandler.GetAllPostComments)
			comment.GET("/showcomments", postHandler.ShowAllPostComments)
		}

		story := r.Group("/story")
		{
			story.POST("", postHandler.CreateStory)
			story.GET("", postHandler.GetStory)
			story.DELETE("", postHandler.DeleteStory)
			story.PUT("/like", postHandler.LikeStory)
			story.PUT("/unlike", postHandler.UnLikeStory)
			story.GET("/details", postHandler.StoryDetails)
		}

	}
	return &ServerHTTP{engine: r}
}
func (s *ServerHTTP) Start() {
	log.Printf("Starting Server on 8000")
	err := s.engine.Run(":8000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
