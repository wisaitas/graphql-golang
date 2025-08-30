package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/wisaitas/graphql-golang/internal/app/middleware/config"
)

func NewApp() {
	service := newService()
	resolver := newResolver(service)
	graphqlType := newGraphqlType()
	schema := newSchema(resolver, graphqlType)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		graphQLHandler(w, r, schema)
	})

	fmt.Println("GraphQL server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

func graphQLHandler(w http.ResponseWriter, r *http.Request, schema graphql.Schema) {
	config.NewCORS(r)

	var req graphQLRequest

	switch r.Method {
	case "GET":
		req.Query = r.URL.Query().Get("query")
	case "POST":
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
	})

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// // // GraphiQLHandler handler สำหรับ GraphiQL interface
// // func (s *Server) GraphiQLHandler(w http.ResponseWriter, r *http.Request) {
// // 	html := `
// // <!DOCTYPE html>
// // <html>
// // <head>
// //     <title>GraphiQL</title>
// //     <link href="https://unpkg.com/graphiql@3.0.6/graphiql.min.css" rel="stylesheet" />
// // </head>
// // <body style="margin: 0;">
// //     <div id="graphiql" style="height: 100vh;"></div>
// //     <script
// //         crossorigin
// //         src="https://unpkg.com/react@18/umd/react.production.min.js"
// //     ></script>
// //     <script
// //         crossorigin
// //         src="https://unpkg.com/react-dom@18/umd/react-dom.production.min.js"
// //     ></script>
// //     <script
// //         crossorigin
// //         src="https://unpkg.com/graphiql@3.0.6/graphiql.min.js"
// //     ></script>
// //     <script>
// //         const fetcher = GraphiQL.createFetcher({ url: '/graphql' });
// //         const root = ReactDOM.createRoot(document.getElementById('graphiql'));
// //         root.render(React.createElement(GraphiQL, { fetcher: fetcher }));
// //     </script>
// // </body>
// // </html>`
// // 	w.Header().Set("Content-Type", "text/html")
// // 	w.Write([]byte(html))
// // }

// // Start เริ่มต้น HTTP server
// func (s *Server) Start(port string) error {
// 	http.HandleFunc("/graphql", s.GraphQLHandler)
// 	// http.HandleFunc("/graphiql", s.GraphiQLHandler)
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	http.Redirect(w, r, "/graphiql", http.StatusFound)
// 	// })

// 	fmt.Printf("🚀 GraphQL server running at http://localhost%s\n", port)
// 	// fmt.Printf("📊 GraphiQL interface at http://localhost%s/graphiql\n", port)

// 	return http.ListenAndServe(port, nil)
// }
