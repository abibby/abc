```abc
type Test struct {
    a int;
    b int;
}

func main() int {
    var a Test = Test{
        a: 1,
        b: 2,
    };
}
```

```c
#include "runtime.h"

int main();

typedef struct Test {
    int a;
    int b;
} Test;

int main() {
    Test a = (Test){
        .a = 1,
        .b = 2
    };
}
```
