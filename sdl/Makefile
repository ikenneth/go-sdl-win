include $(GOROOT)/src/Make.inc

TARG=sdl

GOFILES:=constants.$(O).go sdl.go

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g sdl constants.c > constants.$(O).go
	gofmt -w constants.$(O).go
