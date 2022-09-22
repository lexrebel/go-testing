.PHONY: test-cover
test-cover:
	go test -coverpkg=./... -coverprofile coverage/cover.out -v ./...
	go tool cover -func=coverage/cover.out
	go tool cover -html coverage/cover.out -o coverage/coverage.html

test: 
	go test -v .

cover: 
	go test -cover .

cover-html:
	go test -coverprofile=cover.out && go tool cover -html=cover.out