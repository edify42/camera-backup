test:
	go test ./... -v
test-coverage:
	go test ./... -coverprofile cp.out