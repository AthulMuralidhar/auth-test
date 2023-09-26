package main

import (
	"fmt"
	"github.com/mikespook/gorbac"
)

func main() {
	rbac := gorbac.New()

	//rA := gorbac.NewRole("role-a")

	roleA := gorbac.NewStdRole("role-a")

	//pA := gorbac.NewPermission("permission-a")  // seems like things are renamed
	pA := gorbac.NewStdPermission("permission-a")

	roleA.Assign(pA)

	// init happens here
	rbac.Add(roleA)

	// checking
	if rbac.IsGranted("role-a", pA, nil) {
		fmt.Println("The role-a has been granted permissions a ")
	}
}
