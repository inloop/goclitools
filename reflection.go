package goclitools

import (
	"reflect"
	"strings"
)

// ReflectionFill ...
func ReflectionFill(o interface{}) error {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		typeField := t.Field(i)
		field := v.Field(i)
		tag := typeField.Tag

		if field.IsValid() && field.CanSet() {
			switch typeField.Type.Kind() {
			case reflect.Bool:
				field.SetBool(Confirm(typeField.Name))
				break
			case reflect.String:
				optionsString := tag.Get("options")
				if optionsString != "" {
					options := strings.Split(optionsString, ",")
					index, err := PromptWithChoice(typeField.Name, options)
					if err != nil {
						return err
					}
					field.SetString(options[index])
				} else {
					field.SetString(Prompt(typeField.Name))
				}
				break
			}
		}
	}
	return nil
}
