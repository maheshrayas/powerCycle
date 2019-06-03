package computeEngine

import (
	"encoding/json"
	"log"
"fmt"
	"sync"
"time"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	confg "github.com/maheshrayas/powerCycle/common/configuration"
)

//InitVMClient   Initialize
func (v *VMInstances) InitVMClient() error {
	v.instanceDetails = map[string]*CeInstances{}
	c, err := google.DefaultClient(v.Ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	v.computeService, err = compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (v *VMInstances) getZones(project string, region string) []string {
	resp, err := v.computeService.Regions.Get(project, region).Context(v.Ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	jsonRegions, _ := json.Marshal(resp)
	var regions Region
	json.Unmarshal(jsonRegions, &regions)
	return confg.ParseRegion(&regions.Zones)
}

//GetInstances Lists all the running instances
func (v *VMInstances) GetInstances(project string) []CeInstances {
	fmt.Println(time.Now())
	var wg sync.WaitGroup
	Instances := make([]CeInstances, 0)
	for _, region := range v.Config.Defaults.Regions{
		fmt.Println("Checking for reqion",region)
		for _, zone := range v.getZones(project, region) {
			fmt.Println("Checking for zone",region)
			req := v.computeService.Instances.List(project, zone)
			if err := req.Pages(v.Ctx, func(page *compute.InstanceList) error {
				for _, instance := range page.Items {
					wg.Add(1)
					v.instanceDetails[instance.Name] = &CeInstances{Name: instance.Name,
						Labels: instance.Labels, Zone: zone, State: instance.Status}
					go v.valdiateTags(project, zone, instance.Name, &wg)
				}
				return nil
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Wait()
	return Instances
}
