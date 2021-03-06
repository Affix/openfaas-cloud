package function

import (
	"encoding/json"
	"os"
	"testing"
)

func TestBuildURLWithoutPrettyURL_WithSlash(t *testing.T) {
	os.Setenv("gateway_public_url", "http://localhost:8080")
	os.Setenv("gateway_pretty_url", "")

	event := &eventInfo{
		owner:   "alexellis",
		service: "tester",
	}

	val := buildPublicStatusURL("success", event)
	want := "http://localhost:8080/function/alexellis-tester"

	if val != want {
		t.Errorf("building PublicURL: want %s, got %s", want, val)
		t.Fail()
	}
}

func TestBuildURLWithoutPrettyURL_WithOutSlash(t *testing.T) {
	os.Setenv("gateway_public_url", "http://localhost:8080")
	os.Setenv("gateway_pretty_url", "")

	event := &eventInfo{
		owner:   "alexellis",
		service: "tester",
	}

	val := buildPublicStatusURL("success", event)
	want := "http://localhost:8080/function/alexellis-tester"

	if val != want {
		t.Errorf("building PublicURL: want %s, got %s", want, val)
		t.Fail()
	}
}

func TestBuildURLWithPrettyURL(t *testing.T) {
	os.Setenv("gateway_public_url", "http://localhost:8080")
	os.Setenv("gateway_pretty_url", "https://user.openfaas-cloud.com/function")

	event := &eventInfo{
		owner:   "alexellis",
		service: "tester",
	}

	val := buildPublicStatusURL("success", event)
	want := "https://alexellis.openfaas-cloud.com/tester"

	if val != want {
		t.Errorf("building PublicURL: want %s, got %s", want, val)
		t.Fail()
	}
}

func TestBuildURLWithUndefinedStatusGivesOriginalURL(t *testing.T) {

	event := &eventInfo{
		owner:   "alexellis",
		service: "tester",
		url:     "http://original-value.local",
	}

	val := buildPublicStatusURL("not-supported", event)
	want := event.url

	if val != want {
		t.Errorf("building PublicURL: want %s, got %s", want, val)
		t.Fail()
	}
}

func TestGetEvent_ReadSecrets(t *testing.T) {

	valSt := []string{"s1", "s2"}
	val, _ := json.Marshal(valSt)
	os.Setenv("Http_Secrets", string(val))
	owner := "alexellis"
	os.Setenv("Http_Owner", owner)

	eventInfo, err := getEvent()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	expected := []string{owner + "-s1", owner + "-s2"}
	for _, val := range eventInfo.secrets {
		found := false
		for _, expectedVal := range expected {
			if expectedVal == val {
				found = true
			}
		}
		if !found {
			t.Errorf("Wanted secret: %s, didn't find it in list", val)
		}
	}
}

func TestGetEvent_EmptyEnvVars(t *testing.T) {
	_, err := getEvent()

	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

}
