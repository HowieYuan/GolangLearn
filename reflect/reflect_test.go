package reflect_test

//golang 中的反射

import (
	"fmt"
	"reflect"
	"testing"
)

// TypeOf ValueOf 分别可以获得类型名称和值
// Type 函数又可以将 Value 类型转换成 Type 类型
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

//检查参数类型
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)

}

//检查参数类型
func CheckType(v interface{}) {
	//TypeOf 函数获得参数类型
	t := reflect.TypeOf(v)
	// Kind 函数获得一个 Kind 类型，Kind 类型实质是一个 uint 和 reflect 包中类型常量对应
	switch t.Kind() {
	//以下类型都是 uint 常量
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

//=============================================================

//struct tag
//类似 Java 中的注解
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"` //format相当于key，normal相当于value
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}

	//按名字获取成员变量
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		//获得 format 的值
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	//按名字获取方法，并利用反射调用该方法
	//Call 方法接收方法参数
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
