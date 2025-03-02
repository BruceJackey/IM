/**
* @Auth:jcx
* @Description:
* @CreateDate:2025/03/02 14:57:55
 */
package mq

import (
	"context"
	"fmt"
	"ginchat/utils"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

// TestPublish 测试发布消息到redis
func TestPublish(t *testing.T) {
	msg := "当前时间: " + time.Now().Format("15:04:05")
	err := utils.Publish(ctx, utils.PublishKey, msg)
	if err != nil {

		fmt.Println(err)
	}
}
