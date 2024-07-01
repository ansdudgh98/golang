package network

import (
	"fmt"
	"mygoproject/types"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit    sync.Once
	userRouterInstace *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstace = &userRouter{
			router: router,
		}

		router.registerGET("/", userRouterInstace.get)
		router.registerPOST("/", userRouterInstace.create)
		router.registerUPDATE("/", userRouterInstace.update)
		router.registerDELETE("/", userRouterInstace.delete)

	})

	return userRouterInstace
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create")

}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")

}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get")

	u.router.okResponse(c, &types.UserResponse{
		APIResponse: &types.APIResponse{
			Result:      1,
			Description: "성공입니다.",
		},
		User: nil,
	})

}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")
}
