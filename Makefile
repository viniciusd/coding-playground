build:
	emcc hello.c -s WASM=1 -s EXPORTED_FUNCTIONS="['_hello']" -o hello.js
