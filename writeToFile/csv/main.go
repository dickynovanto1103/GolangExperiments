package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

type Data struct {
	a, b *int
}

func parse(data interface{}) []string {
	dataValue := reflect.ValueOf(data)
	numField := dataValue.NumField()
	dataArray := make([]string, numField)

	for i:=0;i<numField;i++ {
		switch dataValue.Field(i).Kind() {
		case reflect.Int:
			val := dataValue.Field(i).Int()
			dataArray[i] = fmt.Sprintf("%v", val)
		case reflect.Ptr:
			val := dataValue.Field(i).Pointer()
			fmt.Println("val:", val)
		default:
			fmt.Println("halo")
		}

	}
	return dataArray
}

func main() {
	wr := csv.NewWriter(os.Stdout)

	val := 2
	data := &Data{
		a: &val,
		b: &val,
	}
	fmt.Println("data:", data)
	dataString := parse(*data)
	fmt.Println("dataString: ", dataString)
	wr.Write(dataString)
	wr.Flush()
}
