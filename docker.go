package main

import (
	"io"
	"os"

	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type DockerAuth struct {
	Username      string
	Password      string
	ServerAddress string
}

func NewDockerClient() *client.Client {

	dockerApiVersion := os.Getenv("DockerApiVersion")

	if dockerApiVersion == "" {
		dockerApiVersion = "1.38"
	}
	c, err := client.NewClientWithOpts(client.WithVersion(dockerApiVersion))
	CheckIfError(err)
	return c // enforce the default value here
}

func PullPublicImage(img string) bool {

	dockerClient := NewDockerClient()

	ctx := context.Background()

	pull, err := dockerClient.ImagePull(ctx, img, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		io.Copy(os.Stdout, pull)
		return true
	}

}

func Tag(sourcetag, destinationtag string) {

	dockerClient := NewDockerClient()

	ctx := context.Background()
	dockerClient.ImageTag(ctx, sourcetag, destinationtag)
}
