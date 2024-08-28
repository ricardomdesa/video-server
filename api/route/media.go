package route

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func MediaRouter(group *gin.RouterGroup) {
	group.GET("/:id/stream", streamHandler)
	group.GET("/:id/:segName", streamHandler)
}

func streamHandler(c *gin.Context) {
	ID := c.Param("id")
	log.Infof("ID received %v", ID)
	mId, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error get ID": err.Error()})
		return
	}
	segName := c.Param("segName")
	fmt.Println(segName)
	log.Infof("segName received %v", segName)

	if segName == "" {
		mediaBase := getMediaBase(mId)
		m3u8Name := "index.m3u8"
		serveHlsM3u8(c, mediaBase, m3u8Name)
	} else {
		mediaBase := getMediaBase(mId)
    serveHlsTs(c, mediaBase, segName)
	}
}

func getMediaBase(mId int) string {
	mediaRoot := "assets/media"
	return fmt.Sprintf("%s/%d", mediaRoot, mId)
}

func serveHlsM3u8(c *gin.Context, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, m3u8Name)
	fmt.Printf("media file: %s\n", mediaFile)
	c.Set("Content-Type", "application/x-mpegURL")
	c.File(mediaFile)
}

func serveHlsTs(c *gin.Context, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, segName)
	c.Set("Content-Type", "video/MP2T")
	c.File(mediaFile)
}
