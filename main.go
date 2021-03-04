package main

import (
	"github.com/confetti-framework/contract/inter"
	"os"
	"src/app/bootstrap"
)

func main() {
	/*
	   |--------------------------------------------------------------------------
	   | Turn On The Lights
	   |--------------------------------------------------------------------------
	   |
	   |
	   |
	*/
	app := bootstrap.NewAppFromBoot()

	/*
	   |--------------------------------------------------------------------------
	   | Run The Application
	   |--------------------------------------------------------------------------
	   |
	   | Once we have the application, we can handle the command
	   | through the kernel.
	   |
	*/
	kernel := app.Make((*inter.ConsoleKernel)(nil)).(inter.ConsoleKernel)
	code := kernel.Handle()
	os.Exit(int(code))
}
