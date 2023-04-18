package main

import (
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/decode"
	"reflect"
	"strings"
)

type filterArgs struct {
	Checkpoint   graphql.NullString
	IsOffline    graphql.NullBool
	TransferType graphql.NullString
	IsVerified   graphql.NullBool
}

type dealsArgs struct {
	Query  graphql.NullString
	Filter *filterArgs
	Cursor *graphql.ID
	Offset graphql.NullInt
	Limit  graphql.NullInt
}

func main() {
	str := []byte("{\"query\": \"123123\",\"filter\": null,\"cursor\": null,\"offset\": 0,\"limit\": 100}")
	var m map[string]interface{}
	err := json.Unmarshal(str, &m)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%+v\n", m)

	args := dealsArgs{
		Query:  graphql.NullString{},
		Filter: nil,
		Cursor: nil,
		Offset: graphql.NullInt{},
		Limit:  graphql.NullInt{},
	}

	t := reflect.TypeOf(args)
	v := reflect.ValueOf(&args)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		lower := strings.ToLower(field.Name)
		if _, ok := m[lower]; !ok || m[lower] == nil {
			fmt.Printf("%s %v\n", lower, m[lower])
			continue
		}

		vn := reflect.New(t.Field(i).Type)
		if err := vn.Interface().(decode.Unmarshaler).UnmarshalGraphQL(m[lower]); err != nil {
			fmt.Printf("%v", err)
		} else {
			v.Elem().Field(i).Set(vn.Elem())
			fmt.Printf("%s %v\n", lower, m[lower])
		}

	}
	fmt.Printf("%+v\n", args)

}
