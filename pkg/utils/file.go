package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/consolelabs/indexer-api/pkg/logger"
)

func DownloadFileByUri(path, fileName, uri string, maxRetries int64) error {
	if maxRetries <= 0 {
		return fmt.Errorf("[DownloadFileByUri] failed to fetch content from uri %s", uri)
	}

	res, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// retry if timeout
	if res.StatusCode != http.StatusOK {
		logger.L.Fields(logger.Fields{
			"path":     path,
			"fileName": fileName,
			"uri":      uri,
		}).Errorf(err, "[DownloadFileByUri] http.Get() failed with code %d", res.StatusCode)
		return DownloadFileByUri(path, fileName, uri, maxRetries-1)
	}

	// create dir if not existed yet
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return fmt.Errorf("[DownloadFileByUri] os.Mkdir failed: %v\n", err)
		}
	}

	// open a file for writing
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("[DownloadFileByUri] os.Create failed: %v\n", err)
	}

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("[DownloadFileByUri] io.Copy failed: %v\n", err)
	}
	return nil
}

// func ResizeImage(filePath string) error {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	fStat, _ := f.Stat()
// 	// no need resizing if size < 10KB
// 	if fStat.Size() <= 10000 {
// 		return nil
// 	}

// 	buffer, err := io.ReadAll(f)
// 	if err != nil {
// 		return err
// 	}

// 	contentType := http.DetectContentType(buffer)
// 	var imageType bimg.ImageType
// 	switch contentType {
// 	case "image/png":
// 		imageType = bimg.PNG
// 	case "image/jpeg":
// 		imageType = bimg.JPEG
// 	case "image/gif":
// 		imageType = bimg.UNKNOWN
// 	default:
// 		imageType = bimg.UNKNOWN
// 	}
// 	if imageType == bimg.UNKNOWN {
// 		return fmt.Errorf("[ProcessImage] unsupported image file type: %s", contentType)
// 	}

// 	converted, err := bimg.NewImage(buffer).Convert(imageType)
// 	if err != nil {
// 		return fmt.Errorf("[ProcessImage] failed to convert image format:%v\n", err)
// 	}

// 	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: 72})
// 	if err != nil {
// 		return fmt.Errorf("[ProcessImage] failed to process image%v\n:", err)
// 	}

// 	err = bimg.Write(filePath, processed)
// 	if err != nil {
// 		return fmt.Errorf("[ProcessImage] bimg.Write failed: %v\n", err)
// 	}

// 	return nil
// }
