package gonotes

// Variables
var foo string // without initialization
var bar string = "I am an initialized variable"

// String
var thisIsTheString string = "I am a string"

// Bool
var isItTrue bool = true
var isItRobot bool = false

// Numeric types (ripped from -> https://kps.hashnode.dev/learn-go-the-complete-course)
var i int = 404                       // Platform dependent
var i8 int8 = 127                     // -128 to 127
var i16 int16 = 32767                 // -2^15 to 2^15 - 1
var i32 int32 = -2147483647           // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807   // -2^63 to 2^63 - 1
var ui uint = 404                     // Platform dependent
var ui8 uint8 = 255                   // 0 to 255
var ui16 uint16 = 65535               // 0 to 2^16
var ui32 uint32 = 2147483647          // 0 to 2^32
var ui64 uint64 = 9223372036854775807 // 0 to 2^64
var uiptr uintptr                     // Integer representation of a memory address
