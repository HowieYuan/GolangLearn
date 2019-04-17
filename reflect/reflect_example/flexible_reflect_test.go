package reflect_example_test

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

// 万能编程
// 利用反射机制，将成员变量赋值到不同的结构体类型中
func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	//同一个成员变量参数可以赋值给 Employee 和 Customer
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	// 必须是指针
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem() 获取指针指向的值
		// 如果非结构体类型，则报错
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type.")
		}
	}

	// 成员变量为空，报错
	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	//循环遍历赋值
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		//赋值
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return nil
}
