package main

import (
	"fmt"
	"github.com/MarkSalvi/glHelper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
)

const (
	windWidth  = 1920
	windHeight = 1080
)

func init() {
	runtime.LockOSThread()
}

func main() {

	vertices := []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0}

	cubePositions := []mgl32.Vec3{
		mgl32.Vec3{0.0, 0.0, 0.0},
		mgl32.Vec3{2.0, 5.0, -15.0},
		mgl32.Vec3{-1.5, -2.2, -3.0},
		//mgl32.Vec3{-3.8, -2.0, -12.3},
		//mgl32.Vec3{2.4, -0.4, -3.5},
		//mgl32.Vec3{-1.7, 3.0, -7.5},
		//mgl32.Vec3{1.3, -2.0, -2.5},
		//mgl32.Vec3{1.5, 2.0, -2.5},
		//mgl32.Vec3{1.5, 0.2, -1.5},
		//mgl32.Vec3{-1.3, 1.0, -1.5},
	}

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()
	var flags uint32
	flags = sdl.WINDOW_INPUT_GRABBED | sdl.WINDOW_OPENGL | sdl.WINDOW_FULLSCREEN

	window, err := sdl.CreateWindow("CIAO", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windWidth, windHeight, flags)

	if err != nil {
		panic(err)
	}
	_, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}

	defer window.Destroy()
	gl.Init()
	version := glHelper.GetVersion()
	fmt.Println("OpenGL Version: ", version)

	shaderProgram, err := glHelper.NewShader("shaders\\cubevertex.vert", "shaders\\cubeFrag.frag")
	if err != nil {
		panic(err)
	}

	lightProgram, err := glHelper.NewShader("shaders\\cubevertex.vert", "shaders\\lightFrag.frag")
	if err != nil {
		panic(err)
	}

	//texture := glHelper.LoadTexture("assets\\cobblestone.png")

	glHelper.GenBindBuffer(gl.ARRAY_BUFFER)
	cubeVAO := glHelper.GenBindVertexArray()

	glHelper.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(0)
	//gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(2*4))
	//gl.EnableVertexAttribArray(1)

	lightVAO := glHelper.GenBindVertexArray()

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(0)
	glHelper.BindVertexArray(0)

	keyboardState := sdl.GetKeyboardState()
	position := mgl32.Vec3{0.0, 0.0, 3.0}
	worldUp := mgl32.Vec3{0.0, 1.0, 0.0}

	camera := glHelper.NewCamera(position, worldUp, -90.0, 0.0, 0.01, 0.5)
	elapsedTime := float32(0)
	sdl.SetRelativeMouseMode(true)
	sdl.WarpMouseGlobal((windWidth+20)/2, (windHeight+20)/2)
	gl.Enable(gl.DEPTH_TEST)

	//game Loop
	for {
		frameStart := time.Now()
		var mouseX, mouseY int32

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				sdl.SetRelativeMouseMode(false)
				return
			case *sdl.MouseMotionEvent:
				mouseX, mouseY = t.XRel, t.YRel
				//todo fix max and min pitch camera problem
			}
		}

		dir := keyStrokes(keyboardState)
		camera.UpdateCamera(dir, elapsedTime, float32(mouseX), float32(mouseY))

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shaderProgram.Use()
		shaderProgram.SetVec3("objectColor", mgl32.Vec3{0.5, 0.5, 1.0})
		shaderProgram.SetVec3("lightColor", mgl32.Vec3{1.0, 1.0, 1.0})
		projectionMatrix := mgl32.Perspective(mgl32.DegToRad(90.0), float32(windWidth)/float32(windHeight), 0.1, 100.0)
		shaderProgram.SetMat4("view", camera.GetViewMatrix())
		shaderProgram.SetMat4("projection", projectionMatrix)

		//glHelper.BindTexture(texture)
		glHelper.BindVertexArray(cubeVAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)

		for _, pos := range cubePositions {

			modelMatrix, err := glHelper.Rotate3DMat4(mgl32.DegToRad(-55.0), mgl32.Vec3{1.0, 0.0, 0.0})
			modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
			if err != nil {
				panic(err)
			}
			shaderProgram.SetMat4("model", modelMatrix)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		lightProgram.Use()
		lightProgram.SetMat4("view", camera.GetViewMatrix())
		lightProgram.SetMat4("projection", projectionMatrix)
		lightModel := mgl32.Ident4()
		lightModel = mgl32.Translate3D(1.0, 1.0, -2.0).Mul4(lightModel)
		lightModel.Mul4x1(mgl32.Vec4{0.2, 0.2, 0.2, 1.0})
		lightProgram.SetMat4("model", lightModel)

		glHelper.BindVertexArray(lightVAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 36)

		window.GLSwap()
		shaderProgram.CheckShaderForChanges()
		lightProgram.CheckShaderForChanges()

		elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)

	}

}

func keyStrokes(keyboardState []uint8) glHelper.Direction {
	dir := glHelper.Nowhere
	if keyboardState[sdl.SCANCODE_A] != 0 {
		dir = glHelper.Left
	} else if keyboardState[sdl.SCANCODE_D] != 0 {
		dir = glHelper.Right
	} else if keyboardState[sdl.SCANCODE_W] != 0 {
		dir = glHelper.Forward
	} else if keyboardState[sdl.SCANCODE_S] != 0 {
		dir = glHelper.Backward
	} else if keyboardState[sdl.SCANCODE_ESCAPE] != 0 {
		sdl.SetRelativeMouseMode(false)
		os.Exit(0)
	}
	return dir

}

func mouseAndKeyboard() {

}
