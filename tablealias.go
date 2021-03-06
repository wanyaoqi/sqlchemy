package sqlchemy

import (
	"fmt"
	"sync"
)

var (
	tableID     = 0
	tableIDLock *sync.Mutex
)

func init() {
	tableIDLock = &sync.Mutex{}
}

func getTableAliasName() string {
	tableIDLock.Lock()
	defer tableIDLock.Unlock()
	tableID += 1
	return fmt.Sprintf("t%d", tableID)
}
