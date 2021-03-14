module github.com/c3sr/docker

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

require (
	github.com/Microsoft/hcsshim v0.8.15 // indirect
	github.com/c3sr/config v1.0.1
	github.com/c3sr/ldcache v1.0.0
	github.com/c3sr/logger v1.0.1
	github.com/c3sr/model v1.0.0
	github.com/c3sr/nvidia-smi v1.0.0
	github.com/c3sr/utils v1.0.0
	github.com/c3sr/uuid v1.0.1
	github.com/c3sr/vipertags v1.0.0
	github.com/carlescere/scheduler v0.0.0-20170109141437-ee74d2f83d82
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/containerd/continuity v0.0.0-20210313171317-968621f0704d // indirect
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.5+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-plugins-helpers v0.0.0-20200102110956-c9a8a2d92ccc
	github.com/dustin/go-humanize v1.0.0
	github.com/flyaways/golang-lru v0.0.0-20190617091412-ec8b77fcae6c
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/imdario/mergo v0.3.12
	github.com/junegunn/go-shellwords v0.0.0-20170411071455-02e3cf038dce
	github.com/k0kubun/pp/v3 v3.0.7
	github.com/moby/sys/mount v0.2.0 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/runc v1.0.0-rc93
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.7.0
	github.com/unknwon/com v1.0.1
	go.uber.org/goleak v1.1.10
	google.golang.org/grpc/examples v0.0.0-20210312231957-21976fa3e38a // indirect
	gotest.tools/v3 v3.0.3 // indirect
)
