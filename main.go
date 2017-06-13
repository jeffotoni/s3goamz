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
		// Private           = ACL("private")
		// PublicRead        = ACL("public-read")
		// PublicReadWrite   = ACL("public-read-write")
		// AuthenticatedRead = ACL("authenticated-read")
		// BucketOwnerRead   = ACL("bucket-owner-read")
		// BucketOwnerFull   = ACL("bucket-owner-full-control")
		// set up for multipart upload
		multi, err := conn.InitMulti("/"+FileUpload, filetype, BucketOwnerRead)

	} else {

		fmt.Prinln("Erro, File does not exist!")
		os.Exit(0)
	}

	// bucket := connection.Bucket("name-your-bucket") // change this your bucket name

}
