$(document).ready(function() {
  var loader = new THREE.FontLoader();
	loader.load( 'fonts/helvetiker_regular.typeface.json', function ( response ) {
		font = response
    init()
    cubes()
    render()
	});
})

var data = "this is a test".split("")
var scene, camera, renderer, root;
var mouseX = 0, mouseY = 0;

var font;

var windowHalfX = window.innerWidth / 2;
var windowHalfY = window.innerHeight / 2;

function init() {
  scene = new THREE.Scene();
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
  camera.position.z = 800;
  // camera.position.y = 500

  renderer = new THREE.WebGLRenderer();
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);

  document.addEventListener( 'mousemove', onDocumentMouseMove, false );
}


function onDocumentMouseMove(event) {
	mouseX = (event.clientX - windowHalfX) * 10;
	mouseY = (event.clientY - windowHalfY) * 10;
}


function render() {
  // camera.position.x += ( mouseX - camera.position.x ) * .05;
	// camera.position.y += ( - mouseY - camera.position.y ) * .05;
  root.rotation.x += 0.1
  camera.lookAt( scene.position );

  requestAnimationFrame(render);
  renderer.localClippingEnabled = false;
  renderer.render(scene, camera);
}

function text(val, root, x, y) {
  var textGeo = new THREE.TextGeometry( val, {
		font: font,
		size: 20,
		height: 1,
		curveSegments: 2
	});

  textGeo.translate(-5, -5, 0)
  textGeo.computeBoundingBox();
	textGeo.computeVertexNormals();


  var material = new THREE.MeshBasicMaterial({color:0x00ff00})

  var centerOffset = -0.5 * ( textGeo.boundingBox.max.x - textGeo.boundingBox.min.x );
	textMesh1 = new THREE.Mesh( textGeo, material );
	textMesh1.position.x = x;
	textMesh1.position.y = y;
  root.add(textMesh1);
}

function cubes() {
  var size = 80
  var geometry = new THREE.TorusGeometry(size, 2.5, 6, 0)
  var material = new THREE.MeshBasicMaterial({color:0x00ff00, wireframe: true})

  root = new THREE.Mesh( geometry, material);
	root.position.x = 10;

  var hs = size/2

  text("a", root, hs, 0)
  text("b", root, hs/2, hs*-1)
  text("c", root, (hs/2)*-1, hs*-1)
  text("d", root, hs*-1, 0)
  text("e", root, (hs/2)*-1, hs)
  text("f", root, hs/2, hs)

	scene.add( root );

  var amount = 200, object, parent = root;

  for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.x = 150;
    object.rotation.x = 25
		parent.add(object);
		parent = object;

    text("a", object, hs, 0)
    text("b", object, hs/2, hs*-1)
    text("c", object, (hs/2)*-1, hs*-1)
    text("d", object, hs*-1, 0)
    text("e", object, (hs/2)*-1, hs)
    text("f", object, hs/2, hs)
	}

  parent = root

  for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.x = -150;
    object.rotation.x = -25
		parent.add(object);
		parent = object;

    text("a", object, hs, 0)
    text("b", object, hs/2, hs*-1)
    text("c", object, (hs/2)*-1, hs*-1)
    text("d", object, hs*-1, 0)
    text("e", object, (hs/2)*-1, hs)
    text("f", object, hs/2, hs)
	}

}
