all: build service bin

build:
	go build -o $@ cmd/mc-server/main.go

service:
	cp service/mc-server.service /usr/lib/systemd/system/mc-server.service

bin:
	cp mc-server /opt/bin/mc-server

clean:
	rm /usr/lib/systemd/system/mc-server.service
	rm /opt/bin/mc-server

