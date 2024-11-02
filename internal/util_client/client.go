package util_client

import (
	"github.com/metadiv-go-tech/metadiv_util_client"
	"github.com/metadiv-go-tech/metagin/v2"
)

var (
	UtilClientHost = metagin.Environment("UTIL_CLIENT_HOST", true)
	StaticSpace    = metagin.Environment("STATIC_SPACE", true)
)

var Client = metadiv_util_client.NewClient(UtilClientHost.String())
