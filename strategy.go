package kuji

type KujiStrategy interface {
	PickOneByKey(s string) (string, error)
	RegisterCandidatesWithKey(s string, c []KujiCandidate) (int64, error)
}
