//go:build required
// +build required

package imgui

import (
	_ "github.com/jxwr/cimgui-go/lib/linux/x64"
	_ "github.com/jxwr/cimgui-go/lib/macos/arm64"
	_ "github.com/jxwr/cimgui-go/lib/macos/x64"
	_ "github.com/jxwr/cimgui-go/lib/windows/x64"
)
