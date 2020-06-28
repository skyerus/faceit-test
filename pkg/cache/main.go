package cache

import (
	"sync"

	"github.com/skyerus/faceit-test/pkg/user"
)

// Cache - application cache
type Cache struct {
	mux   sync.Mutex
	users *UsersCache
}

// UsersCache - cache of users
type UsersCache struct {
	mux   sync.Mutex
	users map[int]user.User
	keys  []int
}

const maxNumOfCachedUsers = 10000
const numOfUsersToClearOnCleanup = 1000

// InstantiateCache - init app cache
func InstantiateCache() *Cache {
	return &Cache{users: instantiateUsersCache()}
}

func instantiateUsersCache() *UsersCache {
	usersMap := make(map[int]user.User, maxNumOfCachedUsers)
	keys := make([]int, maxNumOfCachedUsers)
	return &UsersCache{users: usersMap, keys: keys}
}

// ClearCache - clear all cache
func (c *Cache) ClearCache() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.users = instantiateUsersCache()
}

// GetUsersCache - safely get users cache
func (c *Cache) GetUsersCache() *UsersCache {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.users
}

// GetUser - get user from cache
func (uc *UsersCache) GetUser(ID int) (u user.User, found bool) {
	uc.mux.Lock()
	defer uc.mux.Unlock()
	u = uc.users[ID]
	if !userFound(u) {
		return u, false
	}

	return u, true
}

// AddUser - add user to cache (if cache is maxed out some older users will be removed)
func (uc *UsersCache) AddUser(u user.User) {
	uc.mux.Lock()
	defer uc.mux.Unlock()
	cachedUser := uc.users[u.ID]
	if userFound(cachedUser) {
		uc.shiftIDToEnd(u.ID)
		uc.users[u.ID] = u
	} else {
		if len(uc.users) == maxNumOfCachedUsers {
			uc.cleanup()
		}
		uc.users[u.ID] = u
		uc.keys = append(uc.keys, u.ID)
	}
}

// DeleteUser - delete user from cache store
func (uc *UsersCache) DeleteUser(ID int) {
	uc.mux.Lock()
	defer uc.mux.Unlock()
	u := uc.users[ID]
	if !userFound(u) {
		return
	}
	delete(uc.users, ID)
	var i int
	for i = 0; i < len(uc.keys); i++ {
		if uc.keys[i] == ID {
			break
		}
	}
	uc.keys = append(uc.keys[:i], uc.keys[i+1:]...)
}

func (uc *UsersCache) cleanup() {
	for i := 0; i < min(maxNumOfCachedUsers, numOfUsersToClearOnCleanup); i++ {
		delete(uc.users, uc.keys[i])
	}
	uc.keys = uc.keys[:min(maxNumOfCachedUsers, numOfUsersToClearOnCleanup)]
}

func min(first int, second int) int {
	if first > second {
		return second
	}

	return first
}

func (uc *UsersCache) shiftIDToEnd(ID int) {
	var i int
	for i = 0; i < len(uc.keys); i++ {
		if uc.keys[i] == ID {
			break
		}
	}
	uc.keys = append(uc.keys[:i], uc.keys[i+1:]...)
	uc.keys = append(uc.keys, ID)
}

func userFound(u user.User) bool {
	return u.ID != 0
}
