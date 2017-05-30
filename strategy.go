package kuji

type KujiStrategy interface {
	PickOneByKey(string) (string, error)
	PickOneByKeyAndIndex(string, int64) (string, error)
	PickAndDeleteOneByKey(string) (string, error)
	RegisterCandidatesWithKey(string, []KujiCandidate) (int64, error)
}
