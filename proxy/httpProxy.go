package proxy

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"tiny-proxy/parser"
)

func HttpProxyHandle(conn net.Conn) {

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Read from conn error | %v", err)
		return
	}

	httpProxy(n, buf, conn)

	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
}

func httpProxy(bufLen int, buf []byte, client net.Conn) {

	req := parser.ParseHttpReqMsg(string(buf[:bufLen]))
	url := req.URL.String()

	log.Printf("client data | %v", req)

	var forwardReq *http.Request
	if strings.EqualFold("GET", req.Method) {
		forwardReq, _ = http.NewRequest("GET", url, nil)
	} else if strings.EqualFold("POST", req.Method) {
		// todo
	} else if strings.EqualFold("CONNECT", req.Method) {
		// todo
	}

	copyHeader(req, forwardReq)

	httpClient := http.Client{}
	resp, err := httpClient.Do(forwardReq)
	if err != nil {
		log.Printf("request remote error | %v", err)
		return
	}
	respBuf, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Printf("DumpResp error | %v", err)
		return
	}
	_, err = client.Write(respBuf)
	if err != nil {
		log.Printf("Write resp buf error | %v", err)
		return
	}
}

func copyHeader(src, tar *http.Request) {
	for s := range src.Header {
		tar.Header.Set(s, src.Header.Get(s))
	}
}
