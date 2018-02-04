.PHONY: default
default: all ;

all:
			make clean
			make build

build:
			go build -o build/director main.go

clean:
			rm -rf build

run:
			build/director
