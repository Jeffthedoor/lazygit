package context

import (
	"github.com/Jeffthedoor/generics/slices"
	"github.com/Jeffthedoor/gocui"
	"github.com/Jeffthedoor/lazygit/pkg/gui/presentation"
	"github.com/Jeffthedoor/lazygit/pkg/gui/types"
)

type MenuContext struct {
	*MenuViewModel
	*ListContextTrait
}

var _ types.IListContext = (*MenuContext)(nil)

func NewMenuContext(
	view *gocui.View,

	onFocus func(...types.OnFocusOpts) error,
	onRenderToMain func(...types.OnFocusOpts) error,
	onFocusLost func() error,

	c *types.HelperCommon,
	getOptionsMap func() map[string]string,
) *MenuContext {
	viewModel := NewMenuViewModel()

	return &MenuContext{
		MenuViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				ViewName:        "menu",
				Key:             "menu",
				Kind:            types.PERSISTENT_POPUP,
				OnGetOptionsMap: getOptionsMap,
				Focusable:       true,
			}), ContextCallbackOpts{
				OnFocus:        onFocus,
				OnFocusLost:    onFocusLost,
				OnRenderToMain: onRenderToMain,
			}),
			getDisplayStrings: viewModel.GetDisplayStrings,
			list:              viewModel,
			viewTrait:         NewViewTrait(view),
			c:                 c,
		},
	}
}

// TODO: remove this thing.
func (self *MenuContext) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.DisplayString
}

type MenuViewModel struct {
	menuItems []*types.MenuItem
	*BasicViewModel[*types.MenuItem]
}

func NewMenuViewModel() *MenuViewModel {
	self := &MenuViewModel{
		menuItems: nil,
	}

	self.BasicViewModel = NewBasicViewModel(func() []*types.MenuItem { return self.menuItems })

	return self
}

func (self *MenuViewModel) SetMenuItems(items []*types.MenuItem) {
	self.menuItems = items
}

// TODO: move into presentation package
func (self *MenuViewModel) GetDisplayStrings(_startIdx int, _length int) [][]string {
	return slices.Map(self.menuItems, func(item *types.MenuItem) []string {
		if item.DisplayStrings != nil {
			return item.DisplayStrings
		}

		styledStr := item.DisplayString
		if item.OpensMenu {
			styledStr = presentation.OpensMenuStyle(styledStr)
		}
		return []string{styledStr}
	})
}
