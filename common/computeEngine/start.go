package computeEngine

import (
	"log"
	"fmt"
)

func (v *VMInstances) StartVMInstances(project string, zone string, name string) {
	_, err := v.computeService.Instances.Start(project, zone, name).Context(v.Ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Started instance %s", name)
}
