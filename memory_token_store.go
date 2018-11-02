package zuora

import (
	"sync"
	"time"
)

type memoryTokenStore struct{}

var cachedToken *Token
var mutex sync.RWMutex

func (m *memoryTokenStore) Token() (bool, *Token) {
	mutex.RLock()
	defer mutex.RUnlock()
	if cachedToken != nil {
		expired := time.Now().UTC().Sub(time.Now().UTC().Add(time.Duration(cachedToken.ExpiresIn)*time.Second)) > 0
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
}
