package inmemdb

import (
	"sync"

	"github.com/02_protogo/02_urlshortner/internal/types"
)

var (
	DbMapMutex sync.Mutex
	Dbmap      = make(map[string]types.UrlNode)
)


