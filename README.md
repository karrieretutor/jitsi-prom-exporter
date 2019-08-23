# Jitsi Prometheus exporter
Exporter that grabs various metrics from [Jitsi](https://jitsi.org), especially form the video bridges, and publishes them as [Prometheu](https://prometheus.io) metrics.
The basic idea is to enter the jvbbrewery MUC room and listen to the presence broadcasts from the JVBs.

# env configuration
env | description | default value
--- | --- | ---
`XMPP_USER` | xmpp user for authentication |
`XMPP_PW` | xmpp password for authentication |
`XMPP_AUTH_DOMAIN` | xmpp domain to authenticate against |
`XMPP_SERVER` | xmpp server host name | 
`XMPP_PORT` | xmpp port to use | `5222`

# Register user in prosody
You will have to create a XMPP user in Prosody, do this with:
```bash
prosodyctl --config <abs path to jitsi-meet.cfg.lua> register <user> <auth-domain> <password>
```

