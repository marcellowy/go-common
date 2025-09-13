package snowflake

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node    *snowflake.Node
	once    sync.Once
	initErr error
)

// Init 初始化雪花节点，只能调用一次
// machineID 范围 0 ~ 1023
func Init(machineID int64) error {
	once.Do(func() {
		if machineID < 0 || machineID > 1023 {
			initErr = fmt.Errorf("machineID must be between 0 and 1023")
			return
		}
		node, initErr = snowflake.NewNode(machineID)
	})
	return initErr
}

// NextID 生成唯一 ID (int64)
// 注意：Init 必须先调用成功
func NextID() (int64, error) {
	if node == nil {
		return 0, fmt.Errorf("snowflake node is not initialized, call Init() first")
	}
	return node.Generate().Int64(), nil
}

// NextIDString 生成唯一 ID 的字符串形式
func NextIDString() (string, error) {
	if node == nil {
		return "", fmt.Errorf("snowflake node is not initialized, call Init() first")
	}
	return node.Generate().String(), nil
}
