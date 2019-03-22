build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app get_url_one_param.go

scratch: build
	docker build -f Dockerfile.scratch -t pgluffy/go-app .