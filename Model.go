package gouda

import (
	"reflect"
	"fmt"
)

type ModelInterface interface {
	TableName() string
}

type Model struct {
	tablename  string
	attributes map[string]reflect.Type
}

type ModelStore struct{}


func attributes(m interface{}) map[string]reflect.Type {
	var st *reflect.StructType
	if _, ok := reflect.Typeof(m).(*reflect.PtrType); ok {
		st = reflect.Typeof(m).(*reflect.PtrType).Elem().(*reflect.StructType)
	} else {
		st = reflect.Typeof(m).(*reflect.StructType)
	}

	fmt.Println(st.NumField())

	ret := make(map[string]reflect.Type)

	for i := 0; i < st.NumField(); i++ {
		p := st.Field(i)
		fmt.Println(p.Name)
		if !p.Anonymous {
			ret[p.Name] = p.Type
		}
	}

	return ret
}

func (m Model) TableName() string { return m.tablename }


func M(m ModelInterface) *Model {
	mod := new(Model)
	mod.tablename = m.TableName()
	mod.attributes = attributes(m)
	return mod
}


func (m *Model) Attributes() map[string]reflect.Type {
	return m.attributes
}

func (m *Model) AttributesNames() (ret  []string) {
	ret=make([]string,len(m.attributes))
	i:=0
	for k,_:=range m.attributes {
	ret[i]=k
	i++
	}
	return ret
}


