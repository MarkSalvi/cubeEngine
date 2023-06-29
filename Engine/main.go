package main

import (
	"fmt"
	"github.com/MarkSalvi/glHelper"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"os"
	"runtime"
	"time"
	"unsafe"
)

const (
	windWidth  = 1920
	windHeight = 1080
)

var vertices = []float32{
	// positions          // normals           // texture coords
	-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,
	0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 0.0,
	0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
	0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
	-0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 1.0,
	-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,

	-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,
	0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 0.0,
	0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
	0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
	-0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 1.0,
	-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,

	-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,
	-0.5, 0.5, -0.5, -1.0, 0.0, 0.0, 1.0, 1.0,
	-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
	-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
	-0.5, -0.5, 0.5, -1.0, 0.0, 0.0, 0.0, 0.0,
	-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,

	0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,
	0.5, 0.5, -0.5, 1.0, 0.0, 0.0, 1.0, 1.0,
	0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
	0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
	0.5, -0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 0.0,
	0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,

	-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,
	0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 1.0, 1.0,
	0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
	0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
	-0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 0.0, 0.0,
	-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,

	-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
	0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 1.0, 1.0,
	0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
	0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
	-0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
	-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
}

func init() {
	runtime.LockOSThread()
}

func main() {

	cubePositions := []mgl32.Vec3{
		mgl32.Vec3{0.0, 0.0, 0.0},
		mgl32.Vec3{2.0, 5.0, -5.0},
		mgl32.Vec3{-1.5, -2.2, -3.0},
		mgl32.Vec3{-3.8, -2.0, -12.3},
		mgl32.Vec3{2.4, -0.4, -3.5},
		mgl32.Vec3{-1.7, 3.0, -7.5},
		mgl32.Vec3{1.3, -2.0, -2.5},
		mgl32.Vec3{1.5, 2.0, -2.5},
		mgl32.Vec3{1.5, 0.2, -1.5},
		mgl32.Vec3{-1.3, 1.0, -1.5},
	}

	lightCubePositions := []mgl32.Vec3{
		mgl32.Vec3{2.0, 2.0, 2.0},
		mgl32.Vec3{5.0, -3.0, -10.0},
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

	lightProgram, err := glHelper.NewShader("shaders\\lightvertex.glsl", "shaders\\lightFrag.frag")
	if err != nil {
		panic(err)
	}

	diffuseMap := glHelper.LoadTexture("assets\\container2.png")
	specularMap := glHelper.LoadTexture("assets\\container2_specular.png")

	glHelper.GenBindBuffer(gl.ARRAY_BUFFER)
	cubeVAO := glHelper.GenBindVertexArray()

	glHelper.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.Ptr(3*unsafe.Sizeof(float32(0))))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.Ptr(6*unsafe.Sizeof(float32(0))))
	gl.EnableVertexAttribArray(2)

	lightVAO := glHelper.GenBindVertexArray()

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(0)
	glHelper.BindVertexArray(0)

	keyboardState := sdl.GetKeyboardState()
	position := mgl32.Vec3{0.0, 0.0, 3.0}
	worldUp := mgl32.Vec3{0.0, 1.0, 0.0}

	camera := glHelper.NewCamera(position, worldUp, -90.0, 0.0, 0.01, 0.5)

	sdl.SetRelativeMouseMode(true)
	sdl.WarpMouseGlobal((windWidth+20)/2, (windHeight+20)/2)
	gl.Enable(gl.DEPTH_TEST)

	var tempoCubo float64
	inizio := time.Now()

	shaderProgram.Use()
	shaderProgram.SetInt("material.diffuse", 0)
	shaderProgram.SetInt("material.specular", 1)

	elapsedTime := float64(0)
	fpsCounter := 0
	timeCounter := time.Now()

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
			}
		}

		dir := keyStrokes(keyboardState)
		camera.UpdateCamera(dir, float32(elapsedTime), float32(mouseX), float32(mouseY))

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shaderProgram.Use()
		//initializing all the uniforms
		shaderProgram.SetVec3("viewPos", camera.GetCameraPosition())
		//setting material properties
		shaderProgram.SetFloat("material.shininess", 64)

		//setting light properties

		//directional light
		shaderProgram.SetVec3("dirLight.direction", mgl32.Vec3{-0.2, -1.0, -0.3})
		shaderProgram.SetVec3("dirLight.ambient", mgl32.Vec3{0.05, 0.05, 0.05})
		shaderProgram.SetVec3("dirLight.diffuse", mgl32.Vec3{0.4, 0.4, 0.4})
		shaderProgram.SetVec3("dirLight.specular", mgl32.Vec3{0.5, 0.5, 0.5})

		//point light1
		pointLightColor1 := mgl32.Vec3{
			(float32(math.Cos(tempoCubo)) * 0.7) + 0.2,
			(float32(math.Cos(tempoCubo)) * 2.0) + 0.2,
			(float32(math.Cos(tempoCubo)) * 1.3) + 0.2,
		}

		shaderProgram.SetVec3("pointLights[0].position", lightCubePositions[0])
		shaderProgram.SetVec3("pointLights[0].ambient", glHelper.MulVec3(pointLightColor1, mgl32.Vec3{0.05, 0.05, 0.05}))
		shaderProgram.SetVec3("pointLights[0].diffuse", glHelper.MulVec3(pointLightColor1, mgl32.Vec3{0.8, 0.8, 0.8}))
		shaderProgram.SetVec3("pointLights[0].specular", mgl32.Vec3{1, 1, 1})
		shaderProgram.SetFloat("pointLights[0].constant", 1)
		shaderProgram.SetFloat("pointLights[0].linear", 0.027)
		shaderProgram.SetFloat("pointLights[0].quadratic", 0.0028)

		//point light 2
		pointLightColor2 := mgl32.Vec3{
			(float32(math.Sin(tempoCubo)) * 1.3) + 0.2,
			(float32(math.Sin(tempoCubo)) * 0.3) + 0.2,
			(float32(math.Sin(tempoCubo)) * 2) + 0.2,
		}

		shaderProgram.SetVec3("pointLights[1].position", lightCubePositions[1])
		shaderProgram.SetVec3("pointLights[1].ambient", glHelper.MulVec3(pointLightColor2, mgl32.Vec3{0.05, 0.05, 0.05}))
		shaderProgram.SetVec3("pointLights[1].diffuse", glHelper.MulVec3(pointLightColor2, mgl32.Vec3{0.8, 0.8, 0.8}))
		shaderProgram.SetVec3("pointLights[1].specular", mgl32.Vec3{1, 1, 1})
		shaderProgram.SetFloat("pointLights[1].constant", 1)
		shaderProgram.SetFloat("pointLights[1].linear", 0.027)
		shaderProgram.SetFloat("pointLights[1].quadratic", 0.0028)

		//bind diffuse map
		gl.ActiveTexture(gl.TEXTURE0)
		glHelper.BindTexture(diffuseMap)
		//bind specular map
		gl.ActiveTexture(gl.TEXTURE1)
		glHelper.BindTexture(specularMap)

		//setting view/projection matrix
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

		tempoCubo = time.Since(inizio).Seconds()
		glHelper.BindVertexArray(lightVAO)
		for i := 0; i < 2; i++ {
			lightModel := mgl32.Ident4()
			lightModel = mgl32.Translate3D(lightCubePositions[i].X(), lightCubePositions[i].Y(), lightCubePositions[i].Z()).Mul4(lightModel)
			lightModel.Mul4x1(mgl32.Vec4{0.2, 0.2, 0.2, 1.0})
			lightProgram.SetMat4("model", lightModel)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
			lightCubePositions[i] = mgl32.Vec3{
				float32(math.Cos(tempoCubo)) * 5,
				float32(math.Sin(tempoCubo)) * -3,
				lightCubePositions[i].Z(),
			}

		}

		window.GLSwap()
		shaderProgram.CheckShaderForChanges()
		lightProgram.CheckShaderForChanges()

		//framerate stuff
		elapsedTime = time.Since(frameStart).Seconds() * 1000
		fpsCounter++
		if time.Since(timeCounter).Seconds() > 1 {
			fmt.Println("FPS: ", fpsCounter)
			fpsCounter = 0
			timeCounter = time.Now()
		}

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
	} else if keyboardState[sdl.SCANCODE_F1] != 0 {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	} else if keyboardState[sdl.SCANCODE_F2] != 0 {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}

	return dir

}
