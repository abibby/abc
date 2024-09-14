```abc
func main() int {
    var a int = 1;
    var b string = "test";

    function(b);
}

func function(str string) void {
    println(str);
}

```

```c
#include "runtime.h"

int main();

void function(string str);

int main() {
    int a = 1;
    string b = new_string("test");
    function(b);
}

void function(string str) {
    println(str);
}
```
