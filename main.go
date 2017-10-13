package main

import (
	"net/http"
	"bytes"
	"github.com/graphql-go/graphql"
	"./gql/queries"
	"./gql/mutations"
	"github.com/graphql-go/graphql/gqlerrors"
	"fmt"
	"encoding/json"
)

func main() {

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queries.Root,
		Mutation: mutations.Root,
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(request.Body)

		result, err := executeQuery(schema, buf.String())

		if err != nil { fmt.Fprint(writer, err) } else {
			json.NewEncoder(writer).Encode(result)
		}
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func executeQuery(schema graphql.Schema, requestStr string) (*graphql.Result, gqlerrors.FormattedErrors) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestStr,
	})

	if len(result.Errors) > 0 {
		return nil, result.Errors
	}
	return result, nil
}
