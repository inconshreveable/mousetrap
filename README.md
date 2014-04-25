# mousetrap

mousetrap is a tiny library that answers a single question.

On a Windows machine, was the process invoked via the command line or by
someone double clicking the executable?

### The interface

The library exposes a single interface:

    func InvokedFromCommandLine() (bool, error)

On Windows, when a command line program is "double-clicked"
from the explorer, the terminal closes immediately after program termination.
Because the default behavior of most programs is to print the help and exit when
invoked with no arguments, this leads to an extremely poor user experience for
users unfamiliar with command line tools. mousetrap provides a way to detect these
invocations so that you can provide a helpful error message.
