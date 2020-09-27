package meituanAdRtaServer

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/liushooter/parim/meituanAdRta"
	"io/ioutil"
	"log"
	"net/http"
)

func Request(url string, rtaRequest meituanAdRta.RtaRequest) meituanAdRta.RtaResponse {
	payload, err := proto.Marshal(rtaRequest)
	if err != nil {
		log.Fatal("RtaRequest marshaling error: ", err)
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/x-protobuf;charset=UTF-8")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	rtaResponse := &meituanAdRta.RtaResponse{}
	err = proto.Unmarshal(body, rtaResponse)

	if err != nil {
		log.Fatal("RtaResponse unmarshaling error: ", err)
		panic(err)
	}

	return rtaResponse

}
