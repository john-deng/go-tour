package reflection

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}


type Manager struct {
	User
	title string
}


func (u User) Hello(name string)  {
	fmt.Println("Hello, " + name)
}

func UserInfo(u interface{})  {
	t := reflect.TypeOf(u)
	fmt.Println("Type: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Wrong Type!")
		return
	}

	v := reflect.ValueOf(u)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func ManagerInfo(m interface{})  {
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}

func PrimitiveType()  {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(888)

	fmt.Printf("PremitiveType: %d\n", x)
}

func SetFieldByName(o interface{}, name, value string)  {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Error type")
		return
	}

	v = v.Elem()

	f := v.FieldByName(name)
	if !f.IsValid() {
		fmt.Printf("Invalid field name: %s\n", name)
		return
	}

	if f.Kind() == reflect.String {
		f.SetString(value)
	}
}

func ReflectMethod(o interface{}, fnName string, param string)  {
	v := reflect.ValueOf(o)
	mv := v.MethodByName(fnName)

	args := []reflect.Value{reflect.ValueOf(param)}

	mv.Call(args)
}

func Run()  {
	fmt.Println(">>> Test Reflection ...")
	u := User{1, "John", 25}
	UserInfo(&u)

	m := Manager{User: User{1, "John", 25}, title: "Directer"}
	ManagerInfo(m)

	PrimitiveType()

	SetFieldByName(&u, "Unknown", "Michael")
	fmt.Println("user: ", u.Name)
	SetFieldByName(&u, "Name", "Michael")
	fmt.Println("user: ", u.Name)

	ReflectMethod(&u, "Hello", "John")
}