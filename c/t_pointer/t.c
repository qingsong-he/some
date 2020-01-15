#include <lcthw/dbg.h>

// in 64bit os
void func1(char a[]) { 
	// char*
	debug("%ld\n", sizeof(a)); // 8 
} 

void func2(char a[16]) { 
	// char*
	debug("%ld\n", sizeof(a)); // 8 
} 

int main(void) { 

	char a11[] = "hello";
	debug("%ld %ld\n", sizeof(a11), sizeof(a11) / sizeof(char)); // 6 6

	char a[2]; 
	debug("%ld\n", sizeof(a) / sizeof(char)); // 2 

	char b[0]; 
	debug("%ld\n", sizeof(b)); // 0 
	debug("%ld\n", sizeof(b) / sizeof(char)); // 0 

	char* p1 = b;
	debug("%ld\n", sizeof(p1)); // 8 
	debug("%ld\n", sizeof(p1) / sizeof(char)); // 8 

	char* p2 = &b;
	debug("%ld\n", sizeof(p2)); // 8 
	debug("%ld\n", sizeof(p2) / sizeof(char)); // 8 

	char *c; 
	debug("%ld\n", sizeof(c)); // 8 
	debug("%ld\n", sizeof(char*)); // 8 

	func1(NULL); 
	func1(a); 
	func2(NULL); 
	func2(a); 
	return 0; 
} 
