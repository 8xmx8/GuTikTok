package main

import (
	"GuTikTok/src/constant/config"
	"GuTikTok/src/extra/profiling"
	"GuTikTok/src/extra/tracing"
	"GuTikTok/src/utils/logging"
	"GuTikTok/src/web/about"
	"GuTikTok/src/web/auth"
	comment2 "GuTikTok/src/web/comment"
	favorite2 "GuTikTok/src/web/favorite"
	feed2 "GuTikTok/src/web/feed"
	message2 "GuTikTok/src/web/message"
	"GuTikTok/src/web/middleware"
	publish2 "GuTikTok/src/web/publish"
	relation2 "GuTikTok/src/web/relation"
	user2 "GuTikTok/src/web/user"
	"context"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"time"
)

func main() {
	// Set Trace Provider
	tp, err := tracing.SetTraceProvider(config.WebServiceName)

	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicf("Error to set the trace")
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorf("Error to set the trace")
		}
	}()

	g := gin.Default()
	// Configure Prometheus
	p := ginprometheus.NewPrometheus("GuGoTik-WebGateway")
	p.Use(g)
	// Configure Gzip
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	// Configure Tracing
	g.Use(otelgin.Middleware(config.WebServiceName))
	g.Use(middleware.TokenAuthMiddleware())
	g.Use(middleware.RateLimiterMiddleWare(time.Second, 1000, 1000))

	// Configure Pyroscope
	profiling.InitPyroscope("GuGoTik.GateWay")

	// Register Service
	// Test Service
	g.GET("/about", about.Handle)
	// Production Service
	rootPath := g.Group("/douyin")
	user := rootPath.Group("/user")
	{
		user.GET("/", user2.UserHandler)
		user.POST("/login/", auth.LoginHandle)
		user.POST("/register/", auth.RegisterHandle)
	}
	feed := rootPath.Group("/feed")
	{
		feed.GET("/", feed2.ListVideosByRecommendHandle)
	}
	comment := rootPath.Group("/comment")
	{
		comment.POST("/action/", comment2.ActionCommentHandler)
		comment.GET("/list/", comment2.ListCommentHandler)
		comment.GET("/count/", comment2.CountCommentHandler)
	}
	relation := rootPath.Group("/relation")
	{
		//todo: frontend
		relation.POST("/action/", relation2.ActionRelationHandler)
		relation.POST("/follow/", relation2.FollowHandler)
		relation.POST("/unfollow/", relation2.UnfollowHandler)
		relation.GET("/follow/list/", relation2.GetFollowListHandler)
		relation.GET("/follower/list/", relation2.GetFollowerListHandler)
		relation.GET("/friend/list/", relation2.GetFriendListHandler)
		relation.GET("/follow/count/", relation2.CountFollowHandler)
		relation.GET("/follower/count/", relation2.CountFollowerHandler)
		relation.GET("/isFollow/", relation2.IsFollowHandler)
	}

	publish := rootPath.Group("/publish")
	{
		publish.POST("/action/", publish2.ActionPublishHandle)
		publish.GET("/list/", publish2.ListPublishHandle)
	}
	//todo
	message := rootPath.Group("/message")
	{
		message.GET("/chat/", message2.ListMessageHandler)
		message.POST("/action/", message2.ActionMessageHandler)
	}
	favorite := rootPath.Group("/favorite")
	{
		favorite.POST("/action/", favorite2.ActionFavoriteHandler)
		favorite.GET("/list/", favorite2.ListFavoriteHandler)
	}
	// Run Server
	if err := g.Run(config.WebServiceAddr); err != nil {
		panic("Can not run GuGoTik Gateway, binding port: " + config.WebServiceAddr)
	}
}
