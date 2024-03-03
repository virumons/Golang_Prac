// permissions
/* example - 0644

digit - 1 - special permission
digit - 2 - owners permissions
digit - 3 - grp permissions
digit - 4 - oth permissions

4 - read
2 - write
1 - exec

0 - no special permissions
6 - 4 + 2 permission for owner
4 - read permission for grp
4 - read for oth
*/

package main

import (
	"fmt"
	
	"os"
)

func main(){
	// grp := os.Getgid()
	// uid := os.Getuid()
	// pid := os.Getpid()
	// // wd := os.Getwd()
	//  fmt.Println(grp,uid,pid)

	ctrmk := os.Mkdir("test",0755)
	fmt.Println("created",ctrmk)

	file,_ := os.Create("test/tests.txt")
	fmt.Println("created",file)
	
}
