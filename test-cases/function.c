#include "runtime.h"

void function(string str) {
    println(str);
}
int main() {
    int a = 1;
    string b = new_string("test");
    function(b);
}
