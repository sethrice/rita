package useragent

import (
	"github.com/activecm/rita/pkg/data"
	"github.com/activecm/rita/pkg/host"
)

// Repository for uconn collection
type Repository interface {
	CreateIndexes() error
	Upsert(useragentMap map[string]*Input, hostMap map[string]*host.Input)
}

// Input ....
type Input struct {
	Name     string
	Seen     int64
	OrigIps  data.UniqueIPSet
	Requests data.StringSet
	JA3      bool
}

// Result represents a user agent and how many times that user agent
// was seen in the dataset
type Result struct {
	UserAgent string `bson:"user_agent"`
	TimesUsed int64  `bson:"seen"`
}
