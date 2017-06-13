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

package wcfg

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strings"

	check "github.com/jeffotoni/s3goamz/pkg/check"
)

type Credentials struct {
	Key    string
	Secret string
}

//
//
//
var (

	//
	//Object of type Config that
	//we will use to access our struct
	//
	cfg      Credentials
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	DirConf  = os.Getenv("HOME") + "/.aws/"
	//DirConf  = "./"
	NameConf = "credentials"
	pconf    = DirConf + "" + NameConf
	returns  string
)

//
//
//
func AwsKeys() (string, string) {

	//
	//
	//
	wcfg := GetConfig()

	//
	// config data
	//

	if wcfg.Key != "" && wcfg.Secret != "" {

		return wcfg.Key, wcfg.Secret
	} else {

		return "error key", "error Secret"
	}

}

//
//
//
func GetConfig() *Credentials {

	//
	// file exist ?
	//

	if check.Exists(pconf) == true {

		if err = ReadFileInto(&cfg, pconf); err != nil {
			log.Fatal(err)
		}

		return &cfg

	} else {

		return &cfg
	}

}

func ReadFileInto(C *Credentials, pconf string) error {

	if file, errx := os.Open(pconf); errx == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			line := scanner.Text()
			line = strings.TrimSpace(line)
			line = strings.Trim(line, " ")

			vetor := strings.Split(line, "=")
			keys := strings.Trim(vetor[0], " ")
			keys = strings.ToLower(keys)

			if keys == "aws_access_key_id" || keys == "aws_secret_access_key" {

				if keys == "aws_access_key_id" {

					C.Key = strings.TrimSpace(vetor[1])

				} else if keys == "aws_secret_access_key" {

					C.Secret = strings.TrimSpace(vetor[1])
				}
			} else if keys == "source_profile" {

				//
				// [default]
				// source_profile = default
				//
				// Try to catch default aws per instance
				//

			}
		}

		return nil

	} else {

		return errx
	}
}

func UserHomeDir() string {

	env := "HOME"

	if runtime.GOOS == "windows" {

		env = "USERPROFILE"

	} else if runtime.GOOS == "plan9" {

		env = "home"
	}

	return os.Getenv(env)
}
