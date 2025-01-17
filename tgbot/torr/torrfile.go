package torr

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"
	"torrsru/global"
	"torrsru/tgbot/torr/state"
)

var ERR_STOPPED = errors.New("stopped")

type TorrFile struct {
	hash   string
	name   string
	wrk    *Worker
	offset int64
	size   int64
	id     int

	resp *http.Response
}

func NewTorrFile(wrk *Worker, tfile *state.TorrentFileStat) (*TorrFile, error) {
	if tfile.Length > 2*1024*1024*1024 {
		return nil, errors.New("Размер файла должен быть меньше 2GB")
	}

	tf := new(TorrFile)
	tf.hash = wrk.torrentHash
	tf.name = filepath.Base(tfile.Path)
	tf.wrk = wrk
	tf.size = tfile.Length

	link := global.TSHost + "/stream?link=" + url.QueryEscape(wrk.torrentHash) + "&index=" + strconv.Itoa(tfile.Id) + "&play"
	c := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 5 * time.Minute,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	resp, err := c.Get(link)
	if err != nil {
		return nil, err
	}
	tf.resp = resp
	return tf, nil
}

func (t *TorrFile) Read(p []byte) (n int, err error) {
	if t.wrk.isCancelled {
		return 0, ERR_STOPPED
	}
	n, err = t.resp.Body.Read(p)
	t.offset += int64(n)
	return
}

func (t *TorrFile) Loaded() int64 {
	return t.size - t.offset
}

func (t *TorrFile) Close() {
	if t.resp != nil && t.resp.Body != nil {
		t.resp.Body.Close()
		t.resp = nil
	}
}
