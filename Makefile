build-krisz:
	GOOS=darwin CGO_ENABLED=1 go build -buildmode=plugin

build-dlm:
	GOOS=darwin CGO_ENABLED=0 go build -o dlm