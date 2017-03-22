package models

type Model struct {
    Current bool
    Name string
}

func NewModel(curr bool, name string) *Model {
    return &Model{Current: curr, Name: name}
}

func GetModels() []*Model {
    return []*Model{
        NewModel(true, "Iglesias"),
    }
}
