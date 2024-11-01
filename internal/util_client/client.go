package util_client

import (
	"github.com/metadiv-go-tech/metadiv_util_client"
	"github.com/metadiv-go-tech/metagin"
)

var (
	UtilClientHost = metagin.Environment.String("UTIL_CLIENT_HOST", true)
	StaticSpace    = metagin.Environment.String("STATIC_SPACE", true)
)

var Client = metadiv_util_client.NewClient(UtilClientHost)
