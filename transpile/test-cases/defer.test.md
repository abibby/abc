```abc
func main() int {
    defer {
        println(str);
    };

    var a int = 1;
}
```

```c
#include "runtime.h"

int main();

int main() {
    int a = 1;
    {
        println("deferred");
    }
}
```
