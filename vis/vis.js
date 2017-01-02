$(document).ready(function() {
  init()
  cube()
  render()
})

var data = "this is a test".split("")
var scene, camera, renderer;

function init() {
  scene = new THREE.Scene();
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);

  renderer = new THREE.WebGLRenderer();
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);
}

function render() {
  requestAnimationFrame(render);
  renderer.render(scene, camera);
}

function cube() {
  var geometry = new THREE.BoxGeometry(1,1,1)
  var material = new THREE.MeshBasicMaterial({color:0x00ff00})
  var cube = new THREE.Mesh(geometry, material)
  scene.add(cube)

  camera.position.z = 5;
}
