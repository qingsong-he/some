all: a.out

a.out: so/libso.so

so/libso.so:
	cc -fPIC -shared -o so/libso.so so/so.c

a.out:
	go build -o a.out

clean:
	-rm a.out
	-rm so/libso.so
