run:
	go run cmd/app/main.go

runv2:
	go run cmd/appv2/main.go

test:
	go test ./...

lint:
	@set -e; \
	if command -v golangci-lint >/dev/null 2>&1; then \
		echo '>> using local golangci-lint'; \
		golangci-lint run; \
	elif command -v docker >/dev/null 2>&1; then \
		echo '>> using docker golangci/golangci-lint:$(GCI_VERSION)'; \
		docker run --rm \
			-v "$(PWD)":/app -w /app \
			-v golangci-cache:/root/.cache \
			golangci/golangci-lint:$(GCI_VERSION) golangci-lint run; \
	else \
		echo '!! ไม่พบ golangci-lint และไม่มี Docker'; \
		echo '   ทางเลือก: run "make tools" เพื่อติดตั้ง golangci-lint ในเครื่อง หรือ ติดตั้ง Docker แล้วรันใหม่'; \
		exit 1; \
	fi

format:
	go fmt ./...

check:
	make lint
	make test
	make format