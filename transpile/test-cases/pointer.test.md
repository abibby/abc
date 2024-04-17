```abc
type Test struct {
    string a;
}

func main() int {
    var a *Test = &Test{
        a: "test",
    };

    var b *Test = a;
}
```

```c
#include "runtime.h"

int main();

typedef struct Test {
    string a;
} Test;

int main() {
    Test* a = new_pointer(sizeof(Test), &(Test){
        .a = new_string("test")
    });
    Test* b = a;
}
```
