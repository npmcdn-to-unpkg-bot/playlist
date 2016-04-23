// Tunes web service.
package main

import (
	"github.com/sath33sh/infra/config"
	"github.com/sath33sh/infra/db"
	"github.com/sath33sh/infra/log"
	"github.com/sath33sh/infra/wapi"
	"github.com/sath33sh/tunes/rest"
)

func initFileServer() {
	// Init web file server endpoint.
	wapi.ServeFiles("/app/*filepath", config.Base.GetString("tunes", "root", "."))
}

func main() {
	// Initialize the infra modules, in order.
	config.Init("etc/base.conf")
	log.Init(config.Base.GetString("tunes", "log", ""))
	db.Init()

	// Initialize REST endpoints.
	rest.InitSong()

	// Initialize file server.
	initFileServer()

	// Start CAS.
	port := config.Base.GetInt("tunes", "port", 80)
	secure := config.Base.GetBool("tunes", "secure", true)
	log.Infoln("Starting Tunes web service on port", port)
	wapi.StartServer(port, secure,
		config.Base.GetString("tunes", "cert-file", ""),
		config.Base.GetString("tunes", "key-file", ""))
}
