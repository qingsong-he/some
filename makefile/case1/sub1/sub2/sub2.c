#include <stdio.h>
#include "sub2.h"
#include "sub1/sub1.h"

void sub2(void) {
	puts("sub2");
	sub1();
}
