package likes

import (
	"sync"
)

type LikeStorage struct {
	likes int
	sync.RWMutex
}

func (like *LikeStorage) GetLikes() {
	like.RWMutex.Lock()
	_ = like.likes
	like.RWMutex.Unlock()
}

func (like *LikeStorage) AddLikes(someLikes int) {
	like.RWMutex.Lock()
	like.likes += someLikes
	like.RWMutex.Unlock()
}
