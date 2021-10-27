/* Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4". */
package main

import "fmt"

type IPAddr [4]byte

func (iAdd IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", iAdd[0], iAdd[1], iAdd[2], iAdd[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {7, 7, 7, 7},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", ip, name)
	}
}
