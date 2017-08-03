package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"bytes"

	"fmt"

	"io"

	"os"

	"github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	action = flag.String("action", "create", "choose an action to perform")
	name   = flag.String("name", "", "the name of the resource to interact with")
	urlFmt = "https://localhost:3101/apis/mobile.k8s.io/v1alpha1/namespaces/%s"
)

func main() {
	flag.Parse()

	switch *action {
	case "create":
		create(*name)
		return
	case "delete":
		del(*name)
		return
	case "list":
		list()
		return

	case "get":
		get(*name)
		return
	}

}

func del(name string) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf(urlFmt, "test/mobileapp/"+name), nil)
	if err != nil {
		log.Fatalf("failed to marshal create request %s ", err)
	}
	client := getClient()
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to do request %s ", err)
	}
	defer res.Body.Close()
	readAndPrint(res.Body)
}

func get(name string) {
	req, err := http.NewRequest("GET", fmt.Sprintf(urlFmt, "test/mobileapp/"+name), nil)
	if err != nil {
		log.Fatalf("failed to marshal create request %s ", err)
	}
	client := getClient()
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to do request %s ", err)
	}
	defer res.Body.Close()
	readAndPrint(res.Body)

}

func list() {
	req, err := http.NewRequest("GET", fmt.Sprintf(urlFmt, "test/mobileapp"), nil)
	if err != nil {
		log.Fatalf("failed to marshal create request %s ", err)
	}
	client := getClient()
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to do request %s ", err)
	}
	defer res.Body.Close()
	readAndPrint(res.Body)

}

func create(name string) {
	mobileApp := v1alpha1.MobileApp{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "mobile.k8s.io/v1alpha1",
			Kind:       "MobileApp",
		},
		Spec: v1alpha1.MobileAppSpec{
			ClientType: "android",
		},
	}
	data, err := json.Marshal(&mobileApp)
	if err != nil {
		log.Fatalf("failed to marshal mobile app %s ", err)
	}

	fmt.Println(string(data))

	req, err := http.NewRequest("POST", fmt.Sprintf(urlFmt, "test/mobileapp"), bytes.NewReader(data))
	if err != nil {
		log.Fatalf("failed to marshal create request %s ", err)
	}
	client := getClient()
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to do request %s ", err)
	}
	defer res.Body.Close()
	readAndPrint(res.Body)
}

func readAndPrint(body io.ReadCloser) {

	if _, err := io.Copy(os.Stdout, body); err != nil {
		log.Fatalf("failed to read response body %s ", err)
	}
}

func getClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
