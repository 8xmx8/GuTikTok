package main

import (
	"GuTikTok/config"
	"GuTikTok/src/rpc/relation"
	"GuTikTok/src/rpc/user"
	"fmt"
)

var userClient user.UserServiceClient

var actionRelationLimitKeyPrefix = config.Conf.Redis.RedisPrefix + "relation_freq_limit"

const actionRelationMaxQPS = 3

type RelationServiceImpl struct {
	relation.RelationServiceServer
}

func actionRelationLimitKey(userId uint32) string {
	return fmt.Sprintf("%s-%d", actionRelationLimitKeyPrefix, userId)
}

func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}
