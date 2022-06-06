package common

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

const BUCKETNAME = "dousheng"

func InitMinioClient() *minio.Client {
	endpoint := "178.79.130.90:9000"
	accessKeyID := "doushengadmin"
	secretAccessKey := "doushengadmin"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("创建 MinIO 客户端失败", err)
		return nil
	}

	return minioClient

}
