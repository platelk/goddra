package shaders


const TriangleVertexShaderSource = `#version 300 es
 
// an attribute is an input (in) to a vertex shader.
// It will receive data from a buffer
in vec2 a_position;
 
// Used to pass in the resolution of the canvas
uniform vec2 u_resolution;
 
// translation to add to position
uniform vec2 u_translation;
 
// all shaders have a main function
void main() {
  // Add in the translation
  vec2 position = a_position + u_translation;
 
  // convert the position from pixels to 0.0 to 1.0
  vec2 zeroToOne = position / u_resolution;
 
  // convert from 0->1 to 0->2
  vec2 zeroToTwo = zeroToOne * 2.0;
 
  // convert from 0->2 to -1->+1 (clip space)
  vec2 clipSpace = zeroToTwo - 1.0;
 
  gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
}
`

const TriangleFragmentShaderSource = `#version 300 es

precision mediump float;

uniform vec4 u_color;

// we need to declare an output for the fragment shader
out vec4 outColor;

void main() {
  outColor = u_color;
}
`
