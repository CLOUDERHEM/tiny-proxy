package parser

import (
	"bufio"
	"log"
	"net/http"
	"strings"
)

func ParseHttpReqMsg(requestStr string) *http.Request {
	reader := bufio.NewReader(strings.NewReader(requestStr))

	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Printf("ParseHttpReqMsg error | %v", err)
		return nil
	}
	return req
}
