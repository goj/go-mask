include $(GOROOT)/src/Make.inc

TARG=bitmask
GOFILES=\
		bitmask.go\
		bitmask_template.go\

include $(GOROOT)/src/Make.pkg
