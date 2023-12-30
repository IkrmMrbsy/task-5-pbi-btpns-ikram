
package main

import (
	"PBIP_BTPN_SYARIAH/app"
	"PBIP_BTPN_SYARIAH/router"
)

func main() {
	app.Init()
	r := router.SetupRouter()
	r.Run(":0")
}
