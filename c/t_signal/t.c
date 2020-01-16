#include <stdlib.h>
#include <lcthw/dbg.h> // ref https://github.com/zedshaw/liblcthw
#include <signal.h>
#include <unistd.h>

static void signal_handler(int s) {
	debug("signal alarm: %d", s);
}

int main(int argc, char* argv[]) {
	if (SIG_ERR == signal(SIGINT, signal_handler)) {
		sentinel("sigint_handler");
	}
	if (SIG_ERR == signal(SIGPROF, SIG_DFL)) {
		sentinel("sigint_handler");
	}
	if (SIG_ERR == signal(SIGTERM, SIG_IGN)) {
		sentinel("sigint_handler");
	}

	for (;;) {
		pause();
	}
noerr:
	debug("noerr return of main");
	return EXIT_SUCCESS;
error:
	debug("error return of main");
	return EXIT_FAILURE;
}
