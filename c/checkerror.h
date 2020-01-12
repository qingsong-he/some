#ifndef CHECKERROR_H
#define CHECKERROR_H

#include <string.h>
#include <errno.h>
#include <stdlib.h>

#define checkerror(hasError) {if (hasError) {fprintf(stderr, "checkerror %s:%d %s\n", __FILE__, __LINE__, strerror(errno)); exit(EXIT_FAILURE);}}

#endif
