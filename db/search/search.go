package search

import (
	"github.com/agnivade/levenshtein"
	"strings"
	"time"
	"torrsru/db/sync"
	"torrsru/models/fdb"
	"unicode"
)

var (
	isUpdate   bool
	lastUpdate time.Time
	index      map[string]struct{}
)

func SetUpdate() { isUpdate = true }

func Find(query string, accurate bool) []*fdb.Torrent {
	if lastUpdate.Add(time.Hour).Before(time.Now()) {
		isUpdate = true
	}
	if isUpdate {
		UpdateIndex()
	}

	listRes := map[string]bool{}
	query = cleanString(query)
	for s, _ := range index {
		if strings.Contains(s, query) {
			listRes[s] = true
		}
	}

	if accurate {
		passDst := 1
		for s := range listRes {
			arr := strings.Split(s, ":")
			if len(arr) == 0 {
				continue
			}
			if len(arr) < 2 {
				arr = append(arr, arr[0])
			}
			if len(arr[1]) == 0 {
				arr[1] = arr[0]
			}
			dst1 := levenshtein.ComputeDistance(query, arr[0])
			dst2 := levenshtein.ComputeDistance(query, arr[1])
			if dst2 < dst1 {
				dst1 = dst2
			}

			if dst1 > passDst {
				listRes[s] = false
			}
		}
	}

	var listTorr []*fdb.Torrent
	for s, b := range listRes {
		if b {
			trs := sync.GetTorrents(s)
			listTorr = append(listTorr, trs...)
		}
	}
	return listTorr
}

func UpdateIndex() {
	list := sync.ListNames()
	index = make(map[string]struct{})
	for _, s := range list {
		index[s] = struct{}{}
	}
	lastUpdate = time.Now()
	isUpdate = false
}

func cleanString(input string) string {
	var cleaned string
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned += string(char)
		}
	}
	return strings.ToLower(cleaned)
}
