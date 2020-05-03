
all: sysv systemd upstart clean

test:
	@go test -c -o './bin/service.test' ./examples/service/

test-local:
	@go test -c -o '../../bin/service.test' ../../examples/servcie/

clean:
	-rm ./bin/service.test

sysv: test
	@echo sysv
	@cp ./bin/service.test ./ci/linux_test/sysv/
	@-docker rmi -f service.test.sysv
	@docker build -q --tag="service.test.sysv" sysv
	@-docker run service.test.sysv
	@-docker rm $(shell docker ps -l -q)
	#@-docker rmi -f service.test.sysv
	@-rm ./ci/linux_test/sysv/service.test

systemd: test
	@echo systemd
	@cp ./bin/service.test ./ci/linux_test/systemd/
	@-docker rmi -f service.test.systemd
	@docker build -q --tag="service.test.systemd" systemd
	@-docker run --privileged -v /sys/fs/cgroup:/sys/fs/cgroup:ro service.test.systemd
	@-docker rm $(shell docker ps -l -q)
	#@-docker rmi -f service.test.systemd
	@-rm ./ci/linux_test/systemd/service.test

upstart: test
	@echo upstart
	@cp ./bin/service.test ./ci/linux_test/upstart/
	@-docker rmi -f service.test.upstart
	@docker build -q --tag="service.test.upstart" upstart
	@-docker run service.test.upstart
	@-docker rm $(shell docker ps -l -q)
	#@-docker rmi -f service.test.upstart
	@-rm ./ci/linux_test/upstart/service.test
