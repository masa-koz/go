/* IP
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr, err := net.ResolveTCPAddr("tcp", name)
	if err != nil {
		fmt.Println("ResolveTCPaddr:", err.Error())
		os.Exit(1)
	}
	fmt.Println("addr is", addr.String())
}
