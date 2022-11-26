package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	p  = fmt.Println
	sf = fmt.Sprintf

	topicName string = "alarm"
	server    string = "http://127.0.0.1:7666"
)

func main() {
	p(">>>>>>>>>>>>>> HttpMessageServer <<<<<<<<<<<<<<<<<<<<")
	defer p("Bye bye. never mind i'll find someone like you")

	go func() {
		for {
			err := SetMessage(sf("%v", time.Now()))

			if err != nil {
				p("[ERROR] Set", err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		msg, err := GetMessage()
		if err != nil {
			p("[ERROR] Get", err)
			time.Sleep(time.Second)
			continue
		}

		p(sf("rcv msg - %v", msg))
		// break
	}

}

func GetMessage() (string, error) {
	var msg string
	resp, err := http.Get(sf("%s/GetMessage?topic=%s&timeout=-1", server, topicName))
	if err != nil {
		return msg, err
	}

	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg, err
	}

	return string(resByte), err
}

func SetMessage(msg string) error {

	resp, err := http.Get(sf("%s/SetMessage?topic=%s&timeout=-1&message=%s",
		server, topicName, url.QueryEscape(msg)))
	if err != nil {
		return err
	}

	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(sf("status : %v, msg : %v", resp.StatusCode, string(resByte)))
	}

	return err
}
