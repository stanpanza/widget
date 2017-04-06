package widget

import (
	"sort"
	"strings"

	"github.com/qor/admin"
	"github.com/qor/roles"
)

type GroupedWidgets struct {
	Group   string
	Widgets []*Widget
}

var funcMap = map[string]interface{}{
	"widget_available_scopes": func() []*Scope {
		if len(registeredScopes) > 0 {
			return append([]*Scope{{Name: "Default Visitor", Param: "default"}}, registeredScopes...)
		}
		return []*Scope{}
	},
	"widget_grouped_widgets": func(context *admin.Context) []*GroupedWidgets {
		groupedWidgetsSlice := []*GroupedWidgets{}

	OUTER:
		for _, w := range registeredWidgets {
			if w.Permission.HasPermission(roles.Create, context.Context.Roles...) {
				for _, groupedWidgets := range groupedWidgetsSlice {
					if groupedWidgets.Group == w.Group {
						groupedWidgets.Widgets = append(groupedWidgets.Widgets, w)
						continue OUTER
					}
				}

				groupedWidgetsSlice = append(groupedWidgetsSlice, &GroupedWidgets{
					Group:   w.Group,
					Widgets: []*Widget{w},
				})
			}
		}

		sort.SliceStable(groupedWidgetsSlice, func(i, j int) bool {
			return strings.Compare(groupedWidgetsSlice[i].Group, groupedWidgetsSlice[j].Group) > 0
		})

		return groupedWidgetsSlice
	},
}
