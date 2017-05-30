package main

import (
	"github.com/katsew/kuji"
	"github.com/katsew/kuji-redis"
	"github.com/go-redis/redis"
	"fmt"
)

func main () {
	opt := redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	}
	kr := kuji_redis.NewSimpleStrategy(&opt)
	instance := kuji.NewKuji(kr)

	c := []kuji.KujiCandidate{
		kuji.KujiCandidate{
			Id: 1,
			Weight: 5,
		},
		kuji.KujiCandidate{
			Id: 2,
			Weight: 15,
		},
		kuji.KujiCandidate{
			Id: 3,
			Weight: 80,
		},
	}

	_, err := instance.RegisterCandidatesWithKey("simple", c)
	if (err != nil) {
		panic(err)
	}

	PickOne(instance)

}

func PickOne(k kuji.Kuji) {
	for i := 0; i < 1000; i++ {
		num, err := k.PickOneByKey("simple")
		if err != nil {
			panic(err)
		}
		fmt.Println("Pickup string is", num)
	}
}