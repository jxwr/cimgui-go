// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "cimnodes_wrapper.h"
import "C"
import "unsafe"

type EmulateThreeButtonMouse struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (data EmulateThreeButtonMouse) handle() (result *C.EmulateThreeButtonMouse, releaseFn func()) {
	result = (*C.EmulateThreeButtonMouse)(data.data)
	return result, func() {}
}

func (data EmulateThreeButtonMouse) c() (result C.EmulateThreeButtonMouse, fin func()) {
	resultPtr, finFn := data.handle()
	return *resultPtr, finFn
}

func newEmulateThreeButtonMouseFromC(cvalue *C.EmulateThreeButtonMouse) EmulateThreeButtonMouse {
	result := new(EmulateThreeButtonMouse)
	result.data = unsafe.Pointer(cvalue)
	return *result
}

type NodesIO struct {
	FieldEmulateThreeButtonMouse     EmulateThreeButtonMouse
	FieldLinkDetachWithModifierClick LinkDetachWithModifierClick
	FieldMultipleSelectModifier      MultipleSelectModifier
	FieldAltMouseButton              int32
	FieldAutoPanningSpeed            float32
}

func (data NodesIO) handle() (result *C.ImNodesIO, releaseFn func()) {
	result = new(C.ImNodesIO)
	FieldEmulateThreeButtonMouse := data.FieldEmulateThreeButtonMouse
	EmulateThreeButtonMouseArg, EmulateThreeButtonMouseFin := FieldEmulateThreeButtonMouse.c()
	result.EmulateThreeButtonMouse = EmulateThreeButtonMouseArg
	FieldLinkDetachWithModifierClick := data.FieldLinkDetachWithModifierClick
	LinkDetachWithModifierClickArg, LinkDetachWithModifierClickFin := FieldLinkDetachWithModifierClick.c()
	result.LinkDetachWithModifierClick = LinkDetachWithModifierClickArg
	FieldMultipleSelectModifier := data.FieldMultipleSelectModifier
	MultipleSelectModifierArg, MultipleSelectModifierFin := FieldMultipleSelectModifier.c()
	result.MultipleSelectModifier = MultipleSelectModifierArg
	FieldAltMouseButton := data.FieldAltMouseButton

	result.AltMouseButton = C.int(FieldAltMouseButton)
	FieldAutoPanningSpeed := data.FieldAutoPanningSpeed

	result.AutoPanningSpeed = C.float(FieldAutoPanningSpeed)
	releaseFn = func() {
		EmulateThreeButtonMouseFin()
		LinkDetachWithModifierClickFin()
		MultipleSelectModifierFin()
	}
	return result, releaseFn
}

func (data NodesIO) c() (result C.ImNodesIO, fin func()) {
	resultPtr, finFn := data.handle()
	return *resultPtr, finFn
}

func newNodesIOFromC(cvalue *C.ImNodesIO) NodesIO {
	result := new(NodesIO)
	result.FieldEmulateThreeButtonMouse = newEmulateThreeButtonMouseFromC(&cvalue.EmulateThreeButtonMouse)
	result.FieldLinkDetachWithModifierClick = newLinkDetachWithModifierClickFromC(&cvalue.LinkDetachWithModifierClick)
	result.FieldMultipleSelectModifier = newMultipleSelectModifierFromC(&cvalue.MultipleSelectModifier)
	result.FieldAltMouseButton = int32(cvalue.AltMouseButton)
	result.FieldAutoPanningSpeed = float32(cvalue.AutoPanningSpeed)
	return *result
}

type NodesStyle struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (data NodesStyle) handle() (result *C.ImNodesStyle, releaseFn func()) {
	result = (*C.ImNodesStyle)(data.data)
	return result, func() {}
}

func (data NodesStyle) c() (result C.ImNodesStyle, fin func()) {
	resultPtr, finFn := data.handle()
	return *resultPtr, finFn
}

func newNodesStyleFromC(cvalue *C.ImNodesStyle) NodesStyle {
	result := new(NodesStyle)
	result.data = unsafe.Pointer(cvalue)
	return *result
}

type LinkDetachWithModifierClick struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (data LinkDetachWithModifierClick) handle() (result *C.LinkDetachWithModifierClick, releaseFn func()) {
	result = (*C.LinkDetachWithModifierClick)(data.data)
	return result, func() {}
}

func (data LinkDetachWithModifierClick) c() (result C.LinkDetachWithModifierClick, fin func()) {
	resultPtr, finFn := data.handle()
	return *resultPtr, finFn
}

func newLinkDetachWithModifierClickFromC(cvalue *C.LinkDetachWithModifierClick) LinkDetachWithModifierClick {
	result := new(LinkDetachWithModifierClick)
	result.data = unsafe.Pointer(cvalue)
	return *result
}

type MultipleSelectModifier struct {
	// TODO: contains unsupported fields
	data unsafe.Pointer
}

func (data MultipleSelectModifier) handle() (result *C.MultipleSelectModifier, releaseFn func()) {
	result = (*C.MultipleSelectModifier)(data.data)
	return result, func() {}
}

func (data MultipleSelectModifier) c() (result C.MultipleSelectModifier, fin func()) {
	resultPtr, finFn := data.handle()
	return *resultPtr, finFn
}

func newMultipleSelectModifierFromC(cvalue *C.MultipleSelectModifier) MultipleSelectModifier {
	result := new(MultipleSelectModifier)
	result.data = unsafe.Pointer(cvalue)
	return *result
}
