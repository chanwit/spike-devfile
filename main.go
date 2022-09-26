package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
)

type GitOpsRunExtension struct {
	PortForwards []PortForward `json:"port-forwards,omitempty"`
}

type PortForward struct {
	Namespace string `json:"namespace,omitempty"`
	Resource  string `json:"resource,omitempty"`
	Port      string `json:"port,omitempty"`
}

func main() {
	devfile, err := parser.ParseDevfile(parser.ParserArgs{
		Path: "./devfile.yaml",
	})
	if err != nil {
		panic(err)
	}

	components, err := devfile.Data.GetComponents(common.DevfileOptions{})
	if err != nil {
		panic(err)
	}
	for _, component := range components {
		if component.Kubernetes != nil {
			if len(component.Kubernetes.Endpoints) > 0 {
				var portForwards []PortForward
				for _, endpoint := range component.Kubernetes.Endpoints {
					portForwards = append(portForwards, PortForward{
						Port: strconv.Itoa(int(endpoint.TargetPort)),
					})
				}
				fmt.Println(portForwards)
			}
		}
	}

	commands, err := devfile.Data.GetCommands(common.DevfileOptions{
		CommandOptions: common.CommandOptions{
			CommandType: v1alpha2.ApplyCommandType,
		},
	})
	if err != nil {
		return
	}

	for _, c := range commands {
		attr := c.Attributes
		if v, ok := attr["gitops-run"]; ok {
			e := GitOpsRunExtension{}
			if err := json.Unmarshal(v.Raw, &e); err != nil {
				return
			}
		}
	}

}
