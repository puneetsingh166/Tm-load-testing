# tm-load-test custom client for cosmos chain.
In order to use the tools, you will need:

* Go 1.13+
* 
* ## Building
To build the `tm-load-test custom client` binary go in the `customclient/my-cosmos-tester` directory:
Run this command ``` go build ```

## Usage
`tm-load-test custom client` can be executed in one of two modes: **standalone**, or
**master/slave**.

### Standalone Mode
In standalone mode, `tm-load-test custom client` operates in a similar way to `tm-bench`:

```bash
my-cosmos-tester -c 1 -T 10 -r 1000 -s 250 \
    --broadcast-tx-method async \
    --endpoints ws://tm-endpoint1.somewhere.com:26657/websocket,ws://tm-endpoint2.somewhere.com:26657/websocket
```
To see a description of what all of the parameters mean, simply run:

```bash
my-cosmos-tester --help
```

### Master/Slave Mode
In master/slave mode, which is best used for large-scale, distributed load 
testing, `tm-load-test custom client` allows you to have multiple slave machines connect to
a single master to obtain their configuration and coordinate their operation.

The master acts as a simple WebSockets host, and the slaves are WebSockets
clients.

On the master machine:

```bash
# Run tm-load-test custom client with similar parameters to the standalone mode, but now 
# specifying the number of slaves to expect (--expect-slaves) and the host:port
# to which to bind (--bind) and listen for incoming slave requests.
my-cosmos-tester \
    master \
    --expect-slaves 2 \
    --bind localhost:26670 \
    -c 1 -T 10 -r 1000 -s 250 \
    --broadcast-tx-method async \
    --endpoints ws://tm-endpoint1.somewhere.com:26657/websocket,ws://tm-endpoint2.somewhere.com:26657/websocket
```

On each slave machine:

```bash
# Just tell the slave where to find the master - it will figure out the rest.
my-cosmos-tester slave --master ws://localhost:26680
```

For more help, see the command line parameters' descriptions:

```bash
my-cosmos-tester master --help
my-cosmos-tester slave --help
```

## Monitoring
As of v0.4.1, `tm-load-test custom client` exposes a number of metrics when in master/slave 
mode, but only from the master's web server at the `/metrics` endpoint. So if
you bind your master node to `localhost:26670`, you should be able to get these
metrics from:

```bash
curl http://localhost:26670/metrics
```

The following kinds of metrics are made available here:

* Total number of transactions recorded from the master's perspective (across
  all slaves)
* Total number of transactions sent by each slave
* The status of the master node, which is a gauge that indicates one of the 
  following codes:
  * 0 = Master starting
  * 1 = Master waiting for all peers to connect
  * 2 = Master waiting for all slaves to connect
  * 3 = Load test underway
  * 4 = Master and/or one or more slave(s) failed
  * 5 = All slaves completed load testing successfully
* The status of each slave node, which is also a gauge that indicates one of the
  following codes:
  * 0 = Slave connected
  * 1 = Slave accepted
  * 2 = Slave rejected
  * 3 = Load testing underway
  * 4 = Slave failed
  * 5 = Slave completed load testing successfully
* Standard Prometheus-provided metrics about the garbage collector in 
  `tm-load-test`
* The ID of the load test currently underway (defaults to 0), set by way of the
  `--load-test-id` flag on the master

## Aggregate Statistics
As of `tm-load-test custom client` one can now write simple aggregate statistics to
a CSV file once testing completes by specifying the `--stats-output` flag:

```bash
# In standalone mode
my-cosmos-tester -c 1 -T 10 -r 1000 -s 250 \
    --broadcast-tx-method async \
    --endpoints ws://tm-endpoint1.somewhere.com:26657/websocket,ws://tm-endpoint2.somewhere.com:26657/websocket \
    --stats-output /path/to/save/stats.csv

# From the master in master/slave mode
my-cosmos-tester \
    master \
    --expect-slaves 2 \
    --bind localhost:26670 \
    -c 1 -T 10 -r 1000 -s 250 \
    --broadcast-tx-method async \
    --endpoints ws://tm-endpoint1.somewhere.com:26657/websocket,ws://tm-endpoint2.somewhere.com:26657/websocket \
    --stats-output /path/to/save/stats.csv
```

The output CSV file has the following format at present:

```csv
Parameter,Value,Units
total_time,10.002,seconds
total_txs,9000,count
avg_tx_rate,899.818398,transactions per second
```


