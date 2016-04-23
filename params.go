package main

import "k8s.io/kubernetes/pkg/apis/extensions"

//Params Input Parameters
type Params struct {
	Cluster    string `json:"cluster"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Deployment struct {
		extensions.Deployment
	}
}
