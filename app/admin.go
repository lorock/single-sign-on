package app

import (
	"runtime/debug"

	"github.com/lorock/single-sign-on/utils"
)

func ReloadConfig() {
	debug.FreeOSMemory()
	utils.LoadConfig(utils.CfgFileName)
}
