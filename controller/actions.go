package controller

import (
	"../app"
	"github.com/goadesign/goa"
)

// ActionsController implements the actions resource.
type ActionsController struct {
	*goa.Controller
}

// NewActionsController creates a actions controller.
func NewActionsController(service *goa.Service) *ActionsController {
	return &ActionsController{Controller: service.NewController("ActionsController")}
}

// ID runs the ID action.
func (c *ActionsController) ID(ctx *app.IDActionsContext) error {
	// ActionsController_ID: start_implement
	if ctx.ID == 0 {
		return ctx.NotFound()
	}

	res := &app.Integer{}
	// Put your logic here
	res.ID = ctx.ID
	return ctx.OK(res)
	// ActionsController_ID: end_implement
}

// Hello runs the hello action.
func (c *ActionsController) Hello(ctx *app.HelloActionsContext) error {
	// ActionsController_Hello: start_implement
	name := ctx.Name

	res := &app.Message{}
	res.Message = "Hello " + name

	return ctx.OK(ctx)
	// ActionsController_Hello: end_implement
}

// Ping runs the ping action.
func (c *ActionsController) Ping(ctx *app.PingActionsContext) error {
	// ActionsController_Ping: start_implement
	message := "pong"
	// Put your logic here
	res := &app.Message{}
	res.Message = message
	return ctx.OK(res)
	// ActionsController_Ping: end_implement
}
