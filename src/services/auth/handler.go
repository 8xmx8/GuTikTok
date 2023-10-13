package main

import (
	"GuTikTok/src/extra/tracing"
	"GuTikTok/src/models"
	"GuTikTok/src/rpc/auth"
	"GuTikTok/src/storage/cached"
	"GuTikTok/src/storage/database"
	"GuTikTok/strings"
	"GuTikTok/utils/checks"
	"GuTikTok/utils/logging"
	"context"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/hlandau/passlib.v1"
	"strconv"
)

var BloomFilter *bloom.BloomFilter

type AuthServiceImpl struct {
	auth.AuthServiceServer
}

func (a AuthServiceImpl) New() {

}
func (a AuthServiceImpl) Authenticate(ctx context.Context, request *auth.AuthenticateRequest) (resp *auth.AuthenticateResponse, err error) {
	ctx, span := tracing.Tracer.Start(ctx, "AuthenticateService")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("AuthService.Authenticate").WithContext(ctx)

	userId, ok, err := hasToken(ctx, request.Token)

	if err != nil {
		resp = &auth.AuthenticateResponse{
			StatusCode: strings.AuthServiceInnerErrorCode,
			StatusMsg:  strings.AuthServiceInnerError,
		}
		return
	}

	if !ok {
		resp = &auth.AuthenticateResponse{
			StatusCode: strings.UserNotExistedCode,
			StatusMsg:  strings.UserNotExisted,
		}
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":   err,
			"token": request.Token,
		}).Warnf("AuthService Authenticate Action failed to response when parsering uint")
		logging.SetSpanError(span, err)

		resp = &auth.AuthenticateResponse{
			StatusCode: strings.AuthServiceInnerErrorCode,
			StatusMsg:  strings.AuthServiceInnerError,
		}
		return
	}
	resp = &auth.AuthenticateResponse{
		StatusCode: strings.ServiceOKCode,
		StatusMsg:  strings.ServiceOK,
		UserId:     id,
	}
	return

}
func (a AuthServiceImpl) Register(ctx context.Context, request *auth.RegisterRequest) (resp *auth.RegisterResponse, err error) {

	ctx, span := tracing.Tracer.Start(ctx, "RegisterService")
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("AuthService.Register").WithContext(ctx)

	checkPwd := checks.ValidatePassword(request.Password, 8, 32)
	if !checkPwd {
		resp = &auth.RegisterResponse{
			StatusCode: strings.AuthInputPwdCode,
			StatusMsg:  strings.AuthInPutPwdExisted,
		}
	}
	resp = &auth.RegisterResponse{}
	var user models.User
	result := database.Client.WithContext(ctx).Limit(1).Where("name = ?", request.Username).Find(&user)
	if result.RowsAffected != 0 {
		resp = &auth.RegisterResponse{
			StatusCode: strings.AuthUserExistedCode,
			StatusMsg:  strings.AuthUserExisted,
		}
		return
	}
	var hashedPassword string
	if hashedPassword, err = hashPassword(ctx, request.Password); err != nil {
		logger.WithFields(logrus.Fields{
			"err":      result.Error,
			"username": request.Username,
		}).Warnf("AuthService Register Action failed to response when hashing password")
		logging.SetSpanError(span, err)

		resp = &auth.RegisterResponse{
			StatusCode: strings.AuthServiceInnerErrorCode,
			StatusMsg:  strings.AuthServiceInnerError,
		}
		fmt.Println(hashedPassword)
		//未完成
		return
	}

	result = database.Client.WithContext(ctx).Create(&user)
	if result.Error != nil {
		logger.WithFields(logrus.Fields{
			"err":      result.Error,
			"username": request.Username,
		}).Warnf("AuthService Register Action failed to response when creating user")
		logging.SetSpanError(span, result.Error)

		resp = &auth.RegisterResponse{
			StatusCode: strings.AuthServiceInnerErrorCode,
			StatusMsg:  strings.AuthServiceInnerError,
		}
		return
	}

	return

}
func (a AuthServiceImpl) Login(ctx context.Context, request *auth.LoginRequest) (resp *auth.LoginResponse, err error) {

	return nil, nil

}
func getToken(ctx context.Context, userId int64) (string, error) {
	span := trace.SpanFromContext(ctx)
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("AuthService.Login").WithContext(ctx)
	logger.WithFields(logrus.Fields{
		"userId": userId,
	}).Debugf("Select for user token")
	return cached.GetWithFunc(ctx, "U2T"+strconv.FormatInt(userId, 10),
		func(ctx context.Context, key string) (string, error) {
			span := trace.SpanFromContext(ctx)
			token := uuid.New().String()
			span.SetAttributes(attribute.String("token", token))
			cached.Write(ctx, "T2U"+token, strconv.FormatInt(userId, 10), true)
			return token, nil
		})
}

func hasToken(ctx context.Context, token string) (string, bool, error) {
	return cached.Get(ctx, "T2U"+token)
}
func hashPassword(ctx context.Context, password string) (string, error) {
	_, span := tracing.Tracer.Start(ctx, "PasswordHash")
	defer span.End()
	logging.SetSpanWithHostname(span)
	pwd, err := passlib.Hash(password)
	return pwd, err
}

func checkPasswordHash(ctx context.Context, password, hash string) bool {
	_, span := tracing.Tracer.Start(ctx, "PasswordHashChecked")
	defer span.End()
	logging.SetSpanWithHostname(span)
	newHash, err := passlib.Verify(password, hash)
	return err == nil && newHash == ""
}
