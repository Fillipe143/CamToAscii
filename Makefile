.SILENT:

build:
	go build -o cam_to_ascii main.go

run: build  
	./cam_to_ascii
