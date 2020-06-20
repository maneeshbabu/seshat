.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/seshat main.go
	cd frontend && npm run build && cd ..

clean:
	rm -rf ./bin ./vendor Gopkg.lock
	rm -rf frontend/dist/
deploy: clean build
	sls deploy --verbose
