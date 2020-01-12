CFLAGS += -O2
CFLAGS += -I ../

objs = $(patsubst %.c, %.o, $(wildcard *.c))
libs = ../sub1/libsub1.a ../sub1/sub2/libsub2.a
libsDir := $(dir $(libs))

main1: $(objs) $(libs)

$(libs)::
	cd $(dir $@) && make

.PHONY: clean
clean: $(libsDir)
	-rm *.o
	-rm main1

$(libsDir)::
	-cd $@ && make clean