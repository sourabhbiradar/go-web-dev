# context

context allows us to signal that an operation can be cancelled.

Deadline associated with context, specifying the maximum time for operation to complete.

When root/parent context is cancelled all its child/children context are also cancelled.

Uses :

HTTP request context can carry values through middelware and handlers.

Cancelling database operations.

Cancelling long running Go Routines.



