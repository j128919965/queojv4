package svc

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"queoj/service/user/internal/model"
	"sync"
	"time"
)

type CoinsAndPointsSyncer struct {
	Ch chan model.UserAccount
	mu sync.RWMutex
	m  map[uint64][2]int32
	db *gorm.DB
}

func (c *CoinsAndPointsSyncer) trySync() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.m) == 0 {
		return
	}
	errMap := make(map[uint64][2]int32)
	for userId, ints := range c.m {
		err := c.db.Model(&model.UserAccount{ID:userId}).Updates(map[string]interface{}{
			"coins":gorm.Expr("coins + ?",ints[0]),
			"point":gorm.Expr("point + ?",ints[1]),
		}).Error
		if err != nil {
			logx.Error(fmt.Errorf("sync coins %d and %d points for user %d failed :%s" ,userId,ints[0],ints[1],err))
			errMap[userId] = ints
			continue
		}
	}
	c.m = errMap
}

func NewCoinsAndPointsSyncer(ctx context.Context, db *gorm.DB) *CoinsAndPointsSyncer {
	syncer := &CoinsAndPointsSyncer{
		Ch: make(chan model.UserAccount, 20),
		m:  make(map[uint64][2]int32),
		db: db,
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case account := <- syncer.Ch:
				ints := syncer.m[account.ID]
				ints[0] += account.Coins
				ints[1] += account.Point
				syncer.m[account.ID] = ints
			default:
				syncer.trySync()
				time.Sleep(time.Second * 5)
			}
		}
	}()

	return syncer
}
