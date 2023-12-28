build:
	cd cmd/shortener && rm server-binary || true && go build -o server-binary *.go
run-server:
	make build && ./cmd/shortener/server-binary
test:
	go clean -testcache && go test ./...
test-iter-1:
	make build \
	&& ./shortenertestbeta-darwin-arm64 -test.v -test.run=^TestIteration1$$ -binary-path=cmd/shortener/server-binary
