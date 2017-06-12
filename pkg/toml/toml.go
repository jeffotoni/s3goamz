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

package toml

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//
// Info from config file
//
type Config struct {
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
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	DirConf  = os.Getenv("HOME") + "/.aws/"
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
	fmt.Println(toml.Key)
	fmt.Println(toml.Secret)

	return toml.Key, toml.Secret
}

//
//
//
func GetConfig() Config {

	var config Config

	//
	// file exist ?
	//

	if Exists(pconf) == true {

		if _, err := toml.DecodeFile(pconf, &config); err != nil {
			log.Fatal(err)
		}

		return config

	} else {

		return config
	}

}

func Exists(fileName string) bool {

	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}
