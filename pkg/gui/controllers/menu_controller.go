package controllers

import (
	"github.com/Jeffthedoor/lazygit/pkg/gui/context"
	"github.com/Jeffthedoor/lazygit/pkg/gui/types"
)

type MenuController struct {
	baseController
	*controllerCommon
}

var _ types.IController = &MenuController{}

func NewMenuController(
	common *controllerCommon,
) *MenuController {
	return &MenuController{
		baseController:   baseController{},
		controllerCommon: common,
	}
}

func (self *MenuController) GetKeybindings(opts types.KeybindingsOpts) []*types.Binding {
	bindings := []*types.Binding{
		{
			Key:     opts.GetKey(opts.Config.Universal.Select),
			Handler: self.press,
		},
		{
			Key:     opts.GetKey(opts.Config.Universal.Confirm),
			Handler: self.press,
		},
		{
			Key:     opts.GetKey(opts.Config.Universal.ConfirmAlt1),
			Handler: self.press,
		},
	}

	return bindings
}

func (self *MenuController) GetOnClick() func() error {
	return self.press
}

func (self *MenuController) press() error {
	selectedItem := self.context().GetSelected()

	if err := self.c.PopContext(); err != nil {
		return err
	}

	if err := selectedItem.OnPress(); err != nil {
		return err
	}

	return nil
}

func (self *MenuController) Context() types.Context {
	return self.context()
}

func (self *MenuController) context() *context.MenuContext {
	return self.contexts.Menu
}
