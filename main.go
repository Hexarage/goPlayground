package main

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v4"

	"github.com/inkyblackness/imgui-go-examples"
	"github.com/inkyblackness/imgui-go-examples/internal/example"
	"github.com/inkyblackness/imgui-go-examples/internal/platforms"
	"github.com/inkyblackness/imgui-go-examples/internal/renderers"
)

func main() {
	fmt.Println("Howdy")

	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	example.Run(platform, renderer)
}
