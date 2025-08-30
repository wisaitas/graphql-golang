# install dependency

go get github.com/99designs/gqlgen

# bootstrap for create gqlgen.yml , graph/schema.graphqls , server.go

go run github.com/99designs/gqlgen init
go install github.com/99designs/gqlgen@latest

# graphql generate

go run github.com/99designs/gqlgen generate
gqlgen generate