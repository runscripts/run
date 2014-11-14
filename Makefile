.PHONY: hello test install clean purge

CONF_FILE = runscripts.yml
DATA_DIR  = /var/lib/runscripts
RUN_BIN   = /usr/bin/run

hello:
	go version

test:
	cd utils && go test

install:
	[[ -e /etc/$(CONF_FILE) ]] || cp $(CONF_FILE) /etc/
	mkdir -p $(DATA_DIR) && chmod 777 $(DATA_DIR)
	go build -o $(RUN_BIN) -v run.go

clean:
	rm -f $(RUN_BIN)
	rm -fr $(DATA_DIR)

purge: clean
	rm -f /etc/$(CONF_FILE)
