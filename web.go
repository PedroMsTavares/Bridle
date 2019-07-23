package main

import (
	//	"encoding/json"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type AdmissionResponse struct {
	UID     string
	Allowed bool
}

func Validate(w http.ResponseWriter, r *http.Request) {

	log.SetOutput(os.Stdout)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bodyString := string(b)
	containers := gjson.Get(bodyString, "request.object.spec.containers")
	uid := gjson.Get(bodyString, "request.uid")

	containersarray := containers.Array()

	// lets get all the images and send them to the Go magic
	for _, img := range containersarray {

		image := gjson.Get(img.String(), "image")

		if !(strings.Contains(img.String(), ".ecr.")) {
			go HandlerToEcr(image.String())
		}

	}

	admissionResponse := &AdmissionResponse{
		UID:     uid.String(),
		Allowed: true,
	}
	response, err := json.Marshal(admissionResponse)
	CheckIfError(err)
	w.Header().Set("content-type", "application/json")
	w.Write(response)

}

func HandlerToEcr(img string) {

	pull := PullPublicImage(img)
	if pull {

		// Spliting the image in repo and tag
		i := strings.Split(img, ":")

		// Validate if the repo in ECR is Created
		exists, repouri := ECRRepoExists(i[0])
		if exists {
			authtoken := Ecrauth()
			Tag(img, repouri+":"+i[1])
			PushECR(authtoken, repouri+":"+i[1])
		} else {
			created, repouri := ECRCreateRepo(i[0])
			if created {
				authtoken := Ecrauth()
				Tag(img, repouri+":"+i[1])
				PushECR(authtoken, repouri+":"+i[1])
			}
		}

	}

}
