package handler

import (
	"github.com/gofiber/fiber/v2"
	"gone/database/model"
	"gone/internal/result"
	"gone/pkg/utils"
	"log"
)

// 待办事项管理
type Todo struct{}

// 获取全部 todos
func (t *Todo) GetAllTodos(c *fiber.Ctx) error {
	var todos model.Todos
	var data = todos.FindAll()
	log.Println("data:", data)
	return c.JSON(result.Success("获取全部 todos 成功").WithData(fiber.Map{"list": data}))
}

// 更新或添加
func (t *Todo) UpdateOrAddTodo(c *fiber.Ctx) error {
	// 定义请求参数结构体
	req := struct {
		Id    int    `json:"id" form:"id"`
		Title string `json:"title" form:"title"`
		Done  bool   `json:"done" form:"done"`
	}{}
	// 绑定请求参数
	_ = c.BodyParser(&req)
	// 验证请求参数
	if errs := utils.Validator(req); errs != nil {
		return c.JSON(result.Error("请求参数错误").WithData(fiber.Map{"failed": errs}))
	}
	// 更新
	var todo model.Todos
	if req.Id > 0 {
		todo.FindOne(req.Id)
		todo.Title = req.Title
		todo.Done = req.Done
		todo.UpdateOne(req.Id)
		return c.JSON(result.Success("更新 todo 成功"))
	}
	log.Println("添加Todo")
	todo.Title = req.Title
	todo.Done = req.Done
	todo.AddOne()
	return c.JSON(result.Success("添加 todo 成功"))
}

// 删除待办事项
func (t *Todo) DeleteTodo(c *fiber.Ctx) error {
	// 获取路由参数
	idStr := c.Params("id")
	// 将字符串转换为 int
	var idInt int
	for _, v := range idStr {
		idInt = idInt*10 + int(v-'0')
	}
	if idInt == 0 {
		return c.JSON(result.Error("id 参数有误或为空"))
	}
	var todo model.Todos
	todo.DeleteOne(idInt)
	return c.JSON(result.Success("成功删除待办"))
}

// 完成待办事项
func (t *Todo) DoneTodo(c *fiber.Ctx) error {
	return c.JSON(result.Success("该接口尚未完善"))
}
