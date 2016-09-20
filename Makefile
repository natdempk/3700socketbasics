all:
	$(RM) client
	export GOPATH=${PWD}
	go build -o client main.go

clean:
	$(RM) client
