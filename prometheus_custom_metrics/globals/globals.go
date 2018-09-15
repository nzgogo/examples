package globals

import (
	"github.com/nzgogo/micro"
)

const (
	SrvName    = "gogo-core-user"
	SrvVersion = "v1"
)

var (
	Service = gogo.NewService(SrvName, SrvVersion)
)
