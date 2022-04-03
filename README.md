# openwhisk-benchmark

## launch-hello-exp

It benchmarks the invocation of a hello-world action.

To use it, first build and create the action `hello-go` in `actions` directory.

Then build the `launch-hello-exp` into a binary file.

Next, modify and run `launch-multi.sh` to stress OpenWhisk.

Finally, run `stat.py` to aggregate the result.


## launch-counter

It benchmarks the invocation of an action from the server side.

To use it, first build and create the action `counter` in `actions` directory.

Then build and run `counter-server`.

Next, modify the `host` in `launch.sh` to the API server address of OpenWhisk.

Finally, modify and run `launch-multi.sh` to stress OpenWhisk.
