package main

import (
	"time"

	"github.com/bmclab-git/v2ray-core/v5/main/v2binding"
)

func main() {
	v2binding.StartAPIInstance(".")
	for {
		time.Sleep(time.Minute)
	}
}
