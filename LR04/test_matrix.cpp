#include <cstdio>

#include "matrix.hpp"

size_t test_mull_empty_one_thread() {
    printf("NOT IMPLEMENTED\n");
    return 1;
}

size_t test_mull_incorrect_size_one_thread() {
    printf("NOT IMPLEMENTED\n");
    return 1;
}

size_t test_all() {
    size_t failed = 0;

    printf("TESTING MATRIX MULL\n");
    failed += test_mull_empty_one_thread();
    failed += test_mull_incorrect_size_one_thread();

    return failed;
}

int main () {
    size_t failed = test_all();
    printf("%u tests failed!\n", failed);

    if (failed) {
        printf("FAIL!\n");
    } else {
        printf("SUCCESS!\n");
    }
    
    return 0;
}