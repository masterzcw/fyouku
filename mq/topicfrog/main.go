package main

import (
	"demo/services/mq"
	"fmt"
)

// 主题模式的消费者
func main() {
	mq.ConsumerEx("fyouku.demo.topic", "topic", "*.frog.*", callback)
}

func callback(s string) {
	fmt.Printf("msg is :%s\n", s)
}
