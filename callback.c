#include "callback.h"
#include "_cgo_export.h"

// C wrapper to call callback with real types
int stream_write_wrapper(void *user_data, const void *data, int length)
{
    return goStreamWriteCallback(user_data, (void *)data, length);
}