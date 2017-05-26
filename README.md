# Abstract

This is a lottery library with strategy pattern.

# Usage

First, we save lottery data into some data store (e.g. redis). 

```go
kr := kuji_redis.NewKujiRedisStrategy(&opt)
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

_, err := instance.RegisterCandidatesWithKey("foo", c)

```

Then, pick random one.

```go
num, err := k.PickOneByKey("foo")
if err != nil {
    panic(err)
}
fmt.Println("Pickup string is", num)
```

*this is from example directory, please see it for the detail if you like :)