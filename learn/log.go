/*
日志
 */

package main

import "log"

func init() {
	log.SetPrefix("[UserCenter]")
	log.SetFlags(log.Ldate|log.Lshortfile)
}

func main() {
	log.Fatal("Hello")
}
