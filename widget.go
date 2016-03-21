package widget

import (
	"fmt"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

type WidgetInstance struct {
	Config *qor.Config
}

func New(config *qor.Config) *WidgetInstance {
	return &WidgetInstance{Config: config}
}

type Widget struct {
	Name     string
	Template string
	Setting  *admin.Resource
	Context  func(context Context, setting interface{}) Context
}

var registeredWidgets []*Widget
var viewPaths = []string{}

func (widgetInstance *WidgetInstance) RegisterWidget(w *Widget) {
	registeredWidgets = append(registeredWidgets, w)
}

func GetWidget(name string) (w Widget, err error) {
	for _, w := range registeredWidgets {
		if w.Name == name {
			return *w, nil
		}
	}
	return Widget{}, fmt.Errorf("Widget: failed to find widget %v", name)
}
