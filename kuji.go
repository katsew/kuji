package kuji

type Kuji struct {
	strategy KujiStrategy
}

func NewKuji(strategy KujiStrategy) Kuji {
	return Kuji{
		strategy,
	}
}

func (k Kuji) PickOneByKey(key string) (string, error) {
	return k.strategy.PickOneByKey(key)
}

func (k Kuji) PickOneByKeyAndIndex(key string, index int64) (string, error) {
	return k.strategy.PickOneByKeyAndIndex(key, index)
}

func (k Kuji) PickAndDeleteOneByKey(key string) (string, error) {
	return k.strategy.PickAndDeleteOneByKey(key)
}

func (k Kuji) RegisterCandidatesWithKey(key string, candidates []KujiCandidate) (int64, error) {
	return k.strategy.RegisterCandidatesWithKey(key, candidates)
}

func (k Kuji) Len(key string) (int64, error) {
	return k.strategy.Len(key)
}

func (k Kuji) List(key string) ([]string, error) {
	return k.strategy.List(key)
}