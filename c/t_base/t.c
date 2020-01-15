#include <stdlib.h>
#include <lcthw/dbg.h> // ref https://github.com/zedshaw/liblcthw

int main(int argc, char* argv[]) {
	check(argc != 2, "argc != 2");

noerr:
	debug("noerr return of main");
	return EXIT_SUCCESS;
error:
	debug("error return of main");
	return EXIT_FAILURE;
}
