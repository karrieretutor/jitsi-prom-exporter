# Jitsi Prometheus exporter
Exporter that grabs various metrics from [Jitsi](https://jitsi.org), especially form the video bridges, and publishes them as [Prometheus](https://prometheus.io) metrics.
The basic idea is to enter the jvbbrewery MUC room and listen to the presence broadcasts from the JVBs.

There is a [documentation](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md) of the published statistics by the video bridges.

# Run
There are multiple ways to run the exporter. Once it is running, it will publish the collected metrics on `:8080/metrics`.

## Register user in prosody
You will have to create a XMPP user in Prosody, do this with:
```bash
prosodyctl --config <abs path to jitsi-meet.cfg.lua> register <user> <auth-domain> <password>
```

Also, ensure, that **tls module is enabled in prosody**, without tls module enabled exporter woun't authorize to prosody.

## Configure JVB
Ensure, that **streaming statistics via colibri is enabled in jvb**, sip-communicator.properties should contain configuration like this:
```
org.jitsi.videobridge.ENABLE_STATISTICS=true
org.jitsi.videobridge.STATISTICS_TRANSPORT=muc,colibri
org.jitsi.videobridge.STATISTICS_INTERVAL=1000
```

## Binary
Clone this repo into your `$GOPATH/src/` directory. In the exporter directroy run `go get ./...` which creates the `exporter` binary in `$GOPTAH/bin/`. You can run this binary, it will still pull its configuration (see below) from the environment.

## Docker container
There is an image available on [docker hub](https://hub.docker.com/r/karrieretutor/jitsi) `karrieretutor/jitsi:prom-exporter-latest`. Alternatively build it yourself with the provided dockerfile. The configuration is provided via environment (see below). By running it as a docker container you gain the advantage of choosing the port which is used for publishing the metrics, see [docker cli reference](https://docs.docker.com/engine/reference/commandline/run/#publish-or-expose-port--p---expose).

# env configuration
env | description | default value
--- | --- | ---
`PROMEXP_AUTH_USER` | xmpp user for authentication |
`PROMEXP_AUTH_PASSWORD` | xmpp password for authentication |
`XMPP_SERVER` | xmpp server host name | 
`XMPP_PORT` | xmpp port to use | `5222`
`XMPP_AUTH_DOMAIN` | xmpp domain to authenticate against |
`XMPP_INTERNAL_MUC_DOMAIN` | internal muc domain (this is where the jvbbrewery muc resides) | 
`JVB_BREWERY_MUC` | name of jvbbrewery MUC room to join; it will join with `prom-exporter` as nickname -> `JVB_BREWERY_MUC@INTERNAL_MUC_DOMAIN/prom-exporter` | jvbbrewery
`JVB_METRIC_SUBSYSTEM` | Allows you to customize the metric names: `[<subsystem>_][<namespace>_]metricname`; both are optional |
`JVB_METRIC_NAMESPACE` | see `JVB_METRIC_SUBSYSTEM` | 

`PROMEXP_AUTH_USER` and `XMPP_AUTH_DOMAIN` are used to construct the JID `PROMEXP_AUTH_USER@XMPP_AUTH_DOMAIN`

# Grafana
If u want to use grafana dashboard in [grafana-dashboard.json](https://github.com/karrieretutor/jitsi-prom-exporter/blob/master/examples/grafana-dashboard.json), define in you grafana prometheus datasource called prometheus-jitsi or modify dashboard json to match your datasource name.
