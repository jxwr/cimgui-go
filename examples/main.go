package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	imgui "github.com/jxwr/cimgui-go"
)

var (
	showDemoWindow bool
	value1         int32
	value2         int32
	value3         int32
	values         [2]int32 = [2]int32{value1, value2}
	content        string   = "Let me try"
	r              float32
	g              float32
	b              float32
	a              float32
	color4         [4]float32 = [4]float32{r, g, b, a}
	selected       bool
)

func callback(data imgui.InputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func showWidgetsDemo() {
	if showDemoWindow {
		imgui.ShowDemoWindowV(&showDemoWindow)
	}

	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
	imgui.Begin("Window 1")
	if imgui.ButtonV("Click Me", imgui.NewVec2(80, 20)) {
		fmt.Println("Click Me")
	}
	imgui.TextUnformatted("Unformatted text")
	imgui.Checkbox("Show demo window", &showDemoWindow)
	if imgui.BeginCombo("Combo", "Combo preview") {
		imgui.SelectableBoolPtr("Item 1", &selected)
		imgui.SelectableBool("Item 2")
		imgui.SelectableBool("Item 3")
		imgui.EndCombo()
	}

	if imgui.RadioButtonBool("Radio button1", selected) {
		selected = true
	}

	imgui.SameLine()

	if imgui.RadioButtonBool("Radio button2", !selected) {
		selected = false
	}

	imgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	imgui.Text(content)
	imgui.SliderInt("Slider int", &value3, 0, 100)
	imgui.DragInt("Drag int", &value1)
	imgui.DragInt2("Drag int2", &values)
	value1 = values[0]
	imgui.ColorEdit4("Color Edit3", &color4)
	imgui.End()
}

func main() {
	screenWidth := 1200
	screenHeight := 800

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Raylib Window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	imgui.SetupRaylib()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		imgui.BeginRaylibFrame()
		rl.ClearBackground(rl.RayWhite)

		imgui.Begin("panel")
		if imgui.Button("button") {
			fmt.Printf("button pressed\n")
		}
		imgui.End()

		showWidgetsDemo()

		rl.DrawCircle(100, 100, 40, rl.Blue)

		imgui.EndRaylibFrame()
		rl.EndDrawing()
	}

	imgui.ShutdownRaylib()
}
