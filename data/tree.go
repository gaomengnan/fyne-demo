package data

import (
	"sync"

	"fyne.io/fyne/v2/data/binding"
)

type ObjectTree interface {
	binding.DataTree
	Append(parent, id string, value *SerializationConnectionData) error
	Get() (map[string][]string, map[string]*SerializationConnectionData, error)
	GetValue(id string) (*SerializationConnectionData, error)
	Prepend(parent, id string, value *SerializationConnectionData) error
	Set(ids map[string][]string, values map[string]*SerializationConnectionData) error
	SetValue(id string, value *SerializationConnectionData) error
}

type objectTree struct {
	listeners      sync.Map // map[DataListener]bool
	lock           sync.RWMutex
	ids            map[string][]string
	items          map[string]binding.DataItem
	updateExternal bool
	val            *map[string]*SerializationConnectionData
}

// AddListener attaches a new change listener to this DataItem.
// Listeners are called each time the data inside this DataItem changes.
// Additionally the listener will be triggered upon successful connection to get the current value.
func (o *objectTree) AddListener(_ binding.DataListener) {
	panic("not implemented") // TODO: Implement
}

// RemoveListener will detach the specified change listener from the DataItem.
// Disconnected listener will no longer be triggered when changes occur.
func (o *objectTree) RemoveListener(_ binding.DataListener) {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) GetItem(id string) (binding.DataItem, error) {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) ChildIDs(_ string) []string {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) Append(parent string, id string, value *SerializationConnectionData) error {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) Get() (map[string][]string, map[string]*SerializationConnectionData, error) {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) GetValue(id string) (*SerializationConnectionData, error) {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) Prepend(parent string, id string, value *SerializationConnectionData) error {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) Set(ids map[string][]string, values map[string]*SerializationConnectionData) error {
	panic("not implemented") // TODO: Implement
}

func (o *objectTree) SetValue(id string, value *SerializationConnectionData) error {
	panic("not implemented") // TODO: Implement
}

func NewObjectTree() ObjectTree {
	t := &objectTree{val: &map[string]*SerializationConnectionData{}}
	t.ids = make(map[string][]string)
	t.items = make(map[string]binding.DataItem)
	return t
}
