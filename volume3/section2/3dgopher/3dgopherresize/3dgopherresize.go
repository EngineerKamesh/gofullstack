package main

import (
	"math"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var JS = js.Global
var D = dom.GetWindow().Document()

var container *dom.HTMLDivElement
var camera *js.Object
var cameraTarget *js.Object
var scene *js.Object
var renderer *js.Object
var gopher *js.Object

type MeshPhongMaterialData struct {
	*js.Object
	Color     int `js:"color"`
	Specular  int `js:"specular"`
	Shininess int `js:"shininess"`
}

func main() {
	println("Hello 3D Gopher!")

	dom.GetWindow().AddEventListener("resize", false, func(event dom.Event) {
		handleWindowResize()
	})

	initialize()
	animate()
}

func initialize() {

	container = D.CreateElement("div").(*dom.HTMLDivElement)
	D.QuerySelector("body").AppendChild(container)

	camera = JS.Get("THREE").Get("PerspectiveCamera").New(39, JS.Get("window").Get("innerWidth").Int()/JS.Get("window").Get("innerHeight").Int(), 1, 15)
	camera.Get("position").Call("set", 6, 0.54, 6)
	cameraTarget = JS.Get("THREE").Get("Vector3").New(0, 0, 0)
	scene = JS.Get("THREE").Get("Scene").New()

	fog := JS.Get("THREE").Get("Fog").New(0x708090, 2, 15)
	scene.Set("fog", fog)

	planeMeshData := &MeshPhongMaterialData{Object: js.Global.Get("Object").New()}
	planeMeshData.Color = 0xA9A9A9
	planeMeshData.Specular = 0x000000

	planeMeshPhongMaterial := JS.Get("THREE").Get("MeshPhongMaterial").New(planeMeshData)
	planeBufferGeometry := JS.Get("THREE").Get("PlaneBufferGeometry").New(40, 40)
	plane := JS.Get("THREE").Get("Mesh").New(planeBufferGeometry, planeMeshPhongMaterial)

	plane.Get("rotation").Set("x", -1*(math.Pi/2))
	plane.Get("position").Set("y", -0.5)

	scene.Call("add", plane)

	loader := JS.Get("THREE").Get("OBJLoader").New()
	loader.Call("load", "/obj/gogopher.obj", func(mesh *js.Object) {

		mesh.Get("position").Call("set", 0.06, -0.45, 0.7)
		mesh.Get("scale").Call("set", 0.004, 0.0054, 0.004)
		mesh.Set("castsShadow", true)
		mesh.Set("receiveShadow", true)

		scene.Call("add", mesh)

	})

	hemisphereLight := JS.Get("THREE").Get("HemisphereLight").New(0xf0f0f0, 0x000000)
	scene.Call("add", hemisphereLight)
	addDirectionalLight(-0.5, 3, -1.0, 0xf0f0f0, 1)
	renderer = JS.Get("THREE").Get("WebGLRenderer").New()
	renderer.Call("setClearColor", scene.Get("fog").Get("color"))
	renderer.Call("setPixelRatio", JS.Get("window").Get("devicePixelRatio"))
	renderer.Call("setSize", JS.Get("window").Get("innerWidth").Int(), JS.Get("window").Get("innerHeight").Int())
	renderer.Set("gammaInput", true)
	renderer.Set("gammaOutput", true)
	renderer.Get("shadowMap").Set("enabled", true)
	renderer.Get("shadowMap").Set("renderReverseSided", false)
	container.Underlying().Call("appendChild", renderer.Get("domElement"))
}

func animate() {
	JS.Get("window").Call("requestAnimationFrame", animate)
	render()
}

func render() {
	camera.Call("lookAt", cameraTarget)
	renderer.Call("render", scene, camera)
}

func addDirectionalLight(x float64, y float64, z float64, color int, intensity int) {
	directionalLight := JS.Get("THREE").Get("DirectionalLight").New(color, intensity)
	directionalLight.Get("position").Call("set", x, y, z)
	scene.Call("add", directionalLight)
	directionalLight.Set("castShadow", true)
	d := 1
	directionalLight.Get("shadow").Get("camera").Set("left", -d)
	directionalLight.Get("shadow").Get("camera").Set("right", d)
	directionalLight.Get("shadow").Get("camera").Set("top", d)
	directionalLight.Get("shadow").Get("camera").Set("bottom", -d)
	directionalLight.Get("shadow").Get("camera").Set("near", 1)
	directionalLight.Get("shadow").Get("camera").Set("far", 4)
	directionalLight.Get("shadow").Get("mapSize").Set("width", 1024)
	directionalLight.Get("shadow").Get("mapSize").Set("height", 1024)
	directionalLight.Get("shadow").Set("bias", -0.005)
}

func handleWindowResize() {
	camera.Set("aspect", JS.Get("window").Get("innerWidth").Int()/JS.Get("window").Get("innerHeight").Int())
	camera.Call("updateProjectionMatrix")
	renderer.Call("setSize", JS.Get("window").Get("innerWidth").Int(), JS.Get("window").Get("innerHeight").Int())
}
