package tools

import (
	"context"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vlog"
	"time"
)

// RetryExecute 尝试多次执行一个函数
func RetryExecute(ctx context.Context, tryMaxTimes int, f func(current int) error) error {
	var (
		count = 0
		err   error
	)
	for {
		if count >= tryMaxTimes {
			// 超过次数
			break
		}
		count++
		if err = f(count); err == nil {
			// 执行成功
			return nil
		} else {
			vlog.Warning(ctx, err)
		}
		time.Sleep(time.Second)
	}

	return fmt.Errorf("function call failed")
}
