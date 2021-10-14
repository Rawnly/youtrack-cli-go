bin := youtrack
buildFolder := ./release
buildPath := $(buildFolder)/$(bin)

build:
	rm -rf build
	go build -o $(buildPath)

install:
	go build -o $(buildPath)
	mv $(buildPath) /usr/local/bin/

tar:
	tar -czf $(bin).tar.gz --directory=$(buildFolder) $(bin)
	shasum -a 256 $(bin).tar.gz

prepare: build tar

tag:
	git tag -a $(version) -m "Version: $(version)"
	git push --tags


