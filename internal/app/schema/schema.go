package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/schema/hello"
	"github.com/wisaitas/graphql-golang/internal/app/schema/user"
	"github.com/wisaitas/graphql-golang/internal/app/service"
)

// CreateSchema สร้าง GraphQL schema แบบใหม่ที่แยกตาม business domain
func CreateSchema() (graphql.Schema, error) {
	// สร้าง services
	userService := service.NewUserService()

	// รวม queries จากทุก domain
	allQueries := graphql.Fields{}

	// เพิ่ม hello queries
	for name, field := range hello.GetQueries() {
		allQueries[name] = field
	}

	// เพิ่ม user queries
	for name, field := range user.GetQueries(userService) {
		allQueries[name] = field
	}

	// รวม mutations จากทุก domain
	allMutations := graphql.Fields{}

	// เพิ่ม user mutations
	for name, field := range user.GetMutations(userService) {
		allMutations[name] = field
	}

	// สร้าง Query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: allQueries,
	})

	// สร้าง Mutation type
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: allMutations,
	})

	// สร้าง schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return schema, err
}
