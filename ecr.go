package main

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"io"
	"os"
	"strings"

	"encoding/json"
	"github.com/docker/docker/client"
)

func ECRRepoExists(reponame string) (exists bool, repouri string) {

	var r []*string
	r = append(r, aws.String(reponame))
	svc := ecr.New(session.New(&aws.Config{
		Region: aws.String("eu-west-1")}))

	input := &ecr.DescribeRepositoriesInput{}
	input.SetRepositoryNames(r)

	result, err := svc.DescribeRepositories(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println("RepoNotFound")
				return false, ""
			default:
				fmt.Println(aerr.Error())
			}
		}
	} else {
		return true, *result.Repositories[0].RepositoryUri
	}
	return false, ""

}

func ECRCreateRepo(reponame string) (created bool, repouri string) {
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(reponame),
	}

	svc := ecr.New(session.New(&aws.Config{
		Region: aws.String("eu-west-1")}))

	result, err := svc.CreateRepository(input)

	if err != nil {
		fmt.Println(err)
		return false, ""
	} else {
		println(*result.Repository.RepositoryUri)
		return true, *result.Repository.RepositoryUri
	}

}

// Ecrauth retrives the token
func Ecrauth() (c types.AuthConfig) {

	svc := ecr.New(session.New(&aws.Config{
		Region: aws.String("eu-west-1")}))

	params := &ecr.GetAuthorizationTokenInput{}

	// request the token
	resp, err := svc.GetAuthorizationToken(params)
	CheckIfError(err)

	// fields to send to template
	auth := resp.AuthorizationData[0]

	// extract base64 token
	data, err := base64.StdEncoding.DecodeString(*auth.AuthorizationToken)
	CheckIfError(err)
	// extract username and password
	token := strings.SplitN(string(data), ":", 2)

	b := types.AuthConfig{}
	// construct the login payload
	b.Username = token[0]
	b.Password = token[1]
	b.ServerAddress = *(auth.ProxyEndpoint)

	return b
}

func PushECR(d types.AuthConfig, img string) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))

	CheckIfError(err)

	ctx := context.Background()

	rep, err := cli.RegistryLogin(ctx, d)
	CheckIfError(err)

	fmt.Println(rep)
	authmarshed, _ := json.Marshal(d)
	sEnc := base64.StdEncoding.EncodeToString((authmarshed))

	optpsuh := types.ImagePushOptions{
		RegistryAuth: sEnc,
	}

	push, err := cli.ImagePush(ctx, img, optpsuh)
	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, push)
	}

}
