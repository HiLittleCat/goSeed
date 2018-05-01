package routers

import (
	"github.com/HiLittleCat/goSeed/controller"
)

// Init routers init
func Init() {
	(&controller.User{}).Register()
}
