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

package cfg

import (
	"fmt"
	"log"
	"os"
	"runtime"

	gcfg "gopkg.in/gcfg.v1"
)

type Config struct {
	Default struct {
		Aws_access_key_id     string
		Aws_secret_access_key string
	}
}

//
//
//
var (

	//
	//Object of type Config that
	//we will use to access our struct
	//
	cfg      Config
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
	toml := GetConfig()

	//
	// config data
	//
	fmt.Println(toml.Default.Aws_access_key_id)
	fmt.Println(toml.Default.Aws_secret_access_key)

	return toml.Default.Aws_access_key_id, toml.Default.Aws_secret_access_key
}

//
//
//
func GetConfig() *Config {

	//var config Config

	//
	// file exist ?
	//

	if Exists(pconf) == true {

		if err = gcfg.ReadFileInto(&cfg, pconf); err != nil {
			log.Fatal(err)
		}

		return &cfg

	} else {

		return &cfg
	}

}

func Exists(fileName string) bool {

	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
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
