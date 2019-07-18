package main

import (
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"fmt"

	
	
)

type DockerAuth struct {
	Username      string
	Password      string
	ServerAddress string
}


func PullPublicImage(img string) bool {

	fmt.Println("Pull")
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	CheckIfError(err)

	ctx := context.Background()
	/* rep, err := cli.RegistryLogin(ctx, d)
	CheckIfError(err) */

	// fmt.Println(rep)
	/*authmarshed, _:=json.Marshal(d)
	sEnc := base64.StdEncoding.EncodeToString((authmarshed))
	
	opts := types.ImagePullOptions{
		RegistryAuth: sEnc,
	}
	optpsuh := types.ImagePushOptions{
		RegistryAuth: sEnc,
	}
	*/

	pull, err := cli.ImagePull(ctx, img, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	}else {

	io.Copy(os.Stdout, pull)
		return true
	}
	
	/*
	push, err := cli.ImagePush(ctx, "580918313152.dkr.ecr.eu-west-1.amazonaws.com" + "/site-furniture:2.165.44-TJT-299-2", optpsuh)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, push)
	*/
}


func Tag ( sourcetag, destinationtag string) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
		
	CheckIfError(err)

	ctx := context.Background()
	cli.ImageTag(ctx, sourcetag, destinationtag)
}
