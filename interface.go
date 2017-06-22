package kuji

type Kuji interface {
	Use(string, KujiStrategyConfig) error
	RegisterCandidatesWithKey(string, string, []KujiCandidate) (int64, error)
	PickOneByKey(string, string) (string, error)
	PickOneByKeyAndIndex(string, string, int64) (string, error)
	PickAndDeleteOneByKey(string, string) (string, error)
	Len(string, string) (int64, error)
	List(string, string) ([]string, error)
}
