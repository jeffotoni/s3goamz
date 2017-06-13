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
	"math"
	"net/http"
	"os"

	auth "github.com/jeffotoni/s3goamz/pkg/auth"
	check "github.com/jeffotoni/s3goamz/pkg/check"
	err "github.com/jeffotoni/s3goamz/pkg/err"
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
		file, erro := os.Open(FileUpload)
		err.checkErr(erro)
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

			fmt.Prinln("Error, The file has to be larger than 5mb to send in parts to amazon s3.")
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
		_, erro = buffer.Read(bytes)

		//
		//
		//
		err.checkErr(erro)

		//
		//
		//
		filetype := http.DetectContentType(bytes)

		//
		//
		//
		multi, errx := conn.InitMulti("/"+FileUpload, filetype, BucketOwnerRead)

		//
		//
		//
		check.checkErr(errx)

		//
		//
		//
		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		//
		//
		//
		parts := []s3.Part{}

	} else {

		fmt.Prinln("Erro, File does not exist!")
		os.Exit(0)
	}

	// bucket := connection.Bucket("name-your-bucket") // change this your bucket name

}
