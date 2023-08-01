package main

import (
	"SimpleTT/controller/basic"
	"SimpleTT/controller/interact"
	"SimpleTT/controller/social"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", basic.Feed)
	apiRouter.GET("/user/", basic.UserInfo)
	apiRouter.POST("/user/register/", basic.Register)
	apiRouter.POST("/user/login/", basic.Login)
	apiRouter.POST("/publish/action/", basic.Publish)
	apiRouter.GET("/publish/list/", basic.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", interact.FavoriteAction)
	apiRouter.GET("/favorite/list/", interact.FavoriteList)
	apiRouter.POST("/comment/action/", interact.CommentAction)
	apiRouter.GET("/comment/list/", interact.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", social.RelationAction)
	apiRouter.GET("/relation/follow/list/", social.FollowList)
	apiRouter.GET("/relation/follower/list/", social.FollowerList)
	apiRouter.GET("/relation/friend/list/", social.FriendList)
	apiRouter.GET("/message/chat/", social.MessageChat)
	apiRouter.POST("/message/action/", social.MessageAction)
}
