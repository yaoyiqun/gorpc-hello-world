BINARY="grpc-server"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@go build -o ${BINARY} #-tags=jsoniter

proto:
	#@if [ ! -d "services" ] ; then mkdir services ; fi
	@protoc -I=./proto --go_out=plugins=grpc:./proto ./proto/*.proto
	@protoc -I=./proto --grpc-gateway_out=logtostderr=true:./proto ./proto/*.proto

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@go vet $(VETPACKAGES)

docker:
    @docker build -t wuxiaoxiaoshen/example:latest .

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default fmt fmt-check install test vet docker clean proto
