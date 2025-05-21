package game

import "sync"

// Manager holds the current game state in memory.
type Manager struct {
	mu    sync.RWMutex
	state State
}

// NewManager returns a Manager with a default state.
func NewManager() *Manager {
	m := &Manager{}
	m.state = State{}
	return m
}

// State returns a copy of the current game state.
func (m *Manager) State() State {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.state
}

// SetState replaces the game state. Useful for loading saves.
func (m *Manager) SetState(s State) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.state = s
}
