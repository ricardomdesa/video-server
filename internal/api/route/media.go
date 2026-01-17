package route

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
	log "github.com/sirupsen/logrus"
)

type Repository interface {
	SaveMediaData(ctx context.Context, key string, data interface{}) error
	GetMediaData(ctx context.Context, key string) ([]byte, error)
}

func MediaRouter(group *gin.RouterGroup, s3Repo *persistence.S3Repository) {
	group.GET("/:mod/:id/stream",func(c *gin.Context) {
		streamHandler(c, s3Repo)
	})
	group.GET("/:mod/:id/:segName", func(c *gin.Context) {
		streamHandler(c, s3Repo)
	})
}

func streamHandler(c *gin.Context, s3Repo *persistence.S3Repository) {
	ID := c.Param("id")
	mod := c.Param("mod")
	log.Infof("ID received %v", ID)

	segName := c.Param("segName")
	log.Infof("segName received %v", segName)

	mediaKey := getMediaKey(ID, mod, segName)
	mediaData, err := s3Repo.GetMediaData(c, mediaKey)
	if err != nil {
		log.Errorf("Failed to fetch media data: %v", err)
		c.JSON(500, gin.H{"error": "Failed to fetch media data"})
		return
	}

	if segName == "" {
		serveHlsM3u8(c, mediaData)
	} else {
		serveHlsTs(c, mediaData)
	}
}

func getMediaKey(mId, mod, segName string) string {
	if segName == "" {
		return mod + ":" + mId + ":index.m3u8"
	}
	return mod + ":" + mId + ":" + segName
}

func serveHlsM3u8(c *gin.Context, mediaData []byte) {
	c.Set("Content-Type", "application/x-mpegURL")
	c.Data(200, "application/x-mpegURL", mediaData)
}

func serveHlsTs(c *gin.Context, mediaData []byte) {
	c.Set("Content-Type", "video/MP2T")
	c.Data(200, "video/MP2T", mediaData)
}
