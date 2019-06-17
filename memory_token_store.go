package zuora

import (
	"sync"
	"time"
)

type MemoryTokenStore struct {
	token *Token
	sync.RWMutex
}

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

func (m *MemoryTokenStore) Update(token *Token) {
	m.Lock()
	defer m.Unlock()
	m.token = token
}
