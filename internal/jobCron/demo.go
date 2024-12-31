package jobCron

import "github.com/zeromicro/go-zero/core/logc"

func (j *JobCron) runDemo() {
	logc.Info(j.ctx, "demo running")
}
