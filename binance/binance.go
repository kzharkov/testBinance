package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const BaseWs string = "wss://stream.binance.com:9443/ws/"

type Binance struct {
	apiKey    string
	secretKey string
	client    *http.Client
}

func NewBinanceClient(apiKey, secretKey string) *Binance {
	return &Binance{
		apiKey:    apiKey,
		secretKey: secretKey,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (b *Binance) request(address string, method string, reqStruct interface{}, respStruct interface{}) error {
	urlAddress, err := url.Parse("https://api.binance.com" + address)
	if err != nil {
		return err
	}

	values, err := query.Values(reqStruct)
	if err != nil {
		return err
	}

	values.Add("recvWindow", "60000")
	values.Add("timestamp", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
	values.Add("signature", b.hmac(values.Encode()))

	urlAddress.RawQuery = values.Encode()
	req := &http.Request{
		Method: method,
		URL:    urlAddress,
		Header: http.Header{},
	}

	req.Header.Add("X-MBX-APIKEY", b.apiKey)
	resp, err := b.client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	errCode := &ErrorCode{}

	err = json.Unmarshal(body, errCode)
	if err != nil {
		return err
	}

	if errCode.Code != 0 {
		return errors.New(errCode.Message)
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}

	return nil
}

func (b *Binance) hmac(data string) string {
	hash := hmac.New(sha256.New, []byte(b.secretKey))
	_, err := hash.Write([]byte(data))
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
