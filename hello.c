#include <stdio.h>

int x = 0;

int hello() {
	printf("%d\n", ++x);
    return 0;
}
