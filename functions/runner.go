package functions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maheshrayas/powerCycle/common/computeEngine"
	"github.com/maheshrayas/powerCycle/common/configuration"
	"github.com/maheshrayas/powerCycle/common/gke"
)

//PowerCycle Entry point for the cloud functions
func PowerCycle(w http.ResponseWriter, r *http.Request) {
	var projectID string
	config := &configuration.Configs{}
	config.ReadConfig()
	if config.Projects != nil {
		for _, project := range config.Projects {
			projectID = project.ProjectID
		}
	}
	a := &computeEngine.VMInstances{
		Ctx:    context.Background(),
		Config: config,
	}
	a.InitVMClient()
	b := &gke.K8Clusters{
		Ctx:    context.Background(),
		Config: config,
	}
	b.InitContainerClient()
	computeEngChan := make(chan struct{})
	gkeChan := make(chan struct{})
	go func() {
		Instances := a.GetInstances(projectID)
		fmt.Println(Instances)
		close(computeEngChan)
	}()
	go func() {
		Nodes := b.GetClusters(projectID)
		fmt.Println(Nodes)
		close(gkeChan)
	}()
	<-computeEngChan
	<-gkeChan
}
