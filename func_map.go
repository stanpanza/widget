package widget

import (
	"sort"
	"strings"

	"github.com/qor/admin"
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
			var roles = []interface{}{}
			for _, role := range context.Roles {
				roles = append(roles, role)
			}
			if w.Permission == nil || w.Permission.HasPermission(roles.Create, roles...) {
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
			if groupedWidgetsSlice[i].Group == "" {
				return false
			}
			return strings.Compare(groupedWidgetsSlice[i].Group, groupedWidgetsSlice[j].Group) < 0
		})

		return groupedWidgetsSlice
	},
}
