#include <stdlib.h>
#include <lcthw/dbg.h> // ref https://github.com/zedshaw/liblcthw

int main(int argc, char* argv[]) {
    // output to stderr
    debug("foobar");
    log_err("foobar");
    log_warn("foobar");
    log_info("foobar");

    // requre 'error' tag
	check(argc != 2, "argc != 2");
	check_mem(1);
	check_debug(0, "foobar")
	sentinel("argc != 2");

noerr:
	debug("noerr return of main");
	return EXIT_SUCCESS;
error:
	debug("error return of main");
	return EXIT_FAILURE;
}
