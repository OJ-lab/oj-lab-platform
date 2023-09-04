package tests

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/OJ-lab/oj-lab-services/packages/application"
)

func TestMinio(T *testing.T) {
	// Initialize minio client object.
	minioClient := application.GetMinioClient()

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	// Make a new bucket called mymusic.
	bucketName := "oj-lab-problem-packages"

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload package files
	packagePath := "../test-collection/packages/icpc/hello_world"
	filepath.Walk(packagePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		relativePath := filepath.Join(filepath.Base(packagePath), strings.Replace(path, packagePath, "", 1))
		println(relativePath)
		_, minioErr := minioClient.FPutObject(ctx, bucketName,
			relativePath,
			path,
			minio.PutObjectOptions{})
		if minioErr != nil {
			log.Fatalln(minioErr)
		}
		return minioErr
	})
}
