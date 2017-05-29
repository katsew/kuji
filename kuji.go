package kuji

type Kuji struct {
	strategy KujiStrategy
}

func NewKuji (strategy KujiStrategy) (Kuji) {
	return Kuji{
		strategy,
	}
}

func (k Kuji) PickOneByKey(s string) (string, error) {
	return k.strategy.PickOneByKey(s)
}

func (k Kuji) PickAndDeleteOneByKey(s string) (string, error) {
	return k.strategy.PickAndDeleteOneByKey(s)
}

func (k Kuji) RegisterCandidatesWithKey(s string, c []KujiCandidate) (int64, error) {
	return k.strategy.RegisterCandidatesWithKey(s, c)
}