Personal Encoder
================

Coding challenge for Symantec

### Usage

Build the code into an executable with `go build`

Run the test suite with `go test`

### Purpose

With a given input file, export an output file where each line of input has been encoded. 
This is achieved by taking the ASCII values and turning them into their binary values - g = 0x67 = 01100111. 
Using the binary string and a 3-bit encoding, produce an encoded output. Follows Base64 encoding to pad the output. 
For any padded values, the string "$" is used. 
Every 3 bits is used to map to an index in the string "symantec". This index determines the resulting encoded value. 

`s = 0, y = 1, m = 2, etc...`
