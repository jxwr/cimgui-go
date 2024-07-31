package imgui

/*
#cgo CFLAGS: -I. -Wno-visibility
#include "rlImGui.h"
*/
import "C"

func SetupRaylib() {
	C.rlImGuiSetup(true)
}

func BeginRaylibFrame() {
	C.rlImGuiBegin()
}

func EndRaylibFrame() {
	C.rlImGuiEnd()
}

func ShutdownRaylib() {
	C.rlImGuiShutdown()
}

func ReloadFonts() {
	C.rlImGuiReloadFonts()
}
