package nodemgr

import (
	"context"
	"time"

	"github.com/fananchong/go-xserver/common"
	"github.com/fananchong/go-xserver/internal/db"
	"github.com/fananchong/go-xserver/internal/utility"
)

// RegisterSelfHelper : 管理节点注册自己到 Redis 数据库的帮助类
type RegisterSelfHelper struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// Start : 开始
func (helper *RegisterSelfHelper) Start() {
	helper.ctx, helper.ctxCancel = context.WithCancel(context.Background())
	go helper.loop()
}

func (helper *RegisterSelfHelper) loop() {
	timer := time.NewTicker(time.Second)
	defer timer.Stop()
	for {
		select {
		case <-helper.ctx.Done():
			common.XLOG.Infoln("RegisterSelfHelper close.")
			return
		case <-timer.C:
			helper.register()
		}
	}
}

func (helper *RegisterSelfHelper) register() {
	data := db.NewMgrServer(common.XCONFIG.DbMgr.Name, 0)
	data.SetAddr(utility.GetIPInner())
	data.SetPort(uint16(utility.GetIntranetListenPort()))
	if err := data.Save(); err != nil {
		common.XLOG.Errorln(err)
	}
}

// Close : 结束
func (helper *RegisterSelfHelper) Close() {
	if helper.ctxCancel != nil {
		helper.ctxCancel()
	}
}