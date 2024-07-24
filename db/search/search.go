package search

import (
	"github.com/agnivade/levenshtein"
	"strings"
	"torrsru/db/sync"
	"torrsru/db/utils"
	"torrsru/models/fdb"
	"torrsru/web/global"
)

var (
	index map[string]struct{}
)

func FindTitle(query string) []*fdb.Torrent {
	list := sync.ListTitles()
	find := []string{}
	queryarr := strings.Split(query, " ")
	for i := range queryarr {
		queryarr[i] = utils.ClearStr(queryarr[i])
	}
	for _, s := range list {
		isFound := true
		for _, q := range queryarr {
			if !strings.Contains(s, q) {
				isFound = false
				break
			}
		}
		if isFound {
			find = append(find, s)
		}
	}

	ret := []*fdb.Torrent{}
	for _, t := range find {
		torrs := sync.GetTorrentsByTitle(t)
		ret = append(ret, torrs...)
	}
	return ret
}

func FindName(query string, accurate bool) []*fdb.Torrent {
	if global.IsUpdateIndex {
		UpdateIndex()
	}

	listRes := map[string]bool{}
	query = utils.ClearStr(query)
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
			trs := sync.GetTorrentsByName(s)
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
	global.IsUpdateIndex = false
}
