/***********
*
*
* project s3 Upload
*
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	auth "github.com/jeffotoni/s3goamz/pkg/auth"
	check "github.com/jeffotoni/s3goamz/pkg/check"
	erro "github.com/jeffotoni/s3goamz/pkg/erro"
	"launchpad.net/goamz/s3"
)

const (
	Bucket     = "name-your-bucket"
	FileUpload = "filename.pdf"
)

const (
	Private           = s3.ACL("private")
	PublicRead        = s3.ACL("public-read")
	PublicReadWrite   = s3.ACL("public-read-write")
	AuthenticatedRead = s3.ACL("authenticated-read")
	BucketOwnerRead   = s3.ACL("bucket-owner-read")
	BucketOwnerFull   = s3.ACL("bucket-owner-full-control")
)

const (
	fileChunk = 5 * (1 << 20) // 5MB
)

func main() {

	//
	//
	//
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	//
	//
	//
	yellow := color.New(color.FgYellow)
	boldYellow := yellow.Add(color.Bold)

	//
	// GetAuth (key, secret)
	//
	// connect := auth.GetAuth("xxxxxx-key", "xxxxx-secret")

	//
	// GetAuth
	//
	connect := auth.GetAuth()

	//fmt.Println(connect)

	conn := s3.Bucket{
		S3:   connect,
		Name: Bucket,
	}

	if check.Exists(FileUpload) == true {

		//
		//
		//
		file, errx := os.Open(FileUpload)
		erro.Check(errx)
		defer file.Close()

		//
		//
		//
		fileInfo, _ := file.Stat()

		//
		//
		//
		fileSize := fileInfo.Size()

		if fileSize < fileChunk {

			fmt.Println("Error, The file has to be larger than 5mb to send in parts to amazon s3.")
			os.Exit(0)

		}
		//
		//
		//
		bytes := make([]byte, fileSize)

		// read into buffer
		buffer := bufio.NewReader(file)

		//
		//
		//
		_, errx = buffer.Read(bytes)

		//
		//
		//
		erro.Check(errx)

		//
		//
		//
		filetype := http.DetectContentType(bytes)

		//
		//
		//
		multi, err := conn.InitMulti("/"+FileUpload, filetype, BucketOwnerRead)

		//
		//
		//
		erro.Check(err)

		//
		//
		//
		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		//
		//
		//
		parts := []s3.Part{}

		//
		//
		//
		HeaderPart := strings.NewReader(string(bytes))

		//
		//
		//
		sizeTotal := (fileSize / (1024 * 1024))

		//
		//
		//
		chunkPart := (fileChunk / (1024 * 1024))

		//
		//
		//
		boldRed.Println("File Size:", fileSize, "byte", sizeTotal, "Mb", "Chunk:", chunkPart, "Mb")

		//
		//
		//
		boldYellow.Println("Uploading...")

		//
		//
		//
		for i := uint64(0); i < totalPartsNum; i++ {

			//
			//
			//
			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))

			//
			//
			//
			partBuffer := make([]byte, partSize)

			//
			//
			//
			sizePart, errx2 := io.ReadFull(HeaderPart, partBuffer)

			//
			//
			//
			erro.Check(errx2)

			//
			//
			//
			piece, errx3 := multi.PutPart(int(i)+1, strings.NewReader(string(partBuffer))) // write to S3 bucket part by part

			//
			//
			//
			erro.Check(errx3)

			//
			//
			//
			fmt.Printf("Processing %d piece of %d and uploaded %d bytes.\n ", int(i), int(totalPartsNum), int(sizePart))

			//
			//
			//
			parts = append(parts, piece)
		}

		//
		//
		//
		err = multi.Complete(parts)

		//
		//
		//
		erro.Check(err)

		boldYellow.Println("\n\nPutPart upload completed")

	} else {

		boldRed.Println("Erro, File does not exist!")
		os.Exit(0)
	}

	// bucket := connection.Bucket("name-your-bucket") // change this your bucket name

}
