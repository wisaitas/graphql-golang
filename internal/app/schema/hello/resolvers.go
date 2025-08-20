package hello

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// HelloResolver resolver สำหรับ hello query
func HelloResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if !ok || name == "" {
		return "Hello, World! This is GraphQL in Go! 🚀", nil
	}
	return fmt.Sprintf("Hello, %s! 👋", name), nil
}
