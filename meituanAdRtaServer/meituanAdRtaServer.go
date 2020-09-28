package meituanAdRtaServer

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"github.com/liushooter/parim/meituanAdRta"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Request(url string, rtaRequest *meituanAdRta.RtaRequest) (*meituanAdRta.RtaResponse, error) {
	payload, err := proto.Marshal(rtaRequest)
	if err != nil {
		log.Fatal("RtaRequest marshaling error: ", err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("post request error: ", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-protobuf;charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("send post request error: ", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	rtaResponse := &meituanAdRta.RtaResponse{}
	err = proto.Unmarshal(body, rtaResponse)
	if err != nil {
		log.Fatal("RtaResponse unmarshaling error: ", err)
		return nil, err
	}

	return rtaResponse, nil
}
