.PHONY: deps test install clean purge packages

CONF_FILE = runscripts.yml
DATA_DIR  = /var/lib/runscripts
RUN_BIN   = /usr/bin/run

deps:
	go get github.com/kylelemons/go-gypsy/yaml

test: deps
	cd utils && go test

install: deps
	[ -e /etc/$(CONF_FILE) ] || cp $(CONF_FILE) /etc/
	mkdir -p $(DATA_DIR) && chmod 777 $(DATA_DIR)
	go build -o $(RUN_BIN) -v run.go

clean:
	rm -f $(RUN_BIN)
	rm -rf $(DATA_DIR)

purge: clean
	rm -f /etc/$(CONF_FILE)

packages:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o packages/linux_amd64/run run.go
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -o packages/linux_386/run run.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -v -o packages/linux_arm/run run.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o packages/darwin_amd64/run run.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -v -o packages/darwin_386/run run.go
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -v -o packages/freebsd_amd64/run run.go
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -v -o packages/freebsd_386/run run.go
