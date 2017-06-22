package kuji

import (
	"errors"
)

type kuji struct {
	strategies map[string]KujiStrategyConfig
}

func NewKuji() *kuji {
	return &kuji{
		make(map[string]KujiStrategyConfig),
	}
}

func (k *kuji) Use(strategyName string, strategyConfig KujiStrategyConfig) error {
	if k.strategies[strategyName].Strategy != nil {
		return errors.New("This strategy has already assigned!")
	}
	k.strategies[strategyName] = strategyConfig
	return nil
}

func (k *kuji) PickOneByKey(strategyName string, key string) (string, error) {

	if ok, error := k.strategies[strategyName].Strategy.PickOneByKey(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.PickOneByKey(key)
		}
		return "", error
	} else {
		return ok, nil
	}

}

func (k *kuji) PickOneByKeyAndIndex(strategyName string, key string, index int64) (string, error) {

	if ok, error := k.strategies[strategyName].Strategy.PickOneByKeyAndIndex(key, index); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.PickOneByKeyAndIndex(key, index)
		}
		return "", error
	} else {
		return ok, nil
	}

}

func (k *kuji) PickAndDeleteOneByKey(strategyName string, key string) (string, error) {

	if ok, error := k.strategies[strategyName].Strategy.PickAndDeleteOneByKey(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.PickAndDeleteOneByKey(key)
		}
		return "", error
	} else {
		return ok, nil
	}

}

func (k *kuji) RegisterCandidatesWithKey(strategyName string, key string, candidates []KujiCandidate) (int64, error) {

	if ok, error := k.strategies[strategyName].Strategy.RegisterCandidatesWithKey(key, candidates); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.RegisterCandidatesWithKey(key, candidates)
		}
		return 0, error
	} else {
		return ok, nil
	}

}

func (k *kuji) Len(strategyName string, key string) (int64, error) {

	if ok, error := k.strategies[strategyName].Strategy.Len(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.Len(key)
		}
		return 0, error
	} else {
		return ok, nil
	}

}

func (k *kuji) List(strategyName string, key string) ([]string, error) {

	if ok, error := k.strategies[strategyName].Strategy.List(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.List(key)
		}
		return []string{}, error
	} else {
		return ok, nil
	}

}
