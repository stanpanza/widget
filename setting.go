package widget

import (
	"fmt"
	"time"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/serializable_meta"
)

// QorWidgetSettingInterface qor widget setting interface
type QorWidgetSettingInterface interface {
	GetWidgetType() string
	SetWidgetType(string)
	GetWidgetName() string
	SetWidgetName(string)
	GetGroupName() string
	SetGroupName(string)
	GetScope() string
	SetScope(string)
	GetTemplate() string
	SetTemplate(string)
	serializable_meta.SerializableMetaInterface
}

// QorWidgetSetting default qor widget setting struct
type QorWidgetSetting struct {
	Name       string `gorm:"primary_key"`
	WidgetType string `gorm:"primary_key;size:128"`
	Scope      string `gorm:"primary_key;size:128;default:'default'"`
	GroupName  string
	Template   string
	serializable_meta.SerializableMeta
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ResourceName get widget setting's resource name
func (widgetSetting *QorWidgetSetting) ResourceName() string {
	return "Widget Content"
}

// GetSerializableArgumentKind get serializable kind
func (widgetSetting *QorWidgetSetting) GetSerializableArgumentKind() string {
	if widgetSetting.WidgetType != "" {
		return widgetSetting.WidgetType
	}
	return widgetSetting.Kind
}

// SetSerializableArgumentKind set serializable kind
func (widgetSetting *QorWidgetSetting) SetSerializableArgumentKind(name string) {
	widgetSetting.WidgetType = name
	widgetSetting.Kind = name
}

// GetWidgetType get widget setting's type
func (widgetSetting QorWidgetSetting) GetWidgetType() string {
	return widgetSetting.WidgetType
}

// SetWidgetType set widget setting's type
func (widgetSetting *QorWidgetSetting) SetWidgetType(widgetType string) {
	widgetSetting.WidgetType = widgetType
}

// GetWidgetName get widget setting's group name
func (widgetSetting QorWidgetSetting) GetWidgetName() string {
	return widgetSetting.Name
}

// SetWidgetName set widget setting's group name
func (widgetSetting *QorWidgetSetting) SetWidgetName(name string) {
	widgetSetting.Name = name
}

// GetGroupName get widget setting's group name
func (widgetSetting QorWidgetSetting) GetGroupName() string {
	return widgetSetting.GroupName
}

// SetGroupName set widget setting's group name
func (widgetSetting *QorWidgetSetting) SetGroupName(groupName string) {
	widgetSetting.GroupName = groupName
}

// GetScope get widget's scope
func (widgetSetting QorWidgetSetting) GetScope() string {
	return widgetSetting.Scope
}

// SetScope set widget setting's scope
func (widgetSetting *QorWidgetSetting) SetScope(scope string) {
	widgetSetting.Scope = scope
}

// GetTemplate get used widget template
func (widgetSetting QorWidgetSetting) GetTemplate() string {
	if widget := GetWidget(widgetSetting.GetSerializableArgumentKind()); widget != nil {
		for _, value := range widget.Templates {
			if value == widgetSetting.Template {
				return value
			}
		}

		// return first value of defined widget templates
		for _, value := range widget.Templates {
			return value
		}
	}
	return ""
}

// SetTemplate set used widget's template
func (widgetSetting *QorWidgetSetting) SetTemplate(template string) {
	widgetSetting.Template = template
}

// GetSerializableArgumentResource get setting's argument's resource
func (widgetSetting *QorWidgetSetting) GetSerializableArgumentResource() *admin.Resource {
	widget := GetWidget(widgetSetting.GetSerializableArgumentKind())
	if widget != nil {
		return widget.Setting
	}
	return nil
}

// ConfigureQorResource a method used to config Widget for qor admin
func (widgetSetting *QorWidgetSetting) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		if res.GetMeta("Name") == nil {
			res.Meta(&admin.Meta{Name: "Name"})
		}

		res.Meta(&admin.Meta{
			Name: "Scope",
			Type: "hidden",
			Valuer: func(result interface{}, context *qor.Context) interface{} {
				if scope := context.Request.URL.Query().Get("widget_scope"); scope != "" {
					return scope
				}

				if setting, ok := result.(QorWidgetSettingInterface); ok {
					if scope := setting.GetScope(); scope != "" {
						return scope
					}
				}

				return "default"
			},
			Setter: func(result interface{}, metaValue *resource.MetaValue, context *qor.Context) {
				if setting, ok := result.(QorWidgetSettingInterface); ok {
					setting.SetScope(utils.ToString(metaValue.Value))
				}
			},
		})

		res.Meta(&admin.Meta{
			Name: "Widgets",
			Type: "select_one",
			Valuer: func(result interface{}, context *qor.Context) interface{} {
				if typ := context.Request.URL.Query().Get("widget_type"); typ != "" {
					return typ
				}

				if setting, ok := result.(QorWidgetSettingInterface); ok {
					widget := GetWidget(setting.GetSerializableArgumentKind())
					if widget == nil {
						return ""
					}
					return widget.Name
				}

				return ""
			},
			Collection: func(result interface{}, context *qor.Context) (results [][]string) {
				if setting, ok := result.(QorWidgetSettingInterface); ok {
					if setting.GetWidgetName() == "" {
						for _, widget := range registeredWidgets {
							results = append(results, []string{widget.Name, widget.Name})
						}
					} else {
						groupName := setting.GetGroupName()
						for _, group := range registeredWidgetsGroup {
							if group.Name == groupName {
								for _, widget := range group.Widgets {
									results = append(results, []string{widget, widget})
								}
							}
						}
					}

					if len(results) == 0 {
						results = append(results, []string{setting.GetSerializableArgumentKind(), setting.GetSerializableArgumentKind()})
					}
				}
				return
			},
			Setter: func(result interface{}, metaValue *resource.MetaValue, context *qor.Context) {
				if setting, ok := result.(QorWidgetSettingInterface); ok {
					setting.SetSerializableArgumentKind(utils.ToString(metaValue.Value))
				}
			},
		})

		res.Meta(&admin.Meta{
			Name: "Template",
			Type: "select_one",
			Valuer: func(result interface{}, context *qor.Context) interface{} {
				if setting, ok := result.(QorWidgetSettingInterface); ok {
					return setting.GetTemplate()
				}
				return ""
			},
			Collection: func(result interface{}, context *qor.Context) (results [][]string) {
				if setting, ok := result.(QorWidgetSettingInterface); ok {
					if widget := GetWidget(setting.GetSerializableArgumentKind()); widget != nil {
						for _, value := range widget.Templates {
							results = append(results, []string{value, value})
						}
					}
				}
				return
			},
			Setter: func(result interface{}, metaValue *resource.MetaValue, context *qor.Context) {
				if setting, ok := result.(QorWidgetSettingInterface); ok {
					setting.SetTemplate(utils.ToString(metaValue.Value))
				}
			},
		})

		res.Action(&admin.Action{
			Name: "Preview",
			URL: func(record interface{}, context *admin.Context) string {
				return fmt.Sprintf("%v/%v/%v/!preview", context.Admin.GetRouter().Prefix, res.ToParam(), record.(QorWidgetSettingInterface).GetWidgetName())
			},
			Modes: []string{"edit", "menu_item"},
		})

		res.UseTheme("widget")

		res.IndexAttrs("Name", "CreatedAt", "UpdatedAt")
		res.ShowAttrs("Name", "Scope", "WidgetType", "Template", "Value", "CreatedAt", "UpdatedAt")
		res.EditAttrs(
			"Scope", "Widgets", "Template",
			&admin.Section{
				Title: "Settings",
				Rows:  [][]string{{"Kind"}, {"SerializableMeta"}},
			},
		)
		res.NewAttrs("Name", "Scope", "Widgets", "Template")
	}
}
