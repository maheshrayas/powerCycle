package functions

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/maheshrayas/powerCycle/common/computeEngine"
	"github.com/maheshrayas/powerCycle/common/gke"
	"github.com/maheshrayas/powerCycle/common/configuration"
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
	// a := &computeEngine.VMInstances{
	// 	Ctx: context.Background(),
	// 	Config: config,
	// }
	// a.InitVMClient()
	// Instances := a.GetInstances(projectID)
	// json.NewEncoder(w).Encode(Instances)


	a := &gke.Cluster{
		Ctx: context.Background(),
		Config: config,
	}
	a.InitContainerClient()
	// jw := writers.NewMessageWriter(Instances)
	// jsonString, err := Instances.JSONString()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	log.Println(err.Error())
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(jsonString))
}
