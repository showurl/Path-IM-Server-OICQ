package tests

import (
	"bytes"
	"encoding/json"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/types"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	buf, _ := json.Marshal(types.LoginReq{
		Username: "user-01",
		Password: "123456",
	})
	request, err := http.NewRequest("POST", "http://42.194.149.177:8080/v1/white/login", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	t.Log(response.StatusCode)
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(all))
}
