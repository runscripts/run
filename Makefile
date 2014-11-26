.PHONY: deps test install clean purge reinstall packages

ifeq (`uname`,'Darwin')
RUN_CONF=/usr/local/etc/run.yml
RUN_BIN=/usr/local/bin/run
else
RUN_CONF=/etc/run.yml
RUN_BIN=/usr/bin/run
endif

DATA_DIR=/usr/local/run

BUILD=go build
MAIN=run.go

PACKAGES=linux_amd64 linux_386 linux_arm \
		 darwin_amd64 darwin_386 \
		 freebsd_amd64 freebsd_386

deps:
	go get github.com/kylelemons/go-gypsy/yaml

test: deps
	cd utils && go test

install: deps
	[ -e $(RUN_CONF) ] || cp run.yml $(RUN_CONF)
	mkdir -p $(DATA_DIR) && chmod 777 $(DATA_DIR)
	$(BUILD) -v -o $(RUN_BIN) $(MAIN)

clean:
	rm -f $(RUN_BIN)
	rm -rf $(DATA_DIR)

purge: clean
	rm -f $(RUN_CONF)

reinstall: purge install

packages: $(PACKAGES)

$(PACKAGES):
	echo $@ | awk -F_ '{print "mkdir -p packages/"$$1"_"$$2}' | bash
	echo $@ | awk -F_ '{print "CGO_ENABLED=0 GOOS="$$1" GOARCH="$$2" $(BUILD) -o packages/"$$1"_"$$2"/run $(MAIN)"}' | bash
