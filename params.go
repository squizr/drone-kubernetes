package main

import "k8s.io/kubernetes/pkg/apis/extensions"

//Params Input Parameters
type Params struct {
	Cluster    string `json:"cluster"`
	Token      string `json:"token"`
	Deployment struct {
		extensions.Deployment
	}
}
