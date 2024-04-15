#include "runtime.h"

typedef struct Test {
    int a;
    int b;
} Test;
int main() {
    Test a = {
        .a = 1,
        .b = 2
    };
}
