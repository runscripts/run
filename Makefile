# Install
#   sudo GOPATH=$GOPATH make install
#
# Release
#   NOTICE: Backup your run.conf before run the following command
#   sudo GOPATH=$GOPATH make packages VERSION="X.Y.Z"
#
# Other
#   make test
#   sudo make clean
#   sudo make purge
#   sudo GOPATH=$GOPATH make deb VERSION="X.Y.Z"

.PHONY: help test install clean purge packages deb

ifeq (`uname`,'Darwin')
RUN_CONF=/usr/local/etc/run.conf
else
RUN_CONF=/etc/run.conf
endif

RUN_BIN=/usr/bin/run
DATA_DIR=/usr/local/run
MAN_PAGE=/usr/share/man/man1/run.1.gz

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

help:
	echo 'sudo GOPATH=$$GOPATH make install'

test:
	cd flock && go test -cover -v
	cd utils && go test -cover -v

install:
	$(_OS) $(_ARCH) $(BUILD) -o $(RUN_BIN) $(MAIN)
	[ -e $(RUN_CONF) ] || cp run.conf $(RUN_CONF)
	mkdir -p $(DATA_DIR) && chmod 777 $(DATA_DIR)
	cp VERSION $(DATA_DIR)
	gzip -c man/run.1 > $(MAN_PAGE)

clean:
	rm -f $(RUN_BIN)
	rm -rf $(DATA_DIR)
	rm -f $(MAN_PAGE)

purge: clean
	rm -f $(RUN_CONF)

packages: deb $(PACKAGES)

$(PACKAGES):
	echo $@ | awk -F_ '{print "mkdir -p packages/"$$1"_"$$2}' | bash
	echo $@ | awk -F_ '{print "CGO_ENABLED=0 GOOS="$$1" GOARCH="$$2" $(BUILD) -o packages/"$$1"_"$$2"/run $(MAIN)"}' | bash

deb:
	[ -n "$(VERSION)" ] && [ `whoami` = 'root' ]
	make purge
	mkdir -p $(DEB_DIR) && rm -rf $(DEB_DIR)/*.deb
	checkinstall $(DEB_FLAG) --arch="amd64" make install _OS="GOOS=linux" _ARCH="GOARCH=amd64"
	checkinstall $(DEB_FLAG) --arch="386" make install _OS="GOOS=linux" _ARCH="GOARCH=386"
	mv *.deb $(DEB_DIR)
