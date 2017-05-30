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
	kr := kuji_redis.NewShuffleStrategy(&opt)
	instance := kuji.NewKuji(kr)

	c := []kuji.KujiCandidate{
		kuji.KujiCandidate{
			Id: 1,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id: 2,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id: 3,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id: 4,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id: 5,
			Weight: 1,
		},
		kuji.KujiCandidate{
			Id: 6,
			Weight: 1,
		},
	}

	_, err := instance.RegisterCandidatesWithKey("shuffle", c)
	if (err != nil) {
		panic(err)
	}

	PickOne(instance)

}

func PickOne(k kuji.Kuji) {
	for i := 0; i < 1000; i++ {
		num, err := k.PickOneByKey("shuffle")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Pickup string is", num)
	}
}

func PickAndDeleteOne(k kuji.Kuji) {
	for i := 0; i < 1000; i++ {
		num, err := k.PickAndDeleteOneByKey("shuffle")
		if err != nil {
			fmt.Println("No data stored in redis!")
			break
		}
		fmt.Println("Pickup string is", num)
	}
}