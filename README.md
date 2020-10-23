# Delve Unit Test Demo

A very small failing unit test to demonstrate how to use dlv at the command line for debugging tests.

From the command line, run:

    $ dlv test

Then in a dlv prompt, run this...

    (dlv) break TestFibonacci

To list the tests available to break on, run this ...

    (dlv) funcs TestFi*

To break on the inner test, run this...

    (dlv) break TestFibonacci.func1

You can always break on a line number...

    (dlv) break ./fib_test.go:24
