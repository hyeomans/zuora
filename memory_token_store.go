package zuora

import (
	"sync"
	"time"
)

// MemoryTokenStore a simple token storer to avoid creating
// a new token on each request.
type MemoryTokenStore struct {
	token *Token
	sync.RWMutex
}

// Token Stores token in memory
func (m *MemoryTokenStore) Token() (bool, *Token) {
	m.RLock()
	defer m.RUnlock()
	if m.token != nil {
		expired := time.Now().UTC().Sub(time.Now().UTC().Add(time.Duration(m.token.ExpiresIn-100)*time.Second)) > 0
		if !expired {
			return true, m.token
		}
	}

	return false, nil
}

// Update updates token in memory
func (m *MemoryTokenStore) Update(token *Token) {
	m.Lock()
	defer m.Unlock()
	m.token = token
}
