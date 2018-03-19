package goclitools

import (
	"fmt"
	"reflect"
	"strings"
)

// ReflectionFill ...
func ReflectionFill(o interface{}) error {
	return ReflectionFillUsingObject(o, false)
}

// ReflectionFillWithObject ...
func ReflectionFillUsingObject(o interface{}, useExistingValues bool) error {
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
				defaultValue := true
				if useExistingValues {
					defaultValue = field.Bool()
				}
				field.SetBool(ConfirmWithDefault(typeField.Name, defaultValue))
				break
			case reflect.String:
				optionsString := tag.Get("options")
				var promptText string
				if useExistingValues {
					promptText = fmt.Sprintf("%s (%s)", typeField.Name, field.String())
				} else {
					promptText = typeField.Name
				}
				if optionsString != "" {
					options := strings.Split(optionsString, ",")
					index, err := PromptWithChoice(promptText, options)
					if err != nil {
						return err
					}
					field.SetString(options[index])
				} else {
					value := Prompt(promptText)
					if useExistingValues && value == "" {
						break
					}
					field.SetString(value)
				}
				break
			}
		}
	}
	return nil
}
