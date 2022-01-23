package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"

	"log"
)

var client redis.UniversalClient
var ctx context.Context

// 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "10.234.83.71:6379",
	})
	ctx = context.Background()
	rand.Seed(time.Now().UnixNano())
	log.Println("start")
	//setVale(10000, "1w_10", RandStringRunes(10))
	//setVale(50000, "5w_10", RandStringRunes(10))
	//setVale(500000, "50w_10", RandStringRunes(10))

	//setVale(10000, "1w_1000", RandStringRunes(1000))
	//setVale(50000, "5w_1000", RandStringRunes(1000))
	//setVale(500000, "50w_1000", RandStringRunes(1000))

	setVale(10000, "1w_5000", RandStringRunes(5000))
	log.Println("ok")
}

func setVale(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := client.Set(ctx, k, value, -1)
		err := cmd.Err()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
