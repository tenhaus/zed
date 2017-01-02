$(document).ready(function() {
  var loader = new THREE.FontLoader();
	loader.load( 'fonts/helvetiker_regular.typeface.json', function ( response ) {
		font = response
    init()
    cubes()
    text()
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
  camera.position.z = 500;

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

  // camera.lookAt( scene.position );

  requestAnimationFrame(render);
  renderer.localClippingEnabled = false;
  globalPlane = new THREE.Plane( new THREE.Vector3( - 1, 0, 0 ), 0.1 );
  var globalPlanes = [ globalPlane ],
					Empty = Object.freeze( [] );
				renderer.clippingPlanes = Empty;
  renderer.render(scene, camera);
}

function text() {
  var textGeo = new THREE.TextGeometry( "heeeey", {
		font: font,
		size: 80,
		height: 20,
		curveSegments: 2
	});

  textGeo.computeBoundingBox();
	textGeo.computeVertexNormals();

  var material = new THREE.MeshBasicMaterial({color:0x00ffff})

  var centerOffset = -0.5 * ( textGeo.boundingBox.max.x - textGeo.boundingBox.min.x );
	textMesh1 = new THREE.Mesh( textGeo, material );
	textMesh1.position.x = 150;
	textMesh1.position.y = 150;
	textMesh1.position.z = 0;
	textMesh1.rotation.x = 0;
	textMesh1.rotation.y = Math.PI * 2;
	scene.add( textMesh1 );
}

function cubes() {
  var geometry = new THREE.BoxGeometry(100,100,100)
  var material = new THREE.MeshBasicMaterial({color:0x00ff00})

  root = new THREE.Mesh( geometry, material );
	root.position.x = 10;
	scene.add( root );

  var amount = 200, object, parent = root;

  for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.x = 150;
		parent.add(object);
		parent = object;
	}

  parent = root;

  for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.x = -150;
		parent.add(object);
		parent = object;
	}

  parent = root;

	for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.y = -150;
		parent.add(object);
		parent = object;
	}

  parent = root;

	for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.y = 150;
		parent.add(object);
		parent = object;
	}

  parent = root;

	for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.z = -150;
		parent.add(object);
		parent = object;
	}

  parent = root;

	for(var i = 0; i < amount; i ++) {
		object = new THREE.Mesh(geometry, material);
		object.position.z = 150;
		parent.add(object);
		parent = object;
	}

}
