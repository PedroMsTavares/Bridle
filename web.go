package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	


)

type ImageReviewSpec struct {
	Containers  []map[string]string
	Annotations map[string]string
	Namespace   string
}

type ImageReview struct {
	ApiVersion string
	Kind       string
	Spec       ImageReviewSpec
}



func Validate(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var payload ImageReview
	err = json.Unmarshal(b, &payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	

	//Get Image from Docker 
	// Download the image 
	// Create Repo in ECR 
	// Push image



	/*e:=Ecrauth()
	PullImage(e) */

	
	// lets get all the images and send them to the Go magic
	for _, img := range payload.Spec.Containers {

		// cannot forget to change the image here to have the full canonical name example "docker.io/serverlessp/asdasdasd/asdasd:20"
		//go PullImage(img["image"])
		//fmt.Println(img)
		go HandlerToEcr(img["image"])
		
		
	}

	/*output, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output) */

}

func HandlerToEcr(img string){

	pull:=PullPublicImage(img)
	if pull {
		
		// Spliting the image in repo and tag
		i:=  strings.Split(img, ":")

		// Validate if the repo in ECR is Created
		exists, repouri:= ECRRepoExists(i[0])
		if exists{
			authtoken:=Ecrauth()
			Tag(img, repouri +":"+ i[1])
			PushECR(authtoken, repouri +":"+ i[1])
		}else{
			created, repouri:=ECRCreateRepo(i[0])
			if created {
				authtoken:=Ecrauth()
				Tag(img, repouri +":"+ i[1])
				PushECR(authtoken, repouri +":"+ i[1])
			}
		}
		
	}
	
}
