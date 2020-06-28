tester:
	go test ./... -v
test-coverage:
	go test ./... -coverprofile cp.out

test-nocache:
	go test ./... -v -count=1