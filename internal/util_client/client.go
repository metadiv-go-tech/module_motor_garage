package util_client

import (
	"github.com/metadiv-go-tech/metadiv_services_client/html_to_pdf_client"
	"github.com/metadiv-go-tech/metadiv_services_client/static_client"
	"github.com/metadiv-go-tech/metagin/v2"
)

var (
	UtilClientHost = metagin.Environment("UTIL_CLIENT_HOST", true)
	StaticSpace    = metagin.Environment("STATIC_SPACE", true)
)

var StaticClient = static_client.New(UtilClientHost.String(), StaticSpace.String())

var HtmlToPdfClient = html_to_pdf_client.New(UtilClientHost.String())
