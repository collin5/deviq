BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

EXENAME=src/main.go

export GOPATH=$(CURDIR)
export GOBIN=$(GOPATH)/bin

makedir:
	@echo "starting build ..."
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

build:
	$(GOINSTALL) $(EXENAME)
	@echo "All done!"

clean:
	@echo "cleaning up ..."
	@rm -rf $(BUILDPATH)/bin
	@rm -rf $(BUILDPATH)/pkg
	@echo "All done!"

all: makedir build