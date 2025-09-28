#ifndef CALLBACK_H
#define CALLBACK_H

#include <stdlib.h>
#include <plutobook.h>

int stream_write_wrapper(void *user_data, const void *data, int length);
plutobook_resource_data_t *resource_fetch_wrapper(void *closure, const char *url);

#endif