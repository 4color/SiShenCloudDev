package main

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/initapp"
)

func main() {
	initapp.InitRouter()
	initapp.R.Run()
}