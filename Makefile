
INSTALL = /usr/local

build: clean
	@echo building
	@go build -o xstatusbar

clean:
	@echo cleaning xstatusbar
	@rm -f xstatusbar

install: build
	@echo installing to ${INSTALL}/bin
	@mkdir -p ${INSTALL}/bin
	@cp -f xstatusbar ${INSTALL}/bin
	@chmod 755 ${INSTALL}/bin/xstatusbar
	@chmod u+s ${INSTALL}/bin/xstatusbar

uninstall:
	@echo removing ${INSTALL}/bin/xstatusbar
	@rm -f ${INSTALL}/bin/xstatusbar

.PHONY: build clean install uninstall
