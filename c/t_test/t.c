#include <lcthw/tests/minunit.h> // ref https://github.com/zedshaw/liblcthw

char* test_foobar() {
	return NULL;
}

char* test_foobar1() {
	mu_assert(1!=1, "1!=1");
	return NULL;
}

char* test_foobar2() {
	mu_assert(2==2, "2==2");
	return NULL;
}

char* all_tests() {
	mu_suite_start();
	mu_run_test(test_foobar);
	mu_run_test(test_foobar1);
	mu_run_test(test_foobar2);
	return NULL;
}

RUN_TESTS(all_tests);
