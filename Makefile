.PHONY: lint test vendor clean

export GO111MODULE=on

default: prepare lint test

prepare:
	tar -xvzf geolite2.tgz

lint:
	golangci-lint run

test:
	go test -v -cover ./...

yaegi_test:
	yaegi test -v .

vendor:
	go mod vendor

clean:
	# rm -rf ./vendor
	rm -rf *.mmdb
