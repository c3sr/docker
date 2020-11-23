package docker

import "github.com/c3sr/model"

func init() {
	testPushModel = model.Push{
		Push:      true,
		ImageName: "raiproject/zipkin-cpp",
		Registry:  "https://index.docker.io/v1/",
		Credentials: model.DockerHubCredentials{
			Username: "dakkak",
			Password: "XXXX",
		},
	}
}
