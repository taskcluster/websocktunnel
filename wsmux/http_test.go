package wsmux

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

type wrapStream struct {
	*bytes.Buffer
}

func TestGet(t *testing.T) {
	server := httptest.NewServer(genWebSocketHandler(t, wsConn))
	servURL := server.URL
	defer server.Close()
	conn, _, err := websocket.DefaultDialer.Dial(makeWsURL(servURL), nil)
	if err != nil {
		t.Fatal(err)
	}
	session := Client(conn, Config{Log: genLogger("get-test")})
	// session.readDeadline = time.Now().Add(10 * time.Second)
	req, err := http.NewRequest(http.MethodGet, "", nil)
	stream, err := session.Open()
	if err != nil {
		t.Fatal(err)
	}
	err = req.Write(stream)
	if err != nil {
		t.Fatal(err)
	}
	reader := bufio.NewReader(stream)
	resp, err := http.ReadResponse(reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte(getSuccess)) {
		t.Log(bytes.NewBuffer(b).String())
		t.Fatal("message inconsistent")
	}
}

func TestPost(t *testing.T) {
	server := httptest.NewServer(genWebSocketHandler(t, wsConn))
	servURL := server.URL
	defer server.Close()
	conn, _, err := (&websocket.Dialer{}).Dial(makeWsURL(servURL), nil)
	if err != nil {
		t.Fatal(err)
	}
	session := Client(conn, Config{Log: genLogger("post-test")})
	// session.readDeadline = time.Now().Add(10 * time.Second)
	msg := []byte("message to be sent in a post request")
	buffer := new(bytes.Buffer)
	_, _ = buffer.Write(msg)
	req, err := http.NewRequest(http.MethodPost, "", wrapStream{buffer})
	stream, err := session.Open()
	if err != nil {
		t.Fatal(err)
	}
	err = req.Write(stream)
	if err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(stream)
	resp, err := http.ReadResponse(reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, msg) {
		t.Log(bytes.NewBuffer(b).String())
		t.Fatal("message inconsistent")
	}
}

func TestMultiplePost(t *testing.T) {
	server := httptest.NewServer(genWebSocketHandler(t, wsConn))
	servURL := server.URL
	defer server.Close()
	conn, _, err := (&websocket.Dialer{}).Dial(makeWsURL(servURL), nil)
	if err != nil {
		t.Fatal(err)
	}
	session := Client(conn, Config{Log: genLogger("post-test")})

	var wg sync.WaitGroup
	sendAndWait := func() {
		msg := []byte("message to be sent in a post request")
		buffer := new(bytes.Buffer)
		_, _ = buffer.Write(msg)
		defer wg.Done()

		req, err := http.NewRequest(http.MethodPost, "", wrapStream{buffer})
		stream, err := session.Open()
		if err != nil {
			t.Fatal(err)
		}
		err = req.Write(stream)
		if err != nil {
			t.Fatal(err)
		}

		reader := bufio.NewReader(stream)
		resp, err := http.ReadResponse(reader, nil)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(b, msg) {
			t.Log(bytes.NewBuffer(b).String())
			t.Fatal("message inconsistent")
		}
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go sendAndWait()
	}
	wg.Wait()
}

func TestWebSocket(t *testing.T) {
	// t.Skip("No idea why this is failing")
	server := httptest.NewServer(genWebSocketHandler(t, wsConn))
	servURL := server.URL
	defer server.Close()
	conn, _, err := (&websocket.Dialer{}).Dial(makeWsURL(servURL), nil)
	//runtime.Breakpoint()
	if err != nil {
		t.Fatal(err)
	}
	session := Client(conn, Config{Log: genLogger("ws-test")})
	//runtime.Breakpoint()
	// session.readDeadline = time.Now().Add(10 * time.Second)
	wsURL := &url.URL{Host: "tcproxy.net", Scheme: "ws"}
	stream, err := session.Open()
	if err != nil {
		t.Fatal(err)
	}
	//runtime.Breakpoint()
	ws, _, err := websocket.NewClient(stream, wsURL, nil, 1024, 1024)
	if err != nil {
		t.Fatal(err)
	}
	_, b, err := ws.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, []byte(wsSuccess)) {
		t.Log(bytes.NewBuffer(b).String())
		t.Fatal("message inconsistent")
	}

}
