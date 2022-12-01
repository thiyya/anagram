package processor

import (
	"strings"
	"sync"
)

type SafeMap struct {
	safeMap map[string][]string
	mutex   sync.RWMutex
}

var (
	once     sync.Once
	instance SafeMap // singleton
)

func New() *SafeMap {
	once.Do(func() { // atomic, does not allow repeating
		instance = SafeMap{}
		instance.safeMap = make(map[string][]string, 0)
	})

	return &instance
}

func (s *SafeMap) get(key string) string {
	value, ok := s.safeMap[key]
	if !ok {
		return ""
	}
	if len(value) > 1 {
		return strings.Join(value, ",")
	}

	return ""
}

func (s *SafeMap) add(key string, value string) {
	s.mutex.Lock()
	val := s.safeMap[key]
	val = append(val, value)
	s.safeMap[key] = val
	s.mutex.Unlock()
}

func (s *SafeMap) flush() {
	s.safeMap = make(map[string][]string, 0)
}
