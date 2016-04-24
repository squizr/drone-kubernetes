package main

import (
	"fmt"
	"log"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone Kubernetes Plugin built from %s\n", buildCommit)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Cluster) == 0 {
		fmt.Println("Please provide a Cluster")
		os.Exit(1)
		return
	}

	if len(vargs.Token) == 0 {
		fmt.Println("Please provide a Token")
		os.Exit(1)
		return
	}

	if len(vargs.Deployment.ObjectMeta.Name) == 0 {
		fmt.Println("Please provide a Valid Deployment")
		os.Exit(1)
		return
	}

	config := &restclient.Config{
		Insecure:    true,
		Host:        vargs.Cluster,
		BearerToken: vargs.Token,
	}

	client, err := client.New(config)
	if err != nil {
		log.Fatalln("Can't connect to Kubernetes API:", err)
		os.Exit(1)
	}
	log.Println("Deployment Name: ", vargs.Deployment.ObjectMeta.Name)
	log.Println("Checking if Deployed already...")
	deploytment, err := client.Deployments(api.NamespaceDefault).Get(vargs.Deployment.ObjectMeta.Name)
	if err != nil {
		log.Println("Can't get deployment, Creating:", err)
		ndeploytment, err := client.Deployments(api.NamespaceDefault).Create(&vargs.Deployment.Deployment)
		if err != nil {
			log.Fatalln("Can't create deployment: ", err)
		} else {
			log.Println("Deployment Created: ", ndeploytment.ObjectMeta.Name)
		}
	} else {
		log.Println("Deployment Found: ", deploytment.ObjectMeta.Name)
		log.Println("Deployment Updating")
		ndeploytment, err := client.Deployments(api.NamespaceDefault).Update(&vargs.Deployment.Deployment)
		if err != nil {
			log.Fatalln("Can't update deployment: ", err)
		} else {
			log.Println("Deployment Updated: ", ndeploytment.ObjectMeta.Name)
		}
	}
}
