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

package auth

import (
	//_ "./pkg/toml"
	"fmt"
	"os"

	wcfg "github.com/jeffotoni/s3goamz/pkg/wcfg"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

type S struct {
	s3 *s3.S3
}

func (s *S) GetS3() *s3.S3 {

	return s.s3
}

// aws.USEast
// aws.USWest
// aws.USWest2
// aws.EUWest
// aws.APSoutheast
// aws.APSoutheast2
// aws.APNortheast
// aws.SAEast
// aws.CNNorth

//
//
//
func GetAuth(parameter ...string) *s3.S3 {

	// var s *S
	var AccessKey string
	var SecretKey string

	if len(parameter) == 2 {

		//
		// ok
		//
		AccessKey = parameter[0]
		SecretKey = parameter[1]

	} else if len(parameter) == 0 {

		//
		// Looking for keys from the aws emr ./aws/credentials default file
		//
		AccessKey, SecretKey = wcfg.AwsKeys()

		//
		// get Env
		//
		if AccessKey == "error key" && SecretKey == "error Secret" {

			AccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
			SecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

			if AccessKey == "" && SecretKey == "" {

				fmt.Println("Error I could not find your aws keys!")
				os.Exit(0)
			}

		}

	} else {

		fmt.Println("Error de parameter!!")
		os.Exit(0)
	}

	if AccessKey != "error key" && SecretKey != "error Secret" {

		auth := aws.Auth{

			//
			//
			//
			AccessKey: AccessKey,

			//
			// change this to yours
			//
			SecretKey: SecretKey,
		}

		//fmt.Println(auth)

		connect := s3.New(auth, aws.USEast)

		return connect
	}

	connect := s3.New(aws.Auth{}, aws.USEast)

	return connect
}
