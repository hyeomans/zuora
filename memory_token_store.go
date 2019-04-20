package zuora

import (
	"sync"
	"time"
)

type memoryTokenStore struct{}

var cachedToken *Token
var tokenExpiry time.Time
var mutex sync.RWMutex

func (m *memoryTokenStore) Token() (bool, *Token) {
	mutex.RLock()
	defer mutex.RUnlock()
	if cachedToken != nil {
		expired := time.Now().UTC().After(tokenExpiry)
		if !expired {
			return true, cachedToken
		}
	}

	return false, nil
}

func (m *memoryTokenStore) Update(token *Token) {
	mutex.Lock()
	defer mutex.Unlock()
	cachedToken = token
	tokenExpiry = time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
}
