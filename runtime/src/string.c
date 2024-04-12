#include <string.h>

typedef struct String {
    int length;
    char* value;
} string;

string new_string(char* str) {
    string s;
    s.length = strlen(str);
    s.value = str;
    return s;
}
