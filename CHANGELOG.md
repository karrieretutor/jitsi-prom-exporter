# Changelog
## Unreleased

## [0.0.3] 2019-09-06
### Added
- the tcp connections for testing the connection to the xmpp server are now closed

## [0.0.2] 2019-09-05
### Added
- the prom exporter now watches the tcp conn to the xmpp server (this led to some confusion on k8s), if the connection to the xmpp server is lost, the exporter exits 

### Fixed
- Exit the prom-exporter if connection gets closed externally (eg. the XMPP server stops)

## [0.0.1] 2019-08-30
### Added
- initial release