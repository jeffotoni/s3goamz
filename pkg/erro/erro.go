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

package erro

import (
	"fmt"
	"os"
)

//
//
//
func Check(err error) {

	if err != nil {

		fmt.Println(err)
		os.Exit(0)
		//panic(err)
	}
}
