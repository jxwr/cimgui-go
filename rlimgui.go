package imgui

/*
#cgo CFLAGS: -I. -Wno-visibility
#include "rlImGui.h"
*/
import "C"

func ImGuiSetupRL() {
	C.rlImGuiSetup(true)
}

func ImGuiBeginRL() {
	C.rlImGuiBegin()
}

func ImGuiEndRL() {
	C.rlImGuiEnd()
}

func ImGuiShutdownRL() {
	C.rlImGuiShutdown()
}
