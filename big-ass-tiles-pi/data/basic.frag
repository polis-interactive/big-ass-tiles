
#ifdef GL_ES

precision highp float;
#define IN varying
#define OUT out
#define TEXTURE texture2D

#else

#define IN in
#define OUT out
#define TEXTURE texture

#endif

uniform float time;
uniform vec2 resolution;

uniform float Brightness;
uniform float Attack;
uniform float Speed;
uniform float Decay;

void main(void)
{

    float t = time * 2.0 * Speed;

    vec2 uv = gl_FragCoord.xy / resolution.xy;

    vec2 uv_grid = floor(vec2(uv.x * 11.0, uv.y * 3.0)) / vec2(11.0, 3.0);

    vec3 color = vec3(sin(t + 2.0 * uv_grid.x), uv_grid.y, sin(t));

    gl_FragColor = vec4(color, 1.0);
}