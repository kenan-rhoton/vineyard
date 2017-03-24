package models

import (
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "net/url"
)

type Model struct {
    Name string
}

func GetModels() []*Model {
    return []*Model{
        &Model{"Iglesias"},
    }
}

func SetField(obj interface{}, name string, value string) error {
    name = strings.Title(name)
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    switch structFieldType {
    case reflect.TypeOf(1):
        intval, _ := strconv.Atoi(value)
        val = reflect.ValueOf(intval)
    case reflect.TypeOf(true):
        boolval, _ := strconv.ParseBool(value)
        val = reflect.ValueOf(boolval)
    }

    if structFieldType != val.Type() {
        return errors.New("Provided value type didn't match obj field type")
    }

    structFieldValue.Set(val)
    return nil
}

func LoadModel(target interface{}, source url.Values) error {
    if reflect.ValueOf(target).Elem().Kind() != reflect.Struct {
        return errors.New("Not a Struct!")
    }
    for k, v := range source {
        err := SetField(target, k, strings.Join(v,""))
        if err != nil {
            return err
        }
    }
    return nil
}
