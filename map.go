package go_map

type Map struct {
	m     map[string]interface{}
	l     []string
	index int
}

func New() *Map {
	return &Map{map[string]interface{}{}, []string{}, 0}
}
func (m *Map) Set(key string, value interface{}) {
	if _, ok := m.m[key]; !ok {
		m.l = append(m.l, key)
		m.index++
	}
	m.m[key] = value

}
func (m *Map) Get(key string) (interface{}, bool) {
	v, b := m.m[key]
	return v, b
}
func (m *Map) Delete(key string) {
	delete(m.m, key)
	for i, k := range m.l {
		if k == key {
			if i == 0 {
				m.l = m.l[1:]
			} else if i == len(m.l)-1 {
				m.l = m.l[:i]
			} else {
				m.l = append(m.l[:i], m.l[i+1:]...)
			}
		}
	}
	m.index--
}
func (m *Map) Pull2Map() map[string]interface{} {
	return m.m
}
func (m *Map) Pull2List() ([]string, []interface{}) {
	keys := make([]string, 0, m.Len())
	values := make([]interface{}, 0, m.Len())
	for _, k := range m.l {
		keys = append(keys, k)
		values = append(values, m.m[k])
	}
	return keys, values
}
func (m *Map) Len() int {
	return m.index
}
