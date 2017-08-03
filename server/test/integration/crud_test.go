package integration

import (
	"testing"

	"fmt"

	"github.com/feedhenry/mobile-control-panel/server/pkg/apis/mobile"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createMobileApp(name, client string) *mobile.MobileApp {
	return &mobile.MobileApp{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "mobile.k8s.io/v1alpha1",
			Kind:       "MobileApp",
		},
		Spec: mobile.MobileAppSpec{
			ClientType: client,
		},
	}
}

func TestCreate(t *testing.T) {
	stopCh, client, _, err := StartDefaultServer()
	if err != nil {
		t.Fatal(err)
	}
	defer close(stopCh)
	cases := []struct {
		Name      string
		MobileApp *mobile.MobileApp
		ExpectErr bool
	}{
		{
			Name:      "test create ok",
			MobileApp: createMobileApp("testapp", "android"),
		},
		{
			Name:      "test create fails invalid client type",
			MobileApp: createMobileApp("testapp", "windows"),
			ExpectErr: true,
		},
	}

	for _, tc := range cases {
		createdApp, err := client.Mobile().MobileApps("test").Create(tc.MobileApp)
		if !tc.ExpectErr && err != nil {
			t.Fatalf("did not expect error creating MobileApp %s", err)
		}
		if tc.ExpectErr && err == nil {
			t.Fatalf("expected an error but got none")
		}
		fmt.Println(createdApp)
	}

}
