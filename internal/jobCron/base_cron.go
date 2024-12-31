package jobCron

import (
	"context"
	"fmt"
	"github.com/robfig/cron"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"go-api/internal/jobCron/job"
	"go-api/internal/svc"
)

type JobCron struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*job.ActiveCount
	OlinePlayCountFirst  bool
	UserActiveCountFirst bool
}

func NewJobCron(ctx context.Context, svcCtx *svc.ServiceContext) *JobCron {
	logic := &JobCron{
		ctx:                  ctx,
		svcCtx:               svcCtx,
		ActiveCount:          job.NewActiveCount(ctx, svcCtx),
		OlinePlayCountFirst:  true,
		UserActiveCountFirst: true,
	}
	return logic
}

func (j *JobCron) registerFunc(runFunc func(), spec string) {
	runFunc()
	c := cron.New()
	err := c.AddFunc(spec, func() {
		runFunc()
	})
	if err != nil {
		logx.Errorf("spec err: %v", err)
	}
	c.Start()
}

// Run 需要增加定时函数在这里注册即可
func (j *JobCron) Run() {
	defer func() {
		if r := recover(); r != nil {
			logc.Error(j.ctx, "JobCron:", fmt.Sprintf("Recovered in f, r=%v\n", r))
		}
	}()
	j.registerFunc(j.runUserActiveCountFirst, "0 10 0 * *") //每天的0点10分执行一次
}

func (j *JobCron) runUserActiveCountFirst() {
	if j.UserActiveCountFirst {
		j.UserActiveCountFirst = false
		return
	}
	j.ActiveCount.Run()
}
