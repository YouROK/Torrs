package static

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteStaticFiles(route *gin.Engine) {

	route.GET("/img/T.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(FilesimgTpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", FilesimgTpng)
	})

	route.GET("/img/android-chrome-192x192.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidchrome192x192png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidchrome192x192png)
	})

	route.GET("/img/android-chrome-512x512.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidchrome512x512png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidchrome512x512png)
	})

	route.GET("/img/apple-touch-icon.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappletouchiconpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappletouchiconpng)
	})

	route.GET("/img/browserconfig.xml", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgbrowserconfigxml))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/xml; charset=utf-8", Filesimgbrowserconfigxml)
	})

	route.GET("/img/copy.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgcopysvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesimgcopysvg)
	})

	route.GET("/img/down.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgdownsvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesimgdownsvg)
	})

	route.GET("/img/favicon-16x16.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfavicon16x16png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgfavicon16x16png)
	})

	route.GET("/img/favicon-32x32.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfavicon32x32png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgfavicon32x32png)
	})

	route.GET("/img/favicon.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfaviconico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgfaviconico)
	})

	route.GET("/img/ico/anidub.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicoanidubico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicoanidubico)
	})

	route.GET("/img/ico/anifilm.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicoanifilmico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicoanifilmico)
	})

	route.GET("/img/ico/anilibria.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicoanilibriaico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicoanilibriaico)
	})

	route.GET("/img/ico/animelayer.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicoanimelayerico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicoanimelayerico)
	})

	route.GET("/img/ico/baibako.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicobaibakoico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicobaibakoico)
	})

	route.GET("/img/ico/bitru.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicobitruico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicobitruico)
	})

	route.GET("/img/ico/hdrezka.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicohdrezkaico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicohdrezkaico)
	})

	route.GET("/img/ico/kinozal.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicokinozalico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicokinozalico)
	})

	route.GET("/img/ico/lostfilm.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicolostfilmico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicolostfilmico)
	})

	route.GET("/img/ico/megapeer.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicomegapeerico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicomegapeerico)
	})

	route.GET("/img/ico/nnmclub.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgiconnmclubico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgiconnmclubico)
	})

	route.GET("/img/ico/rutor.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicorutorico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicorutorico)
	})

	route.GET("/img/ico/rutracker.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicorutrackerico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicorutrackerico)
	})

	route.GET("/img/ico/selezen.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicoselezenico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicoselezenico)
	})

	route.GET("/img/ico/toloka.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicotolokaico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicotolokaico)
	})

	route.GET("/img/ico/torrentby.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicotorrentbyico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicotorrentbyico)
	})

	route.GET("/img/ico/underverse.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgicounderverseico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgicounderverseico)
	})

	route.GET("/img/magnet.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmagnetsvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesimgmagnetsvg)
	})

	route.GET("/img/mstile-150x150.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmstile150x150png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgmstile150x150png)
	})

	route.GET("/img/plus.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgplussvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesimgplussvg)
	})

	route.GET("/img/safari-pinned-tab.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgsafaripinnedtabsvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesimgsafaripinnedtabsvg)
	})

	route.GET("/img/site.webmanifest", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgsitewebmanifest))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/manifest+json", Filesimgsitewebmanifest)
	})

	route.GET("/img/up.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgupsvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesimgupsvg)
	})

	route.GET("/js/crypt/base64.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjscryptbase64js))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjscryptbase64js)
	})

	route.GET("/js/crypt/jsbn.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjscryptjsbnjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjscryptjsbnjs)
	})

	route.GET("/js/crypt/prng4.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjscryptprng4js))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjscryptprng4js)
	})

	route.GET("/js/crypt/rng.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjscryptrngjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjscryptrngjs)
	})

	route.GET("/js/crypt/rsa.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjscryptrsajs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjscryptrsajs)
	})

	route.GET("/js/crypt/sha1.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjscryptsha1js))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjscryptsha1js)
	})
}
