# Contextual Log Information

Contextual information in log messages is important to fulfill the information needs that are required by the engineer to debug.


Two components of contextual information:
1. The value of the log message.
2. The stack trace and location of the log message.

## The Value
I found this blog post is good at giving us an example of the contextual log message. No further explanation is needed from me: https://reflectoring.io/logging-context/

To see more example in Go, check `00_acontextual_log`

## The Trace
Ability to identify the exact line of code that produces the log will help the engineer debug faster.
Not only the exact line of code but also the function stack call that leads into the log to give the engineer a wider picture.

To see my preferred solution for this, see `01_function_call_stack` for the problem, and `02_final_version` for the solution.
