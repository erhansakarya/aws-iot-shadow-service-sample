package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// using the config value, create the iot client
	iotClient := iot.NewFromConfig(cfg)
	// using the config value, create the iotdataplane client
	idpClient := iotdataplane.NewFromConfig(cfg)

	// pull list of the things and get its shadows
	if thingList, err := iotClient.ListThings(context.TODO(), &iot.ListThingsInput{}); err == nil {
		for _, thing := range thingList.Things {
			fmt.Println("Thing Name: " + *thing.ThingName)
			if shadowPayload, err := idpClient.GetThingShadow(context.TODO(), &iotdataplane.GetThingShadowInput{ThingName: thing.ThingName}); err == nil {
				fmt.Println("Shadow Payload: " + string(shadowPayload.Payload))
			}
		}
	}
}
