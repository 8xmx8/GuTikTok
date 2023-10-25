package main

import (
	"GuTikTok/src/extra/tracing"
	"GuTikTok/src/models"
	"GuTikTok/src/rpc/user"
	"GuTikTok/src/storage/cached"
	"GuTikTok/strings"
	"GuTikTok/utils/logging"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

type UserServiceImpl struct {
	user.UserServiceServer
}

func (a UserServiceImpl) New() {

}

func (a UserServiceImpl) GetUserInfo(ctx context.Context, request *user.UserRequest) (resp *user.UserResponse, err error) {
	ctx, span := tracing.Tracer.Start(ctx, "GetUserInfo")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("UserService.GetUserInfo").WithContext(ctx)

	var userModel models.User
	userModel.ID = request.UserId
	ok, err := cached.ScanGet(ctx, "UserInfo", &userModel)

	if err != nil {

		resp = &user.UserResponse{
			StatusCode: strings.UserServiceInnerErrorCode,
			StatusMsg:  strings.UserServiceInnerError,
		}
		return
	}

	if !ok {
		resp = &user.UserResponse{
			StatusCode: strings.UserNotExistedCode,
			StatusMsg:  strings.UserNotExisted,
			User:       nil,
		}
		logger.WithFields(logrus.Fields{
			"user": request.UserId,
		}).Infof("Do not exist")
		return
	}

	resp = &user.UserResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		User: &user.User{
			Id:              request.UserId,
			Name:            userModel.Name,
			FollowCount:     nil,
			FollowerCount:   nil,
			IsFollow:        false,
			Avatar:          &userModel.Avatar,
			BackgroundImage: &userModel.BackgroundImage,
			Signature:       &userModel.Signature,
			TotalFavorited:  nil,
			WorkCount:       nil,
			FavoriteCount:   nil,
		},
	}

	var wg sync.WaitGroup

	wg.Add(6)

	return
}

func (a UserServiceImpl) GetUserExistInformation(ctx context.Context, request *user.UserExistRequest) (resp *user.UserExistResponse, err error) {
	ctx, span := tracing.Tracer.Start(ctx, "GetUserExisted")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("UserService.GetUserExisted").WithContext(ctx)

	var userModel models.User
	userModel.ID = request.UserId
	ok, err := cached.ScanGet(ctx, "UserInfo", &userModel)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("Error when selecting user info")
		logging.SetSpanError(span, err)
		resp = &user.UserExistResponse{
			StatusCode: strings.UserServiceInnerErrorCode,
			StatusMsg:  strings.UserServiceInnerError,
			Existed:    false,
		}
		return
	}

	if !ok {
		resp = &user.UserExistResponse{
			StatusCode: strings.ServiceOKCode,
			StatusMsg:  strings.ServiceOK,
			Existed:    false,
		}
		logger.WithFields(logrus.Fields{
			"user": request.UserId,
		}).Infof("User do not exist")
		return
	}

	resp = &user.UserExistResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		Existed:    true,
	}
	return
}
