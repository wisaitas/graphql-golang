package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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

	json.NewEncoder(w).Encode(result)
}

// // GraphQLRequest โครงสร้างสำหรับ GraphQL request
// type GraphQLRequest struct {
// 	Query     string                 `json:"query"`
// 	Variables map[string]interface{} `json:"variables,omitempty"`
// }

// // Server โครงสร้างสำหรับ GraphQL server
// type Server struct {
// 	Schema graphql.Schema
// }

// // NewServer สร้าง server instance ใหม่
// func NewServer() (*Server, error) {
// 	schema, err := schema.CreateSchema()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create schema: %v", err)
// 	}

// 	return &Server{Schema: schema}, nil
// }

// // GraphQLHandler handler สำหรับ GraphQL endpoint
// func (s *Server) GraphQLHandler(w http.ResponseWriter, r *http.Request) {
// 	// ตั้งค่า CORS headers
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.Header().Set("Content-Type", "application/json")

// 	// จัดการ OPTIONS request สำหรับ CORS
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	var req GraphQLRequest

// 	// รองรับทั้ง GET และ POST
// 	if r.Method == "GET" {
// 		req.Query = r.URL.Query().Get("query")
// 	} else if r.Method == "POST" {
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 			return
// 		}
// 	} else {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// รัน GraphQL query
// 	result := graphql.Do(graphql.Params{
// 		Schema:         s.Schema,
// 		RequestString:  req.Query,
// 		VariableValues: req.Variables,
// 	})

// 	// ส่งผลลัพธ์กลับ
// 	json.NewEncoder(w).Encode(result)
// }

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
