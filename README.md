# mdmdirector-client-go

## Basic Usage

```
package main

import (
	"fmt"
	"log"
	"os"

	mdmdirector "github.com/gavinelder/mdmdirector-client-go"
)

func main() {
	host := os.Getenv("MDMDIRECTOR_URL")
	username := os.Getenv("MDMDIRECTOR_USER")
	password := os.Getenv("MDMDIRECTOR_PASS")
	c, err := mdmdirector.NewClient(&host, &username, &password)
	if err != nil {
		log.Fatal(err)
	}
	sp, err := c.GetSharedProfiles()
	if err != nil {
		log.Fatal(err)
	}
	for _, profile := range *sp {
		fmt.Println(profile.ID)
	}
}
}
```

## Implementations.

```sh
	r.HandleFunc("/webhook", director.WebhookHandler).Methods("POST")
	r.HandleFunc("/profile", utils.BasicAuth(director.PostProfileHandler)).Methods("POST")
	r.HandleFunc("/profile", utils.BasicAuth(director.DeleteProfileHandler)).Methods("DELETE")
	r.HandleFunc("/profile", utils.BasicAuth(director.GetSharedProfiles)).Methods("GET")
	r.HandleFunc("/profile/{udid}", utils.BasicAuth(director.GetDeviceProfiles)).Methods("GET")
	r.HandleFunc("/device", utils.BasicAuth(director.DeviceHandler)).Methods("GET")
	r.HandleFunc("/device/command/{command}", utils.BasicAuth(director.PostDeviceCommandHandler)).Methods("POST")
	r.HandleFunc("/device/serial/{serial}", utils.BasicAuth(director.SingleDeviceSerialHandler)).Methods("GET")
	r.HandleFunc("/device/push/{udid}", utils.BasicAuth(director.PushDeviceHandler)).Methods("GET")
	r.HandleFunc("/device/{udid}", utils.BasicAuth(director.SingleDeviceHandler)).Methods("GET")
	r.HandleFunc("/installapplication", utils.BasicAuth(director.PostInstallApplicationHandler)).Methods("POST")
	r.HandleFunc("/installapplication", utils.BasicAuth(director.GetSharedApplicationss)).Methods("GET")
	r.HandleFunc("/command/pending", utils.BasicAuth(director.GetPendingCommands)).Methods("GET")
	r.HandleFunc("/command/pending/delete", utils.BasicAuth(director.DeletePendingCommands)).Methods("GET")
	r.HandleFunc("/command/error", utils.BasicAuth(director.GetErrorCommands)).Methods("GET")
	r.HandleFunc("/command", utils.BasicAuth(director.GetAllCommands)).Methods("GET")
	r.HandleFunc("/health", director.HealthCheck).Methods("GET")
```
