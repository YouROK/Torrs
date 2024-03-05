package sync

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"torrsru/models/fdb"
	"torrsru/web/global"
)

var (
	mu     sync.Mutex
	isSync bool
)

func StartSync() {
	for !global.Stoped {
		syncDB()
		time.Sleep(time.Minute * 20)
	}
}

func syncDB() {
	mu.Lock()
	if isSync {
		mu.Unlock()
		return
	}
	isSync = true
	defer func() { isSync = false }()

	filetime := GetFileTime()

	mu.Unlock()
	for {
		ftstr := strconv.FormatInt(filetime, 10)
		//log.Println("Get:", ftstr)
		resp, err := http.Get("http://85.17.54.98:9117/sync/fdb/torrents?time=" + ftstr)
		if err != nil {
			log.Fatal("Error connect to fdb:", err)
			return
		}

		var js *fdb.FDBRequest
		err = json.NewDecoder(resp.Body).Decode(&js)
		if err != nil {
			log.Fatal("Error decode json:", err)
			return
		}
		resp.Body.Close()

		err = saveTorrent(js.Collections)
		if err != nil {
			log.Fatal("Error save torrents:", err)
			return
		}

		err = SetFileTime(filetime)
		if err != nil {
			log.Fatal("Error set ftime:", err)
			return
		}
		log.Println("Save:", ftstr)

		if !js.Nextread {
			break
		}

		for _, col := range js.Collections {
			if col.Value.FileTime > filetime {
				filetime = col.Value.FileTime
			}
		}
		js = nil
	}
}

func getHash(magnet string) string {
	pos := strings.Index(magnet, "btih:")
	if pos == -1 {
		return ""
	}
	magnet = magnet[pos+5:]
	pos = strings.Index(magnet, "&")
	if pos == -1 {
		return strings.ToLower(magnet)
	}
	return strings.ToLower(magnet[:pos])
}
