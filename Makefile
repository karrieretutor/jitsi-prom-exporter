EXPORTER_DOCKER_TAG=v0.6.2

build:
	docker build -t registryv2.landr.com/landr-jitsi-prom-exporter:${EXPORTER_DOCKER_TAG} .

push:
	docker push registryv2.landr.com/landr-jitsi-prom-exporter:${EXPORTER_DOCKER_TAG}

.PHONE: build