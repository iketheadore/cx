package base

// constant codes
const (
	// https://www.khronos.org/registry/OpenGL/api/GL/glcorearb.h
	// gl_1_0
	CONST_GL_DEPTH_BUFFER_BIT = iota
	CONST_GL_STENCIL_BUFFER_BIT
	CONST_GL_COLOR_BUFFER_BIT
	CONST_GL_FALSE
	CONST_GL_TRUE
	CONST_GL_POINTS
	CONST_GL_LINES
	CONST_GL_LINE_LOOP
	CONST_GL_LINE_STRIP
	CONST_GL_TRIANGLES
	CONST_GL_TRIANGLE_STRIP
	CONST_GL_TRIANGLE_FAN
	CONST_GL_QUADS
	CONST_GL_NEVER
	CONST_GL_LESS
	CONST_GL_EQUAL
	CONST_GL_LEQUAL
	CONST_GL_GREATER
	CONST_GL_NOTEQUAL
	CONST_GL_GEQUAL
	CONST_GL_ALWAYS
	CONST_GL_ZERO
	CONST_GL_ONE
	CONST_GL_SRC_ALPHA
	CONST_GL_ONE_MINUS_SRC_ALPHA
	CONST_GL_FRONT
	CONST_GL_BACK
	CONST_GL_FRONT_AND_BACK
	CONST_GL_NO_ERROR
	CONST_GL_INVALID_ENUM
	CONST_GL_INVALID_VALUE
	CONST_GL_INVALID_OPERATION
	CONST_GL_STACK_OVERFLOW
	CONST_GL_STACK_UNDERFLOW
	CONST_GL_OUT_OF_MEMORY
	CONST_GL_LINE_SMOOTH
	CONST_GL_POLYGON_SMOOTH
	CONST_GL_CULL_FACE
	CONST_GL_DEPTH_RANGE
	CONST_GL_DEPTH_TEST
	CONST_GL_DEPTH_WRITEMASK
	CONST_GL_DEPTH_CLEAR_VALUE
	CONST_GL_DEPTH_FUNC
	CONST_GL_STENCIL_TEST
	CONST_GL_STENCIL_CLEAR_VALUE
	CONST_GL_STENCIL_FUNC
	CONST_GL_STENCIL_VALUE_MASK
	CONST_GL_STENCIL_FAIL
	CONST_GL_STENCIL_PASS_DEPTH_FAIL
	CONST_GL_STENCIL_PASS_DEPTH_PASS
	CONST_GL_STENCIL_REF
	CONST_GL_STENCIL_WRITEMASK
	CONST_GL_DITHER
	CONST_GL_BLEND
	CONST_GL_SCISSOR_TEST
	CONST_GL_POLYGON_SMOOTH_HINT
	CONST_GL_TEXTURE_2D
	CONST_GL_TEXTURE_WIDTH
	CONST_GL_TEXTURE_HEIGHT
	CONST_GL_DONT_CARE
	CONST_GL_UNSIGNED_BYTE
	CONST_GL_FLOAT
	CONST_GL_INVERT
	CONST_GL_TEXTURE
	CONST_GL_COLOR
	CONST_GL_DEPTH
	CONST_GL_STENCIL
	CONST_GL_STENCIL_INDEX
	CONST_GL_DEPTH_COMPONENT
	CONST_GL_RGBA
	CONST_GL_KEEP
	CONST_GL_REPLACE
	CONST_GL_INCR
	CONST_GL_DECR
	CONST_GL_NEAREST
	CONST_GL_LINEAR
	CONST_GL_NEAREST_MIPMAP_NEAREST
	CONST_GL_LINEAR_MIPMAP_NEAREST
	CONST_GL_NEAREST_MIPMAP_LINEAR
	CONST_GL_LINEAR_MIPMAP_LINEAR
	CONST_GL_TEXTURE_MAG_FILTER
	CONST_GL_TEXTURE_MIN_FILTER
	CONST_GL_TEXTURE_WRAP_S
	CONST_GL_TEXTURE_WRAP_T
	CONST_GL_REPEAT

	// gl_1_1
	CONST_GL_RGBA8
	CONST_GL_VERTEX_ARRAY

	// gl_1_2
	CONST_GL_TEXTURE_WRAP_R
	CONST_GL_CLAMP_TO_EDGE

	// gl_1_3
	CONST_GL_TEXTURE0
	CONST_GL_MULTISAMPLE_ARB // remove _ARB
	CONST_GL_CLAMP_TO_BORDER

	// gl_1_4
	CONST_GL_DEPTH_COMPONENT16
	CONST_GL_DEPTH_COMPONENT24
	CONST_GL_DEPTH_COMPONENT32
	CONST_GL_MIRRORED_REPEAT
	CONST_GL_INCR_WRAP
	CONST_GL_DECR_WRAP

	// gl_1_5
	CONST_GL_ARRAY_BUFFER
	CONST_GL_STREAM_DRAW
	CONST_GL_STREAM_READ
	CONST_GL_STREAM_COPY
	CONST_GL_STATIC_DRAW
	CONST_GL_STATIC_READ
	CONST_GL_STATIC_COPY
	CONST_GL_DYNAMIC_DRAW
	CONST_GL_DYNAMIC_READ
	CONST_GL_DYNAMIC_COPY

	// gl_2_0
	CONST_GL_STENCIL_BACK_FUNC
	CONST_GL_STENCIL_BACK_FAIL
	CONST_GL_STENCIL_BACK_PASS_DEPTH_FAIL
	CONST_GL_STENCIL_BACK_PASS_DEPTH_PASS
	CONST_GL_FRAGMENT_SHADER
	CONST_GL_VERTEX_SHADER
	CONST_GL_STENCIL_BACK_REF
	CONST_GL_STENCIL_BACK_VALUE_MASK
	CONST_GL_STENCIL_BACK_WRITEMASK

	// gl_3_0
	CONST_GL_DEPTH_COMPONENT32F
	CONST_GL_DEPTH32F_STENCIL8
	CONST_GL_FRAMEBUFFER_UNDEFINED
	CONST_GL_DEPTH_STENCIL_ATTACHMENT
	CONST_GL_DEPTH_STENCIL
	CONST_GL_DEPTH24_STENCIL8
	CONST_GL_FRAMEBUFFER_COMPLETE
	CONST_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	CONST_GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	CONST_GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	CONST_GL_FRAMEBUFFER_UNSUPPORTED
	CONST_GL_COLOR_ATTACHMENT0
	CONST_GL_DEPTH_ATTACHMENT
	CONST_GL_STENCIL_ATTACHMENT
	CONST_GL_FRAMEBUFFER
	CONST_GL_RENDERBUFFER
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE
	CONST_GL_STENCIL_INDEX1
	CONST_GL_STENCIL_INDEX4
	CONST_GL_STENCIL_INDEX8
	CONST_GL_STENCIL_INDEX16

	// gl_3_2
	CONST_GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS

	// gl_4_4
	CONST_GL_MIRROR_CLAMP_TO_EDGE

	// Fixed pipeline. Deprecated ?
	CONST_GL_POLYGON
	CONST_GL_MODELVIEW
	CONST_GL_PROJECTION
	CONST_GL_MODELVIEW_MATRIX
	CONST_GL_LIGHTING
	CONST_GL_LIGHT0
	CONST_GL_AMBIENT
	CONST_GL_DIFFUSE
	CONST_GL_POSITION
	CONST_GL_TEXTURE_ENV
	CONST_GL_TEXTURE_ENV_MODE
	CONST_GL_MODULATE
	CONST_GL_DECAL
	CONST_GL_POINT_SMOOTH

	// glfw
	CONST_GLFW_FALSE
	CONST_GLFW_TRUE
	CONST_GLFW_PRESS
	CONST_GLFW_RELEASE
	CONST_GLFW_REPEAT
	CONST_GLFW_KEY_UNKNOWN
	CONST_GLFW_CURSOR
	CONST_GLFW_STICKY_KEYS
	CONST_GLFW_STICKY_MOUSE_BUTTONS
	CONST_GLFW_CURSOR_NORMAL
	CONST_GLFW_CURSOR_HIDDEN
	CONST_GLFW_CURSOR_DISABLED
	CONST_GLFW_RESIZABLE
	CONST_GLFW_CONTEXT_VERSION_MAJOR
	CONST_GLFW_CONTEXT_VERSION_MINOR
	CONST_GLFW_OPENGL_PROFILE
	CONST_GLFW_OPENGL_COREPROFILE
	CONST_GLFW_OPENGL_FORWARD_COMPATIBLE
	CONST_GLFW_MOUSE_BUTTON_LAST
	CONST_GLFW_MOUSE_BUTTON_LEFT
	CONST_GLFW_MOUSE_BUTTON_RIGHT
	CONST_GLFW_MOUSE_BUTTON_MIDDLE

	// gltext
	CONST_GLTEXT_LEFT_TO_RIGHT
	CONST_GLTEXT_RIGHT_TO_LEFT
	CONST_GLTEXT_TOP_TO_BOTTOM

	// os
	CONST_OS_RUN_SUCCESS
	CONST_OS_RUN_EMPTY_CMD
	CONST_OS_RUN_PANIC
	CONST_OS_RUN_START_FAILED
	CONST_OS_RUN_WAIT_FAILED
	CONST_OS_RUN_TIMEOUT

	// cx
	CONST_CX_SUCCESS
	CONST_CX_RUNTIME_ERROR
	CONST_CX_PANIC
	CONST_CX_COMPILATION_ERROR
	CONST_CX_INTERNAL_ERROR
	CONST_CX_ASSERT
	CONST_CX_RUNTIME_INVALID_ARGUMENT
	CONST_CX_RUNTIME_SLICE_INDEX_OUT_OF_RANGE
)

// For the parser. These shouldn't be used in the runtime for performance reasons
var ConstNames map[int]string = map[int]string {}
var ConstCodes map[string]int = map[string]int {}
var Constants map[int]CXConstant = map[int]CXConstant {}

func AddConstCode(code int, name string, typ int, value []byte) {
	ConstNames[code] = name
	ConstCodes[name] = code
	Constants[code] = CXConstant{Type: typ, Value: value}
}

func init() {
	/* gl_1_0 */
	AddConstCode( CONST_GL_DEPTH_BUFFER_BIT       , "gl.DEPTH_BUFFER_BIT"       , TYPE_I32, FromI32(0x00000100))
	AddConstCode( CONST_GL_STENCIL_BUFFER_BIT     , "gl.STENCIL_BUFFER_BIT"     , TYPE_I32, FromI32(0x00000400))
	AddConstCode( CONST_GL_COLOR_BUFFER_BIT       , "gl.COLOR_BUFFER_BIT"       , TYPE_I32, FromI32(0x00004000))
	AddConstCode( CONST_GL_FALSE                  , "gl.FALSE"                  , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GL_TRUE                   , "gl.TRUE"                   , TYPE_I32, FromI32(1))
	AddConstCode( CONST_GL_POINTS                 , "gl.POINTS"                 , TYPE_I32, FromI32(0x0000))
	AddConstCode( CONST_GL_LINES                  , "gl.LINES"                  , TYPE_I32, FromI32(0x0001))
	AddConstCode( CONST_GL_LINE_LOOP              , "gl.LINE_LOOP"              , TYPE_I32, FromI32(0x0002))
	AddConstCode( CONST_GL_LINE_STRIP             , "gl.LINE_STRIP"             , TYPE_I32, FromI32(0x0003))
	AddConstCode( CONST_GL_TRIANGLES              , "gl.TRIANGLES"              , TYPE_I32, FromI32(0x0004))
	AddConstCode( CONST_GL_TRIANGLE_STRIP         , "gl.TRIANGLE_STRIP"         , TYPE_I32, FromI32(0x0005))
	AddConstCode( CONST_GL_TRIANGLE_FAN           , "gl.TRIANGLE_FAN"           , TYPE_I32, FromI32(0x0006))
	AddConstCode( CONST_GL_QUADS                  , "gl.QUADS"                  , TYPE_I32, FromI32(0x0007))
	AddConstCode( CONST_GL_NEVER                  , "gl.NEVER"                  , TYPE_I32, FromI32(0x0200))
	AddConstCode( CONST_GL_LESS                   , "gl.LESS"                   , TYPE_I32, FromI32(0x0201))
	AddConstCode( CONST_GL_EQUAL                  , "gl.EQUAL"                  , TYPE_I32, FromI32(0x0202))
	AddConstCode( CONST_GL_LEQUAL                 , "gl.LEQUAL"                 , TYPE_I32, FromI32(0x0203))
	AddConstCode( CONST_GL_GREATER                , "gl.GREATER"                , TYPE_I32, FromI32(0x0204))
	AddConstCode( CONST_GL_NOTEQUAL               , "gl.NOTEQUAL"               , TYPE_I32, FromI32(0x0205))
	AddConstCode( CONST_GL_GEQUAL                 , "gl.GEQUAL"                 , TYPE_I32, FromI32(0x0206))
	AddConstCode( CONST_GL_ALWAYS                 , "gl.ALWAYS"                 , TYPE_I32, FromI32(0x0207))
	AddConstCode( CONST_GL_ZERO                   , "gl.ZERO"                   , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GL_ONE                    , "gl.ONE"                    , TYPE_I32, FromI32(1))
	AddConstCode( CONST_GL_SRC_ALPHA              , "gl.SRC_ALPHA"              , TYPE_I32, FromI32(0x302))
	AddConstCode( CONST_GL_ONE_MINUS_SRC_ALPHA    , "gl.ONE_MINUS_SRC_ALPHA"    , TYPE_I32, FromI32(0x303))
	AddConstCode( CONST_GL_FRONT                  , "gl.FRONT"                  , TYPE_I32, FromI32(0x404))
	AddConstCode( CONST_GL_BACK                   , "gl.BACK"                   , TYPE_I32, FromI32(0x405))
	AddConstCode( CONST_GL_FRONT_AND_BACK         , "gl.FRONT_AND_BACK"         , TYPE_I32, FromI32(0x408))
	AddConstCode( CONST_GL_NO_ERROR               , "gl.NO_ERROR"               , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GL_INVALID_ENUM           , "gl.INVALID_ENUM"           , TYPE_I32, FromI32(0x500))
	AddConstCode( CONST_GL_INVALID_VALUE          , "gl.INVALID_VALUE"          , TYPE_I32, FromI32(0x501))
	AddConstCode( CONST_GL_INVALID_OPERATION      , "gl.INVALID_OPERATION"      , TYPE_I32, FromI32(0x502))
	AddConstCode( CONST_GL_STACK_OVERFLOW         , "gl.STACK_OVERFLOW"         , TYPE_I32, FromI32(0x503))
	AddConstCode( CONST_GL_STACK_UNDERFLOW        , "gl.STACK_UNDERFLOW"        , TYPE_I32, FromI32(0x504))
	AddConstCode( CONST_GL_OUT_OF_MEMORY          , "gl.OUT_OF_MEMORY"          , TYPE_I32, FromI32(0x505))
	AddConstCode( CONST_GL_LINE_SMOOTH            , "gl.LINE_SMOOTH"            , TYPE_I32, FromI32(0x0B20))
	AddConstCode( CONST_GL_POLYGON_SMOOTH         , "gl.POLYGON_SMOOTH"         , TYPE_I32, FromI32(0x0B41))
	AddConstCode( CONST_GL_CULL_FACE              , "gl.CULL_FACE"              , TYPE_I32, FromI32(0x0B44))
	AddConstCode( CONST_GL_DEPTH_RANGE            , "gl.DEPTH_RANGE"            , TYPE_I32, FromI32(0x0B70))
	AddConstCode( CONST_GL_DEPTH_TEST             , "gl.DEPTH_TEST"             , TYPE_I32, FromI32(0x0B71))
	AddConstCode( CONST_GL_DEPTH_WRITEMASK        , "gl.DEPTH_WRITEMASK"        , TYPE_I32, FromI32(0x0B72))
	AddConstCode( CONST_GL_DEPTH_CLEAR_VALUE      , "gl.DEPTH_CLEAR_VALUE"      , TYPE_I32, FromI32(0x0B73))
	AddConstCode( CONST_GL_DEPTH_FUNC             , "gl.DEPTH_FUNC"             , TYPE_I32, FromI32(0x0B74))
	AddConstCode( CONST_GL_STENCIL_TEST           , "gl.STENCIL_TEST"           , TYPE_I32, FromI32(0x0B90))
	AddConstCode( CONST_GL_STENCIL_CLEAR_VALUE    , "gl.STENCIL_CLEAR_VALUE"    , TYPE_I32, FromI32(0x0B91))
	AddConstCode( CONST_GL_STENCIL_FUNC           , "gl.STENCIL_FUNC"           , TYPE_I32, FromI32(0x0B92))
	AddConstCode( CONST_GL_STENCIL_VALUE_MASK     , "gl.STENCIL_VALUE_MASK"     , TYPE_I32, FromI32(0x0B93))
	AddConstCode( CONST_GL_STENCIL_FAIL           , "gl.STENCIL_FAIL"           , TYPE_I32, FromI32(0x0B94))
	AddConstCode( CONST_GL_STENCIL_PASS_DEPTH_FAIL, "gl.STENCIL_PASS_DEPTH_FAIL", TYPE_I32, FromI32(0x0B95))
	AddConstCode( CONST_GL_STENCIL_PASS_DEPTH_PASS, "gl.STENCIL_PASS_DEPTH_PASS", TYPE_I32, FromI32(0x0B96))
	AddConstCode( CONST_GL_STENCIL_REF            , "gl.STENCIL_REF"            , TYPE_I32, FromI32(0x0B97))
	AddConstCode( CONST_GL_STENCIL_WRITEMASK      , "gl.STENCIL_WRITE_MASK"     , TYPE_I32, FromI32(0x0B98))
	AddConstCode( CONST_GL_DITHER                 , "gl.DITHER"                 , TYPE_I32, FromI32(0x0BD0))
	AddConstCode( CONST_GL_BLEND                  , "gl.BLEND"                  , TYPE_I32, FromI32(0x0BE2))
	AddConstCode( CONST_GL_SCISSOR_TEST           , "gl.SCISSOR_TEST"           , TYPE_I32, FromI32(0x0C11))
	AddConstCode( CONST_GL_POLYGON_SMOOTH_HINT    , "gl.POLYGON_SMOOTH_HINT"    , TYPE_I32, FromI32(0x0C53))
	AddConstCode( CONST_GL_TEXTURE_2D             , "gl.TEXTURE_2D"             , TYPE_I32, FromI32(0x0DE1))
	AddConstCode( CONST_GL_TEXTURE_WIDTH          , "gl.TEXTURE_WIDTH"          , TYPE_I32, FromI32(0x1000))
	AddConstCode( CONST_GL_TEXTURE_HEIGHT         , "gl.TEXTURE_HEIGHT"         , TYPE_I32, FromI32(0x1001))
	AddConstCode( CONST_GL_DONT_CARE              , "gl.DONT_CARE"              , TYPE_I32, FromI32(0x1100))
	AddConstCode( CONST_GL_UNSIGNED_BYTE          , "gl.UNSIGNED_BYTE"          , TYPE_I32, FromI32(0x1401))
	AddConstCode( CONST_GL_FLOAT                  , "gl.FLOAT"                  , TYPE_I32, FromI32(0x1406))
	AddConstCode( CONST_GL_INVERT                 , "gl.INVERT"                 , TYPE_I32, FromI32(0x150A))
	AddConstCode( CONST_GL_TEXTURE                , "gl.TEXTURE"                , TYPE_I32, FromI32(0x1702))
	AddConstCode( CONST_GL_COLOR                  , "gl.COLOR"                  , TYPE_I32, FromI32(0x1800))
	AddConstCode( CONST_GL_DEPTH                  , "gl.DEPTH"                  , TYPE_I32, FromI32(0x1801))
	AddConstCode( CONST_GL_STENCIL                , "gl.STENCIL"                , TYPE_I32, FromI32(0x1802))
	AddConstCode( CONST_GL_STENCIL_INDEX          , "gl.STENCIL_INDEX"          , TYPE_I32, FromI32(0x1901))
	AddConstCode( CONST_GL_DEPTH_COMPONENT        , "gl.DEPTH_COMPONENT"        , TYPE_I32, FromI32(0x1902))
	AddConstCode( CONST_GL_RGBA                   , "gl.RGBA"                   , TYPE_I32, FromI32(0x1908))
	AddConstCode( CONST_GL_KEEP                   , "gl.KEEP"                   , TYPE_I32, FromI32(0x1E00))
	AddConstCode( CONST_GL_REPLACE                , "gl.REPLACE"                , TYPE_I32, FromI32(0x1E01))
	AddConstCode( CONST_GL_INCR                   , "gl.INCR"                   , TYPE_I32, FromI32(0x1E02))
	AddConstCode( CONST_GL_DECR                   , "gl.DECR"                   , TYPE_I32, FromI32(0x1E03))
	AddConstCode( CONST_GL_NEAREST                , "gl.NEAREST"                , TYPE_I32, FromI32(0x2600))
	AddConstCode( CONST_GL_LINEAR                 , "gl.LINEAR"                 , TYPE_I32, FromI32(0x2601))
	AddConstCode( CONST_GL_NEAREST_MIPMAP_NEAREST , "gl.NEAREST_MIPMAP_NEAREST" , TYPE_I32, FromI32(0x2700))
	AddConstCode( CONST_GL_LINEAR_MIPMAP_NEAREST  , "gl.LINEAR_MIPMAP_NEAREST"  , TYPE_I32, FromI32(0x2701))
	AddConstCode( CONST_GL_NEAREST_MIPMAP_LINEAR  , "gl.NEAREST_MIPMAP_LINEAR"  , TYPE_I32, FromI32(0x2702))
	AddConstCode( CONST_GL_LINEAR_MIPMAP_LINEAR   , "gl.LINEAR_MIPMAP_LINEAR"   , TYPE_I32, FromI32(0x2703))
	AddConstCode( CONST_GL_TEXTURE_MAG_FILTER     , "gl.TEXTURE_MAG_FILTER"     , TYPE_I32, FromI32(0x2800))
	AddConstCode( CONST_GL_TEXTURE_MIN_FILTER     , "gl.TEXTURE_MIN_FILTER"     , TYPE_I32, FromI32(0x2801))
	AddConstCode( CONST_GL_TEXTURE_WRAP_S         , "gl.TEXTURE_WRAP_S"         , TYPE_I32, FromI32(0x2802))
	AddConstCode( CONST_GL_TEXTURE_WRAP_T         , "gl.TEXTURE_WRAP_T"         , TYPE_I32, FromI32(0x2803))
	AddConstCode( CONST_GL_REPEAT                 , "gl.REPEAT"                 , TYPE_I32, FromI32(0x2901))

	// gl_1_1
	AddConstCode( CONST_GL_RGBA8                 , "gl.RGBA8"                 , TYPE_I32, FromI32(0x8058))
	AddConstCode( CONST_GL_VERTEX_ARRAY          , "gl.VERTEX_ARRAY"          , TYPE_I32, FromI32(0x8074))

	// gl_1_2
	AddConstCode( CONST_GL_TEXTURE_WRAP_R        , "gl.TEXTURE_WRAP_R"        , TYPE_I32, FromI32(0x8072))
	AddConstCode( CONST_GL_CLAMP_TO_EDGE         , "gl.CLAMP_TO_EDGE"         , TYPE_I32, FromI32(0x812F))

	// gl_1_3
	AddConstCode( CONST_GL_TEXTURE0              , "gl.TEXTURE0"              , TYPE_I32, FromI32(0x84C0))
	AddConstCode( CONST_GL_MULTISAMPLE_ARB       , "gl.MULTISAMPLE_ARB"       , TYPE_I32, FromI32(0x809D)) // remove _ARB
	AddConstCode( CONST_GL_CLAMP_TO_BORDER       , "gl.CLAMP_TO_BORDER"       , TYPE_I32, FromI32(0x812D))

	// gl_1_4
	AddConstCode( CONST_GL_DEPTH_COMPONENT16     , "gl.DEPTH_COMPONENT16"     , TYPE_I32, FromI32(0x81A5))
	AddConstCode( CONST_GL_DEPTH_COMPONENT24     , "gl.DEPTH_COMPONENT24"     , TYPE_I32, FromI32(0x81A6))
	AddConstCode( CONST_GL_DEPTH_COMPONENT32     , "gl.DEPTH_COMPONENT32"     , TYPE_I32, FromI32(0x81A7))
	AddConstCode( CONST_GL_MIRRORED_REPEAT       , "gl.MIRRORED_REPEAT"       , TYPE_I32, FromI32(0x8370))
	AddConstCode( CONST_GL_INCR_WRAP             , "gl.INCR_WRAP"             , TYPE_I32, FromI32(0x8507))
	AddConstCode( CONST_GL_DECR_WRAP             , "gl.DECR_WRAP"             , TYPE_I32, FromI32(0x8508))

	// gl_1_5
	AddConstCode( CONST_GL_ARRAY_BUFFER          , "gl.ARRAY_BUFFER"          , TYPE_I32, FromI32(0x8892))
	AddConstCode( CONST_GL_STREAM_DRAW           , "gl.STREAM_DRAW"           , TYPE_I32, FromI32(0x88E0))
	AddConstCode( CONST_GL_STREAM_READ           , "gl.STREAM_READ"           , TYPE_I32, FromI32(0x88E1))
	AddConstCode( CONST_GL_STREAM_COPY           , "gl.STREAM_COPY"           , TYPE_I32, FromI32(0x88E2))
	AddConstCode( CONST_GL_STATIC_DRAW           , "gl.STATIC_DRAW"           , TYPE_I32, FromI32(0x88E4))
	AddConstCode( CONST_GL_STATIC_READ           , "gl.STATIC_READ"           , TYPE_I32, FromI32(0x88E5))
	AddConstCode( CONST_GL_STATIC_COPY           , "gl.STATIC_COPY"           , TYPE_I32, FromI32(0x88E6))
	AddConstCode( CONST_GL_DYNAMIC_DRAW          , "gl.DYNAMIC_DRAW"          , TYPE_I32, FromI32(0x88E8))
	AddConstCode( CONST_GL_DYNAMIC_READ          , "gl.DYNAMIC_READ"          , TYPE_I32, FromI32(0x88E9))
	AddConstCode( CONST_GL_DYNAMIC_COPY          , "gl.DYNAMIC_COPY"          , TYPE_I32, FromI32(0x88EA))

	// gl_2_0
	AddConstCode( CONST_GL_STENCIL_BACK_FUNC           , "gl.STENCIL_BACK_FUNC"           , TYPE_I32, FromI32(0x8800))
	AddConstCode( CONST_GL_STENCIL_BACK_FAIL           , "gl.STENCIL_BACK_FAIL"           , TYPE_I32, FromI32(0x8801))
	AddConstCode( CONST_GL_STENCIL_BACK_PASS_DEPTH_FAIL, "gl.STENCIL_BACK_PASS_DEPTH_FAIL", TYPE_I32, FromI32(0x8802))
	AddConstCode( CONST_GL_STENCIL_BACK_PASS_DEPTH_PASS, "gl.STENCIL_BACK_PASS_DEPTH_PASS", TYPE_I32, FromI32(0x8803))
	AddConstCode( CONST_GL_FRAGMENT_SHADER             , "gl.FRAGMENT_SHADER"             , TYPE_I32, FromI32(0x8B30))
	AddConstCode( CONST_GL_VERTEX_SHADER               , "gl.VERTEX_SHADER"               , TYPE_I32, FromI32(0x8B31))
	AddConstCode( CONST_GL_STENCIL_BACK_REF            , "gl.STENCIL_BACK_REF"            , TYPE_I32, FromI32(0x8CA3))
	AddConstCode( CONST_GL_STENCIL_BACK_VALUE_MASK     , "gl.STENCIL_BACK_VALUE_MASK"     , TYPE_I32, FromI32(0x8CA4))
	AddConstCode( CONST_GL_STENCIL_BACK_WRITEMASK      , "gl.STENCIL_BACK_WRITEMASK"      , TYPE_I32, FromI32(0x8CA5))

	// gl_3_0
	AddConstCode( CONST_GL_DEPTH_COMPONENT32F                        , "gl.DEPTH_COMPONENT32F"                       , TYPE_I32, FromI32(0x8CAC))
	AddConstCode( CONST_GL_DEPTH32F_STENCIL8                         , "gl.DEPTH32F_STENCIL8"                        , TYPE_I32, FromI32(0x8CAD))
	AddConstCode( CONST_GL_FRAMEBUFFER_UNDEFINED                     , "gl.FRAMEBUFFER_UNDEFINED"                    , TYPE_I32, FromI32(0x8219))
	AddConstCode( CONST_GL_DEPTH_STENCIL_ATTACHMENT                  , "gl.DEPTH_STENCIL_ATTACHMENT"                 , TYPE_I32, FromI32(0x821A))
	AddConstCode( CONST_GL_DEPTH_STENCIL                             , "gl.DEPTH_STENCIL"                            , TYPE_I32, FromI32(0x84F9))
	AddConstCode( CONST_GL_DEPTH24_STENCIL8                          , "gl.DEPTH24_STENCIL8"                         , TYPE_I32, FromI32(0x88F0))
	AddConstCode( CONST_GL_FRAMEBUFFER_COMPLETE                      , "gl.FRAMEBUFFER_COMPLETE"                     , TYPE_I32, FromI32(0x8CD5))
	AddConstCode( CONST_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT         , "gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT"        , TYPE_I32, FromI32(0x8CD6))
	AddConstCode( CONST_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT , "gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT", TYPE_I32, FromI32(0x8CD7))
	AddConstCode( CONST_GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER        , "gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER"       , TYPE_I32, FromI32(0x8CDB))
	AddConstCode( CONST_GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER        , "gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER"       , TYPE_I32, FromI32(0x8CDC))
	AddConstCode( CONST_GL_FRAMEBUFFER_UNSUPPORTED                   , "gl.FRAMEBUFFER_UNSUPPORTED"                  , TYPE_I32, FromI32(0x8CDD))
	AddConstCode( CONST_GL_COLOR_ATTACHMENT0                         , "gl.COLOR_ATTACHMENT0"                        , TYPE_I32, FromI32(0x8CE0))
	AddConstCode( CONST_GL_DEPTH_ATTACHMENT                          , "gl.DEPTH_ATTACHMENT"                         , TYPE_I32, FromI32(0x8D00))
	AddConstCode( CONST_GL_STENCIL_ATTACHMENT                        , "gl.STENCIL_ATTACHMENT"                       , TYPE_I32, FromI32(0x8D20))
	AddConstCode( CONST_GL_FRAMEBUFFER                               , "gl.FRAMEBUFFER"                              , TYPE_I32, FromI32(0x8D40))
	AddConstCode( CONST_GL_RENDERBUFFER                              , "gl.RENDERBUFFER"                             , TYPE_I32, FromI32(0x8D41))
	AddConstCode( CONST_GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE        , "gl.FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"       , TYPE_I32, FromI32(0x8D56))
	AddConstCode( CONST_GL_STENCIL_INDEX1                            , "gl.STENCIL_INDEX1"                           , TYPE_I32, FromI32(0x8D46))
	AddConstCode( CONST_GL_STENCIL_INDEX4                            , "gl.STENCIL_INDEX4"                           , TYPE_I32, FromI32(0x8D47))
	AddConstCode( CONST_GL_STENCIL_INDEX8                            , "gl.STENCIL_INDEX8"                           , TYPE_I32, FromI32(0x8D48))
	AddConstCode( CONST_GL_STENCIL_INDEX16                           , "gl.STENCIL_INDEX16"                          , TYPE_I32, FromI32(0x8D49))

	// gl_3_2
	AddConstCode( CONST_GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS      , "gl.FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS"     , TYPE_I32, FromI32(0x8DA8))

	// gl_4_4
	AddConstCode( CONST_GL_MIRROR_CLAMP_TO_EDGE   , "gl.MIRROR_CLAMP_TO_EDGE", TYPE_I32, FromI32(0x8743))

	// Fixed pipeline. Deprecated ?
	AddConstCode( CONST_GL_POLYGON                , "gl.POLYGON"          , TYPE_I32, FromI32(9))
	AddConstCode( CONST_GL_MODELVIEW              , "gl.MODELVIEW"        , TYPE_I32, FromI32(5888))
	AddConstCode( CONST_GL_PROJECTION             , "gl.PROJECTION"       , TYPE_I32, FromI32(5889))
	AddConstCode( CONST_GL_MODELVIEW_MATRIX       , "gl.MODELVIEW_MATRIX" , TYPE_I32, FromI32(2982))
	AddConstCode( CONST_GL_LIGHTING               , "gl.LIGHTING"         , TYPE_I32, FromI32(2896))
	AddConstCode( CONST_GL_LIGHT0                 , "gl.LIGHT0"           , TYPE_I32, FromI32(16384))
	AddConstCode( CONST_GL_AMBIENT                , "gl.AMBIENT"          , TYPE_I32, FromI32(4608))
	AddConstCode( CONST_GL_DIFFUSE                , "gl.DIFFUSE"          , TYPE_I32, FromI32(4609))
	AddConstCode( CONST_GL_POSITION               , "gl.POSITION"         , TYPE_I32, FromI32(4611))
	AddConstCode( CONST_GL_TEXTURE_ENV            , "gl.TEXTURE_ENV"      , TYPE_I32, FromI32(8960))
	AddConstCode( CONST_GL_TEXTURE_ENV_MODE       , "gl.TEXTURE_ENV_MODE" , TYPE_I32, FromI32(8704))
	AddConstCode( CONST_GL_MODULATE               , "gl.MODULATE"         , TYPE_I32, FromI32(8448))
	AddConstCode( CONST_GL_DECAL                  , "gl.DECAL"            , TYPE_I32, FromI32(8449))
	AddConstCode( CONST_GL_POINT_SMOOTH           , "gl.POINT_SMOOTH"     , TYPE_I32, FromI32(2832))

	// glfw
	AddConstCode( CONST_GLFW_FALSE                     , "glfw.False"                    , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GLFW_TRUE                      , "glfw.True"                     , TYPE_I32, FromI32(1))
	AddConstCode( CONST_GLFW_PRESS                     , "glfw.Press"                    , TYPE_I32, FromI32(1))
	AddConstCode( CONST_GLFW_RELEASE                   , "glfw.Release"                  , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GLFW_REPEAT                    , "glfw.Repeat"                   , TYPE_I32, FromI32(2))
	AddConstCode( CONST_GLFW_KEY_UNKNOWN               , "glfw.KeyUnknown"               , TYPE_I32, FromI32(-1))
	AddConstCode( CONST_GLFW_CURSOR                    , "glfw.Cursor"                   , TYPE_I32, FromI32(208897))
	AddConstCode( CONST_GLFW_STICKY_KEYS               , "glfw.StickyKeys"               , TYPE_I32, FromI32(208898))
	AddConstCode( CONST_GLFW_STICKY_MOUSE_BUTTONS      , "glfw.StickyMouseButtons"       , TYPE_I32, FromI32(208899))
	AddConstCode( CONST_GLFW_CURSOR_NORMAL             , "glfw.CursorNormal"             , TYPE_I32, FromI32(212993))
	AddConstCode( CONST_GLFW_CURSOR_HIDDEN             , "glfw.CursorHidden"             , TYPE_I32, FromI32(212994))
	AddConstCode( CONST_GLFW_CURSOR_DISABLED           , "glfw.CursorDisabled"           , TYPE_I32, FromI32(212995))
	AddConstCode( CONST_GLFW_RESIZABLE                 , "glfw.Resizable"                , TYPE_I32, FromI32(131075))
	AddConstCode( CONST_GLFW_CONTEXT_VERSION_MAJOR     , "glfw.ContextVersionMajor"      , TYPE_I32, FromI32(139266))
	AddConstCode( CONST_GLFW_CONTEXT_VERSION_MINOR     , "glfw.ContextVersionMinor"      , TYPE_I32, FromI32(139267))
	AddConstCode( CONST_GLFW_OPENGL_PROFILE            , "glfw.Opengl.Profile"           , TYPE_I32, FromI32(139272))
	AddConstCode( CONST_GLFW_OPENGL_COREPROFILE        , "glfw.Opengl.Coreprofile"       , TYPE_I32, FromI32(204801))
	AddConstCode( CONST_GLFW_OPENGL_FORWARD_COMPATIBLE , "glfw.Opengl.ForwardCompatible" , TYPE_I32, FromI32(139270))
	AddConstCode( CONST_GLFW_MOUSE_BUTTON_LAST         , "glfw.MouseButtonLast"          , TYPE_I32, FromI32(7))
	AddConstCode( CONST_GLFW_MOUSE_BUTTON_LEFT         , "glfw.MouseButtonLeft"          , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GLFW_MOUSE_BUTTON_RIGHT        , "glfw.MouseButtonRight"         , TYPE_I32, FromI32(1))
	AddConstCode( CONST_GLFW_MOUSE_BUTTON_MIDDLE       , "glfw.MouseButtonMiddle"        , TYPE_I32, FromI32(2))

	// gltext
	AddConstCode( CONST_GLTEXT_LEFT_TO_RIGHT           , "gltext.LeftToRight"            , TYPE_I32, FromI32(0))
	AddConstCode( CONST_GLTEXT_RIGHT_TO_LEFT           , "gltext.RightToLeft"            , TYPE_I32, FromI32(1))
	AddConstCode( CONST_GLTEXT_TOP_TO_BOTTOM           , "gltext.TopToBottom"            , TYPE_I32, FromI32(2))

	// os
	AddConstCode( CONST_OS_RUN_SUCCESS                 , "os.RUN_SUCCESS"                , TYPE_I32, FromI32(OS_RUN_SUCCESS))
	AddConstCode( CONST_OS_RUN_EMPTY_CMD               , "os.RUN_EMPTY_CMD"              , TYPE_I32, FromI32(OS_RUN_EMPTY_CMD))
	AddConstCode( CONST_OS_RUN_PANIC                   , "os.RUN_PANIC"                  , TYPE_I32, FromI32(OS_RUN_PANIC))
	AddConstCode( CONST_OS_RUN_START_FAILED            , "os.RUN_START_FAILED"           , TYPE_I32, FromI32(OS_RUN_START_FAILED))
	AddConstCode( CONST_OS_RUN_WAIT_FAILED             , "os.RUN_WAIT_FAILED"            , TYPE_I32, FromI32(OS_RUN_WAIT_FAILED))
	AddConstCode( CONST_OS_RUN_TIMEOUT                 , "os.RUN_TIMEOUT"                , TYPE_I32, FromI32(OS_RUN_TIMEOUT))

	// cx
	AddConstCode( CONST_CX_SUCCESS                     , "cx.SUCCESS"                    , TYPE_I32, FromI32(CX_SUCCESS))
	AddConstCode( CONST_CX_RUNTIME_ERROR               , "cx.RUNTIME_ERROR"              , TYPE_I32, FromI32(CX_RUNTIME_ERROR))
	AddConstCode( CONST_CX_PANIC                       , "cx.PANIC"                      , TYPE_I32, FromI32(CX_PANIC))
	AddConstCode( CONST_CX_COMPILATION_ERROR           , "cx.COMPILATION_ERROR"          , TYPE_I32, FromI32(CX_COMPILATION_ERROR))
	AddConstCode( CONST_CX_INTERNAL_ERROR              , "cx.INTERNAL_ERROR"             , TYPE_I32, FromI32(CX_INTERNAL_ERROR))
	AddConstCode( CONST_CX_ASSERT                      , "cx.ASSERT"                     , TYPE_I32, FromI32(CX_ASSERT))
	AddConstCode( CONST_CX_RUNTIME_INVALID_ARGUMENT    , "cx.RUNTIME_INVALID_ARGUMENT"   , TYPE_I32, FromI32(CX_RUNTIME_INVALID_ARGUMENT))
	AddConstCode( CONST_CX_RUNTIME_SLICE_INDEX_OUT_OF_RANGE, "cx.RUNTIME_SLICE_INDEX_OUT_OF_RANGE", TYPE_I32, FromI32(CX_RUNTIME_SLICE_INDEX_OUT_OF_RANGE))
}
