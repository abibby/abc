#include "stdlib.h"

void* new_pointer(size_t size, void* value) {
    return memcpy(malloc(size), value, size);
}