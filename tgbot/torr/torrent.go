package torr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"torrsru/global"
	"torrsru/tgbot/torr/state"
)

type TorrentDetails struct {
	Title   string
	Size    string
	Date    time.Time
	Link    string
	Tracker string
	Peer    int
	Seed    int
	Magnet  string
}

func GetTorrentInfo(hash string) (*state.TorrentStatus, error) {
	link := global.TSHost + "/stream?stat&link=" + url.QueryEscape(hash)
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ti *state.TorrentStatus

	err = json.Unmarshal(buf, &ti)
	return ti, err
}
