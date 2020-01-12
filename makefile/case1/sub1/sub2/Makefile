CFLAGS += -I ../../

objs = $(patsubst %.c, %.o, $(wildcard *.c))
headers = $(wildcard *.h)

libsub2.a: $(objs)
	ar crv $@ $<

$(objs): $(headers)

.PHONY: clean
clean:
	-rm *.o
	-rm *.a