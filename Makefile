lambda_env_vars = GOARCH=amd64 GOOS=linux

install:
	dep ensure

build:
	go build -i -ldflags="-s -w" -o bin/fetch-feeds ./fetch-feeds

clean:
	rm -rf bin/*

test:
	go test ./import-csv-to-mysql

coverage:
	./tools/coverage.sh
