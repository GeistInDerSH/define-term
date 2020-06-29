all:
	go build -o def -ldflags="-s -w" main.go
	upx -5 -q def

remake:
	rm -f def
	make all

clean:
	rm -f def

