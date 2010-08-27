# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

QUOTED_GOROOT=$(GOROOT)

include $(GOROOT)/src/Make.$(GOARCH)

TARG=ttf

GOFILES:=constants.$(O).go ttf.go

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g ttf constants.c > constants.$(O).go
	gofmt -w constants.$(O).go
