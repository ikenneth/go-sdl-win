include $(GOROOT)/src/Make.inc

TARG=ttf

GOFILES:=constants.$(O).go ttf.go

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g ttf constants.c > constants.$(O).go
	gofmt -w constants.$(O).go
