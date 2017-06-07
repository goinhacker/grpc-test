
all: proto

deps:
	go get -d -v google.golang.org/grpc/...

proto: deps
	mkdir generated

	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi

	go get -u -v github.com/golang/protobuf/protoc-gen-go

	for dir in $$(git ls-files '../../proto/*.proto' | xargs -n1 dirname | uniq); do \
		protoc -I $$dir --go_out=plugins=grpc:./generated $$dir/*.proto; \
	done

clean:
	rm -rf generated

.PHONY: all deps proto clean
