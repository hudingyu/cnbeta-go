/*
 * @Description: 
 * @Author: hudingyu
 * @Date: 2019-09-26 22:47:26
 * @LastEditTime: 2019-09-26 22:47:26
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}
