package apis

import (
	"github.com/gofiber/fiber/v2"
	"gone/middle"
)

// Api 路由组，访问以下所有路由都需加上 /api
func Api(api fiber.Router) {
	api.Get("/", hello) // 保留的路由，用以验活

	apiV1 := api.Group("/v1", middle.Auth()) // api/v1 路由组

	var u *user                        // 用户管理
	apiV1.Post("/user/login", u.login) // 登录
	apiV1.Post("/user/sth", u.postSth) // 仅作演示

	var t *todo                           // 待办事项管理
	todos := api.Group("todos")           // api/todos 路由组
	todos.Get("/all", t.getAllTodos)      // api/todos/all 获取全部 todos
	todos.Post("/one", t.updateOrAddTodo) // api/todos/one 更新或添加
	todos.Delete("/:id", t.deleteTodo)    // api/todos/:id 删除待办事项
	todos.Post("/done", t.doneTodo)       // api/todos/done 完成待办事项
}

// 服务端 api 路由
func hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("👊 Yes, Iam working!")
}
