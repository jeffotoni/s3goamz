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
	auth "github.com/jeffotoni/s3goamz/pkg/auth"
	"launchpad.net/goamz/s3"
)

const (
	Bucket = "name-your-bucket"
)

func main() {

	//
	// GetAuth (region, key, secret)
	//
	// connect := auth.GetAuth("us-east-1","xxxxxx-key", "xxxxx-secret")

	//
	// GetAuth
	//
	connect := auth.GetAuth()

	//fmt.Println(connect)

	conn := s3.Bucket{
		S3:   connect,
		Name: Bucket,
	}

	// bucket := connection.Bucket("name-your-bucket") // change this your bucket name

	//fmt.Println(conn)

	//
	// File upload..
	//
	//fileToBeUploaded := "namefile.pdf"

}
