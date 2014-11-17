.PHONY: deps test install clean purge

CONF_FILE = runscripts.yml
DATA_DIR = /var/lib/runscripts
BUILD_TARGET = ./build/run`getconf LONG_BIT`
RUN_BIN = /usr/bin/run

deps:
	go get github.com/kylelemons/go-gypsy || echo -n
	go get github.com/kylelemons/go-gypsy/yaml

test: deps
	cd utils && go test

build: deps
	go build -o $(BUILD_TARGET) -v run.go

install: deps
	[ -e /etc/$(CONF_FILE) ] || cp $(CONF_FILE) /etc/
	mkdir -p $(DATA_DIR) && chmod 777 $(DATA_DIR)
	cp $(BUILD_TARGET) $(RUN_BIN)

clean:
	rm -f $(RUN_BIN)
	rm -rf $(DATA_DIR)

purge: clean
	rm -f /etc/$(CONF_FILE)
