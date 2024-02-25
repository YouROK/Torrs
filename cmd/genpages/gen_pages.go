package main

import (
	"fmt"
	"io/fs"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	staticFilesPath := "public"
	srcFilesPath := "web/static/files"
	srcPagesPath := "web/static"

	os.RemoveAll(srcPagesPath)
	os.MkdirAll(srcPagesPath, 0777)
	run("cp", "-r", staticFilesPath, srcFilesPath)

	files := make([]string, 0)

	filepath.WalkDir(staticFilesPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			name := strings.TrimPrefix(path, staticFilesPath)
			if !strings.HasPrefix(filepath.Base(name), ".") {
				files = append(files, name)
				fmt.Println("Find file:", name)
			}
		}
		return nil
	})
	sort.Strings(files)
	fmap := writeEmbed(filepath.Join(srcPagesPath, "html.go"), files)
	writeRoute(filepath.Join(srcPagesPath, "route.go"), fmap)
}

func writeEmbed(fname string, files []string) map[string]string {
	ff, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ff.Close()
	embedStr := `package static

import (
	_ "embed"
)
`
	ret := make(map[string]string)

	for _, f := range files {
		path := "files" + f
		name := cleanName(path)
		embedStr += "\n//go:embed " + path + "\nvar " + name + " []byte\n"
		ret[f] = name
	}

	ff.WriteString(embedStr)
	return ret
}

func writeRoute(fname string, fmap map[string]string) {
	ff, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ff.Close()
	embedStr := `package static

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteStaticFiles(route *gin.Engine) {
`
	mime.AddExtensionType(".map", "application/json")
	mime.AddExtensionType(".webmanifest", "application/manifest+json")
	// sort fmap
	keys := make([]string, 0, len(fmap))
	for key := range fmap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, link := range keys {
		fmime := mime.TypeByExtension(filepath.Ext(link))
		if fmime == "application/xml" || fmime == "application/javascript" {
			fmime = fmime + "; charset=utf-8"
		}
		if fmime == "image/x-icon" {
			fmime = "image/vnd.microsoft.icon"
		}
		embedStr += `
	route.GET("` + link + `", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(` + fmap[link] + `))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "` + fmime + `", ` + fmap[link] + `)
	})
`
	}
	embedStr += "}\n"

	ff.WriteString(embedStr)
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func cleanName(fn string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Title(reg.ReplaceAllString(fn, ""))
}
