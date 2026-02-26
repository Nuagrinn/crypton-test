package main

import "sync"

type MyCystomMap struct {
	m           map[int]int
	mu          sync.RWMutex
	CallKeysCnt int64
	AddKeysCnt  int64
}

func NewMyCustomMap() *MyCystomMap {
	return &MyCystomMap{
		m: make(map[int]int),
	}
}

func (m *MyCystomMap) Lock()   { m.mu.Lock() }
func (m *MyCystomMap) Unlock() { m.mu.Unlock() }

func (m *MyCystomMap) RLock()   { m.mu.RLock() }
func (m *MyCystomMap) RUnlock() { m.mu.RUnlock() }
