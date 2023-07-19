package ippool

type IPSet struct {
	ipaddress map[string]struct{}
}

var exists = struct{}{}

func NewSet() *IPSet {
	s := &IPSet{}
	s.ipaddress = make(map[string]struct{})
	return s
}

func (s *IPSet) Add(value string) {
	s.ipaddress[value] = exists
}

func (s *IPSet) Remove(value string) {
	delete(s.ipaddress, value)
}

func (s *IPSet) Contains(value string) bool {
	_, c := s.ipaddress[value]
	return c
}

func (s *IPSet) GetAll() []string {
	keys := make([]string, 0, len(s.ipaddress))
	for k := range s.ipaddress {
		keys = append(keys, k)
	}
	return keys
}

func (s *IPSet) IsEqual(Pair *IPSet) (isEqual bool, added, deleted IPSet) {
	isEqual = true
	for values := range s.ipaddress {
		c := Pair.Contains(values)
		if !c {
			isEqual = false
			deleted.Add(values)
		}
	}

	for values := range Pair.ipaddress {
		c := s.Contains(values)
		if !c {
			isEqual = false
			added.Add(values)
		}
	}

	return isEqual, added, deleted
}
