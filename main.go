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
	//_ "./pkg/auth"

	"fmt"

	auth "github.com/jeffotoni/s3goamz/pkg/auth"
	"launchpad.net/goamz/s3"
)

const (
	Bucket = "name-your-bucket"
	//Region = aws.USEast
)

func main() {

	//
	// GetAuth
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

	// bucket := connection.Bucket("teste-user") // change this your bucket name

	fmt.Println(conn)

	//
	// File upload..
	//
	//fileToBeUploaded := "namefile.pdf"

}
