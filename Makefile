.PHONY: full build clean

full: clean build

build:
	@mkdir -p build
	CGO_ENABLED=1 go build -trimpath -a -ldflags '-w -s' -o ./build/api-server ./

clean:
	rm -rf build/
