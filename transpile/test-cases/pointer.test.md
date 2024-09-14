```abc
type Test struct {
    a string;
}

func main() int {
    var a *Test = &Test{
        a: "test",
    };

    var b *Test = newTest();
}

func newTest() *Test {
    return &Test{
        a: "new func",
    };
}
```

```c
#include "runtime.h"

int main();

Test* newTest();

typedef struct Test {
    string a;
} Test;

int main() {
    Test* a = new_pointer(sizeof(Test), &(Test){
        .a = new_string("test")
    });
    Test* b = newTest();
}

Test* newTest() {
    return new_pointer(sizeof(Test), &(Test){
        .a = new_string("new func")
    });
}
```
