all: a.out

a.out: lib1/libso.so

lib1/libso.so:
	go build -buildmode=c-shared -o lib1/libso.so lib1/main.go

a.out:
	cc -Ilib1 use.c -Llib1 -lso

clean:
	-rm a.out
	-rm lib1/libso.h
	-rm lib1/libso.so
