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

package check

import "os"

func Exists(fileName string) bool {

	_, err := os.Stat(fileName)

	return !os.IsNotExist(err)
}
