package go_map

import "container/list"

type Map struct {
	l *list.List
}

func New() *Map {
	return &Map{list.New()}
}
func (m *Map) Set(key string, value interface{}) {
	if _, ok, e := m.Get(key); !ok {
		m.l.PushBack([]interface{}{key, value})
	} else {
		e.Value = []interface{}{key, value}
	}
}
func (m *Map) Get(key string) (d interface{}, b bool, l *list.Element) {
	m.traverse(func(data []interface{}, e *list.Element) (exit bool) {
		if data[0].(string) == key {
			d = data[1]
			b = true
			l = e
			return true
		}
		return
	})
	return nil, false, nil
}
func (m *Map) Delete(key string) bool {
	if _, ok, e := m.Get(key); ok {
		m.l.Remove(e)
		return true
	}
	return false
}
func (m *Map) Pull2Map() map[string]interface{} {
	out := make(map[string]interface{})
	m.traverse(func(data []interface{}, e *list.Element) (exit bool) {
		out[data[0].(string)] = data[1]
		return
	})
	return out
}
func (m *Map) Pull2List() ([]string, []interface{}) {
	keys := make([]string, 0, m.l.Len())
	values := make([]interface{}, 0, m.l.Len())
	m.traverse(func(data []interface{}, e *list.Element) (exit bool) {
		keys = append(keys, data[0].(string))
		values = append(values, data[1])
		return
	})

	return keys, values
}

func (m *Map) traverse(f func(data []interface{}, e *list.Element) (exit bool)) {
	var data []interface{}
	for e := m.l.Front(); e != nil; e = e.Next() {
		data = e.Value.([]interface{})
		if exit := f(data, e); exit {
			return
		}
	}
}
