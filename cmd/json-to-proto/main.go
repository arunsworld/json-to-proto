package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fname := flag.String("input", "", "")
	flag.Parse()

	if err := doMain(*fname); err != nil {
		log.Fatal(err)
	}
}

const input = `{
	"primaryKey":  "primary_key",
	"version":  "version",
	"user":  "user",
	"message":  "message",
	"transaction":  {
		"properties":  {
			"bookingId":  "booking_id",
          	"shippingInstructionId":  "shipping_instruction_id"
		},
		"booker":  "booker",
        "carrier":  "carrier"
	},
	"eventCode":  "event_code",
	"eventLocation":  {
		"locationType":  "EVENT_LOCATION_TYPE_CONTAINER_EVENT_LOCATION",
		"locationId":  "location_id",
		"locationIdType":  "LOCATION_ID_TYPE_UNLOC",
		"locationName":  "location_name",
		"locationSystemLiteral":  "location_system_literal"
	}
  }`

type Message struct {
	Name   string
	Fields []Field
}

type Field struct {
	FieldType   string
	FieldLabel  string
	FieldIndex  int
	IsRepeating bool
}

func parseJsonToMessages(input []byte) ([]Message, error) {
	result := []Message{}
	return result, nil
}

func doMain(fname string) error {
	// contents, err := os.ReadFile(fname)
	// if err != nil {
	// 	return err
	// }

	contents := []byte(input)

	result, err := parseJsonToMessages(contents)
	if err != nil {
		return err
	}

	fmt.Printf("%#v", result)

	// i, err := jwalk.Parse([]byte(input))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// switch v := i.(type) {
	// case jwalk.ObjectWalker:
	// 	v.Walk(func(key string, value interface{}) error {
	// 		fmt.Println(key + ":")
	// 		switch v := value.(type) {
	// 		case jwalk.ObjectsWalker:
	// 			v.Walk(func(obj jwalk.ObjectWalker) error {
	// 				fmt.Println("\t-")
	// 				obj.Walk(func(key string, value interface{}) error {
	// 					if v, ok := value.(jwalk.Value); ok {
	// 						fmt.Println("\t", key+":", v.Interface())
	// 					}
	// 					return nil
	// 				})
	// 				return nil
	// 			})
	// 		case jwalk.Value:
	// 			fmt.Println("\t", v.Interface())
	// 		case jwalk.ObjectWalker:
	// 			v.Walk(func(key string, value interface{}) error {
	// 				if v, ok := value.(jwalk.Value); ok {
	// 					fmt.Println("\t", key+":", v.Interface())
	// 				}
	// 				return nil
	// 			})
	// 		}
	// 		return nil
	// 	})
	// }
	return nil
}
