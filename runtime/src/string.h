#ifndef ABC_RUNTIME_STRING
#define ABC_RUNTIME_STRING

typedef struct String {
    int length;
    char* value;
} string;

string new_string(char* str);

#endif
