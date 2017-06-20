package kuji

import "errors"

type Kuji struct {
	strategies map[string]KujiStrategyConfig
}

func NewKuji(strategyName string, config KujiStrategyConfig) Kuji {

	k := Kuji{
		make(map[string]KujiStrategyConfig),
	}
	if strategyName != "" && config.Strategy != nil {
		k.strategies[strategyName] = config
	}
	return k

}

func (k Kuji) Use(strategyName string, strategyConfig KujiStrategyConfig) error {
	if k.strategies[strategyName].Strategy != nil {
		return errors.New("This strategy has already assigned!")
	}
	k.strategies[strategyName] = strategyConfig
	return nil
}

func (k Kuji) PickOneByKey(strategyName string, key string) (string, error) {

	if ok, error := k.strategies[strategyName].Strategy.PickOneByKey(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.PickOneByKey(key)
		}
		return "", error
	} else {
		return ok, nil
	}

}

func (k Kuji) PickOneByKeyAndIndex(strategyName string, key string, index int64) (string, error) {

	if ok, error := k.strategies[strategyName].Strategy.PickOneByKeyAndIndex(key, index); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.PickOneByKeyAndIndex(key, index)
		}
		return "", error
	} else {
		return ok, nil
	}

}

func (k Kuji) PickAndDeleteOneByKey(strategyName string, key string) (string, error) {

	if ok, error := k.strategies[strategyName].Strategy.PickAndDeleteOneByKey(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.PickAndDeleteOneByKey(key)
		}
		return "", error
	} else {
		return ok, nil
	}

}

func (k Kuji) RegisterCandidatesWithKey(strategyName string, key string, candidates []KujiCandidate) (int64, error) {

	if ok, error := k.strategies[strategyName].Strategy.RegisterCandidatesWithKey(key, candidates); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.RegisterCandidatesWithKey(key, candidates)
		}
		return 0, error
	} else {
		return ok, nil
	}

}

func (k Kuji) Len(strategyName string, key string) (int64, error) {

	if ok, error := k.strategies[strategyName].Strategy.Len(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.Len(key)
		}
		return 0, error
	} else {
		return ok, nil
	}

}

func (k Kuji) List(strategyName string, key string) ([]string, error) {

	if ok, error := k.strategies[strategyName].Strategy.List(key); error != nil {
		if k.strategies[strategyName].FailOver != nil {
			return k.strategies[strategyName].FailOver.List(key)
		}
		return []string{}, error
	} else {
		return ok, nil
	}

}
