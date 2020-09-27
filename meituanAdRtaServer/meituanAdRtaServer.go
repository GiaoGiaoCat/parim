package meituanAdRtaServer

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"github.com/liushooter/parim/meituanAdRta"
	"io/ioutil"
	"log"
	"net/http"
)

func Request(url string, rtaRequest *meituanAdRta.RtaRequest) (*meituanAdRta.RtaResponse, error) {
	payload, err := proto.Marshal(rtaRequest)
	if err != nil {
		log.Fatal("RtaRequest marshaling error: ", err)
		return nil, err
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
		return nil, err
	}

	return rtaResponse, nil
}
