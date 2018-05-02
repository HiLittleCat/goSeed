package routers

import (
	"github.com/HiLittleCat/goSeed/controller"
)

// Init routers init
func init() {
	(&controller.User{}).Register()
}
