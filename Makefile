# Install
#   make && sudo GOPATH=$GOPATH make install
#
# Packages
#   NOTICE: Backup your run.yml before run the following command
#   sudo GOPATH=$GOPATH make packages VERSION="X.Y.Z"
#
# Other
#   make test
#   sudo make clean
#   sudo make purge
#   sudo GOPATH=$GOPATH make reinstall

.PHONY: deps test install clean purge reinstall packages deb

ifeq (`uname`,'Darwin')
RUN_CONF=/usr/local/etc/run.yml
RUN_BIN=/usr/local/bin/run
else
RUN_CONF=/etc/run.yml
RUN_BIN=/usr/bin/run
endif

DATA_DIR=/usr/local/run
DEB_DIR=packages/deb

BUILD=go build
MAIN=run.go

PACKAGES=linux_amd64 linux_386 linux_arm \
		 darwin_amd64 darwin_386 \
		 freebsd_amd64 freebsd_386

DEB_FLAG=-y --install=no --fstrans=yes --nodoc --backup=no \
		 --maintainer="wizawu@gmail.com" --pkglicense=MIT \
		 --pkgversion=`date +"%Y%m%d"` --pkgrelease=$(VERSION) \
		 --pkgaltsource="https://github.com/runscripts/run"

deps:
	go get github.com/kylelemons/go-gypsy/yaml

test: deps
	cd utils && go test

install:
	$(GOOS) $(GOARCH) $(BUILD) -o $(RUN_BIN) $(MAIN)
	[ -e $(RUN_CONF) ] || cp run.yml $(RUN_CONF)
	mkdir -p $(DATA_DIR) && chmod 777 $(DATA_DIR)
	cp LICENSE $(DATA_DIR)

clean:
	rm -f $(RUN_BIN)
	rm -rf $(DATA_DIR)

purge: clean
	rm -f $(RUN_CONF)

reinstall: purge install

packages: deb $(PACKAGES)

$(PACKAGES):
	echo $@ | awk -F_ '{print "mkdir -p packages/"$$1"_"$$2}' | bash
	echo $@ | awk -F_ '{print "CGO_ENABLED=0 GOOS="$$1" GOARCH="$$2" $(BUILD) -o packages/"$$1"_"$$2"/run $(MAIN)"}' | bash

deb: deps
	[ -n "$(VERSION)" ] && [ `whoami` = 'root' ]
	make purge
	mkdir -p $(DEB_DIR) && rm -rf $(DEB_DIR)/*.deb
	checkinstall $(DEB_FLAG) --arch="amd64" make install GOOS="GOOS=linux" GOARCH="GOARCH=amd64"
	checkinstall $(DEB_FLAG) --arch="386" make install GOOS="GOOS=linux" GOARCH="GOARCH=386"
	mv *.deb $(DEB_DIR)

