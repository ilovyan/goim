package main

import (
	"github.com/ilovyan/goim/libs/define"
	"strconv"
	"strings"
)

// developer could implement "ThirdAuth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
	var err error
	var tokenUid string
	var tokenRid string
	var roomId64 int64

	if strings.ContainsAny(token, "|") {
		uidAndRid := strings.Split(token, "|")
		if len(uidAndRid) == 2 {
			tokenUid = uidAndRid[0]
			tokenRid = uidAndRid[1]
		}
	}
	if userId, err = strconv.ParseInt(tokenUid, 10, 64); err != nil {
		userId = 0
	}
	if roomId64, err = strconv.ParseInt(tokenRid, 10, 64); err != nil {
		roomId64 = define.NoRoom
	}
	return
}
