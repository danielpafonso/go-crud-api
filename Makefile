.PHONY: all build clean

all: clean build

build:
	@mkdir -p build
	CGO_ENABLED=1 go build -trimpath -a -ldflags '-w -s' -o ./build/api-server ./
	cp init-script.sql build/

clean:
	rm -rf build/
