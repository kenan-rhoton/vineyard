package models

import (
    "gopkg.in/mgo.v2/bson"
    "errors"
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "net/url"
)

type Model interface {
    class() string
    getKey() string
    getValue() interface{}
    validate() error
}

func ListAll() []string {
    return []string{
        "Iglesias",
    }
}

func Insert(m Model) error {
    dupe := Grab(m, m.getValue())
    if dupe == nil {
        return fmt.Errorf("Cannot add: key ", m.getValue(), " exists")
    }
    err := m.validate()
    if err != nil {
        return err
    }
    return TheDB.InsertInto(m.class(), m)
}

func Create(m Model, form url.Values) error {
    err := LoadModel(m, form)
    if err != nil {
        return err
    }
    return Insert(m)
}

func Update(m Model, form url.Values, key_value interface{}) error {
    err := LoadModel(m, form)
    if err != nil {
        return err
    }
    if reflect.ValueOf(key_value) != reflect.ValueOf(m.getValue()) {
        dupe := Grab(m, m.getValue())
        if dupe == nil {
            return fmt.Errorf("Cannot update: key ", m.getValue(), " exists")
        } else {
            LoadModel(m, form)
        }
    }
    err = m.validate()
    if err != nil {
        return err
    }
    return TheDB.Update(m.class(), bson.M{m.getKey(): key_value}, m)
}

func Grab(m Model, key_value interface{}) error {
    return TheDB.GetFrom(m.class(), bson.M{m.getKey(): key_value}, m)
}

func GrabSpecific(m Model, query bson.M) error {
    return TheDB.GetFrom(m.class(), query, m)
}

func GrabAll(m Model, results interface{}) error {
    return TheDB.GetAll(m.class(), results)
}

func Delete(m Model, key_value interface{}) error {
    return TheDB.DeleteFrom(m.class(), bson.M{m.getKey(): key_value})
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
