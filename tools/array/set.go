package array

// TODO 扩展成泛型
type StringSet map[string]struct{}

func NewStringSet(s []string) StringSet {
	set := make(map[string]struct{})
	for i := range s {
		set[s[i]] = struct{}{}
	}
	return set
}

func (s StringSet) Contain(val string) bool {
	_, ok := s[val]
	return ok
}
