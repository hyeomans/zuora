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
	expiration time.Time
}

// Token Stores token in memory
func (m *MemoryTokenStore) Token() (bool, *Token) {
	m.RLock()
	defer m.RUnlock()
	if m.token != nil {
		expired := time.Now().UTC().After(m.expiration)
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
	m.expiration = time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
	m.token = token
}
