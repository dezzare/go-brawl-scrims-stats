package util

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/entity"
)

func ConvertToPlayer(data []byte) entity.Player {
	var p entity.Player
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println(err)
	}
	return p
}

func ParseToPlayer(data interface{}) {

	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Slice:
		// Iterate over the elements of the slice
		fmt.Println("Received a slice of structs:")
		for i := 0; i < v.Len(); i++ {
			fmt.Printf("Element %d: %+v\n", i, v.Index(i).Interface())
		}
	case reflect.Struct:
		// Handle a single struct
		fmt.Println("Received a single struct:")
		fmt.Printf("Data: %+v\n", data)
	default:
		fmt.Println("Unsupported type")
	}
}
