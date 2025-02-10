package main

import (
	_ "github.com/blue-monads/turnix-extern/module"

	"github.com/blue-monads/turnix/backend/distro"

	"github.com/k0kubun/pp"
)

func main() {
	//
	pp.Println("Start turnix-extern")
	dapp, err := distro.NewApp()
	if err != nil {
		pp.Println(err)
		return
	}

	mig, err := dapp.NeedsMigrate()
	if err != nil {
		panic(err)
	}

	if mig {
		err = dapp.RunNormalSeed()
		if err != nil {
			panic(err)
		}
	}

	err = dapp.Start()
	if err != nil {
		pp.Println(err)
		return
	}

	pp.Println("Turnix-extern exited")

}
