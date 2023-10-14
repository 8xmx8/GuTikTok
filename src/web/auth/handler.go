package auth

import (
	"GuTikTok/config"
	"GuTikTok/src/extra/tracing"
	"GuTikTok/src/rpc/auth"
	"GuTikTok/src/web/models"
	"GuTikTok/strings"
	"GuTikTok/utils/binder"
	grpc2 "GuTikTok/utils/grpc"
	"GuTikTok/utils/jsons"
	"GuTikTok/utils/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var Client auth.AuthServiceClient

func LoginHandler(c *gin.Context) {
	var req models.LoginReq
	_, span := tracing.Tracer.Start(c.Request.Context(), "LoginHandler")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("GateWay.Login").WithContext(c.Request.Context())

	if err := binder.Bind(c, &req); err != nil {
		c.JSON(http.StatusOK, models.LoginRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
			UserId:     0,
			Token:      "",
		})
		return
	}
	res, err := Client.Login(c.Request.Context(), &auth.LoginRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logger.WithFields(logrus.Fields{
			"Username": req.UserName,
		}).Warnf("Error when trying to connect with AuthService")
		c.Render(http.StatusOK, jsons.CustomJSON{Data: res, Context: c})
		return
	}
	logger.WithFields(logrus.Fields{
		"Username": req.UserName,
		"Token":    res.Token,
		"UserId":   res.UserId,
	}).Infof("User log in")

	c.Render(http.StatusOK, jsons.CustomJSON{Data: res, Context: c})
}

func RegisterHandler(c *gin.Context) {
	var req models.RegisterReq
	_, span := tracing.Tracer.Start(c.Request.Context(), "LoginHandler")
	defer span.End()
	logger := logging.LogService("GateWay.Register").WithContext(c.Request.Context())

	if err := binder.Bind(c, &req); err != nil {
		c.JSON(http.StatusOK, models.RegisterRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
			UserId:     0,
			Token:      "",
		})
		return
	}

	res, err := Client.Register(c.Request.Context(), &auth.RegisterRequest{
		Username: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"Username": req.UserName,
		}).Warnf("Error when trying to connect with AuthService")
		c.Render(http.StatusOK, jsons.CustomJSON{Data: res, Context: c})
		return
	}

	logger.WithFields(logrus.Fields{
		"Username": req.UserName,
		"Token":    res.Token,
		"UserId":   res.UserId,
	}).Infof("User register in")

	c.Render(http.StatusOK, jsons.CustomJSON{Data: res, Context: c})
}
func init() {
	coon := grpc2.Connect(config.AuthRpcServerName)
	Client = auth.NewAuthServiceClient(coon)
}
