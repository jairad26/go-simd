//go:build arm64
// +build arm64

#include "textflag.h"

// func AddVecSIMD(a, b, result *uint8, n uint)
TEXT ·AddVecSIMD(SB), NOSPLIT, $0-32
    // the arguments are *uint8 so that we dont have to load in the array's 0th element, pointer, and capacity
    // pointer, length, and capacity are 8 bytes each, so we save 16 bytes per slice doing this

    MOVD result+0(FP), R0  // pointer to 0th element of result
    MOVD a+8(FP), R1       // pointer to 0th element of first slice data
    MOVD b+16(FP), R2      // pointer to 0th element of second slice data
    MOVD len+24(FP), R3    // length of slices
    
    MOVD R3, R4            // Copy length for SIMD processing
    AND $~15, R4           // Round down to nearest multiple of 16
    /*
    A quick bit manipulation refresher for myself

    15 in bits: 000...00001111
    ~15 in bits: 111...11110000
    When you AND ~15 with any number, you get the nearest multiple of 16, bc it
    zeroes out the remainder

    You can also AND with 15 to get the remainder of the division, since it zeroes everything
    except the remainder!
    */

    // If no 16-byte chunks after bit manipulation, skip to remainder
    CBZ R4, remainder      // CBZ means compare and branch if zero
    
loop:
    
    // Load 16 bytes from each slice - note R1 and R2 now
    VLD1.P 16(R1), [V0.B16]  // 16 bits of R1 -> load into 16 bits of V0, then increment R1
    VLD1.P 16(R2), [V1.B16]  // 16 bits of R2 -> load into 16 bits of V1, then increment R2
    
    VADD V1.B16, V0.B16, V0.B16 // V1 + V0 -> V0
    
    // Store to result - note R0 now
    VST1.P [V0.B16], 16(R0) // 16 bits of V0 -> store into 16 bits of R0, then increment R0
    
    SUB $16, R4, R4 // Decrement counter by 16

    // Loop until processed all 16-byte chunks
    CBNZ R4, loop // if R4 is not zero, go to loop

remainder:
    // Calculate remaining elements
    AND $15, R3, R4 // $15 AND R3 -> R4, see above for bit manipulation explanation
    CBZ R4, done // if no remainder, go to done
    
remainder_loop:
    // Load bytes from a and b
    MOVBU (R1), R6 // move unsigned byte from R1 into R6
    MOVBU (R2), R7 // move unsigned byte from R2 into R7
    ADD R7, R6, R6 // R6 + R7 -> R6
    // Store to result - note R0 now
    MOVBU R6, (R0) // move byte from R6 into R0
    
    // Increment all pointers - note new register order
    ADD $1, R1
    ADD $1, R2
    ADD $1, R0
    
    SUB $1, R4, R4 // Decrement counter by 1
    CBNZ R4, remainder_loop

done:
    RET


// func SubVecSIMD(a, b, result *uint8, n uint)
TEXT ·SubVecSIMD(SB), NOSPLIT, $0-32
    MOVD result+0(FP), R0
    MOVD a+8(FP), R1
    MOVD b+16(FP), R2
    MOVD len+24(FP), R3

    MOVD R3, R4
    AND $~15, R4

    CBZ R4, remainder

loop:

    VLD1.P 16(R1), [V0.B16]  // from a
    VLD1.P 16(R2), [V1.B16]  // from b

    VSUB V1.B16, V0.B16, V0.B16

    VST1.P [V0.B16], 16(R0)

    SUB $16, R4, R4

    CBNZ R4, loop

remainder:
    AND $15, R3, R4
    CBZ R4, done

remainder_loop:
    MOVBU (R1), R6
    MOVBU (R2), R7
    SUB R7, R6, R6
    MOVBU R6, (R0)

    ADD $1, R1
    ADD $1, R2
    ADD $1, R0

    SUB $1, R4, R4
    CBNZ R4, remainder_loop

done:
    RET

// The function below is inspired from https://github.com/camdencheek/simd_blog/blob/main/dot_arm64.s
// Thank you @camdencheek for the great article https://sourcegraph.com/blog/slow-to-simd
// func DotVecSIMD16(a, b *uint8, len uint) int32
TEXT ·DotVecSIMD16(SB), NOSPLIT, $0-32
    MOVD a_base+0(FP), R0
    MOVD b_base+8(FP), R1
    MOVD len+16(FP), R2

    MOVD R2, R4
    AND $~15, R4

    // Zero V0, which will store 4 packed 32-bit sums
	VEOR V0.B16, V0.B16, V0.B16
    MOVD $0, R8

    CBZ R4, remainder

loop:
    

    VLD1.P 16(R0), [V1.B16]  // load 16 bytes from a
    VLD1.P 16(R1), [V2.B16]  // load 16 bytes from b

    // The following instruction is not supported by the go assembler, so use
	// the binary format. It would be the equivalent of the following instruction:
	//
    // UDOT V1.B16, V2.B16, V0.S4 // V1.B16 * V2.B16 -> V0.S4
    // this creates a dot product for each 4 bytes in V1 and V2, and stores the sum in V0
	//
	// Generated the binary form of the instruction using this godbolt setup:
	// https://godbolt.org/z/r5b1axedY
	WORD $0x6E829420

    SUB $16, R4, R4

    CBNZ R4, loop

remainder:
    // Calculate remaining elements
    AND $15, R2, R3
    CBZ R3, done    // Skip if no remainder

remainder_loop:
    
    // Load single bytes and multiply
    MOVBU.P 1(R0), R5
    MOVBU.P 1(R1), R6
    MUL R5, R6, R7
    ADD R7, R8      // Accumulate in R8
    
    SUB $1, R3

    CBNZ R3, remainder_loop

done:
    // Add remainder sum to vector sum
    VADDV V0.S4, V0 // adds the 4 32-bit values in V0 to a single 32-bit value
    VMOV V0.S[0], R6 // writes the 32-bit value in the first lane of V0 to R6
    ADD R8, R6      // Add remainder sum
    MOVD R6, ret+24(FP)
    RET




// func DotVecSIMD32(a, b *uint8, len uint) int32
TEXT ·DotVecSIMD32(SB), NOSPLIT, $0-32
    MOVD a_base+0(FP), R0
    MOVD b_base+8(FP), R1
    MOVD len+16(FP), R2

    MOVD R2, R4
    AND $~31, R4

	VEOR V0.B16, V0.B16, V0.B16
    VEOR V1.B16, V1.B16, V1.B16
    MOVD $0, R8

    CBZ R4, remainder

loop:
    

    VLD1.P 32(R0), [V2.B16, V3.B16] 
    VLD1.P 32(R1), [V4.B16, V5.B16]

    // The following instruction is not supported by the go assembler, so use
	// the binary format. It would be the equivalent of the following instruction:
	//
    // UDOT V2.B16, V4.B16, V0.S4
    // UDOT V3.B16, V5.B16, V1.S4
	//
	// Generated the binary form of the instruction using this godbolt setup:
	// https://godbolt.org/z/EdbxvTvhz
	WORD $0x6E849440
    WORD $0x6E859461

    SUB $32, R4, R4

    CBNZ R4, loop

    VADD V0.S4, V1.S4, V0.S4

remainder:
    AND $31, R2, R3
    CBZ R3, done

remainder_loop:
    
    MOVBU.P 1(R0), R5
    MOVBU.P 1(R1), R6
    MUL R5, R6, R7
    ADD R7, R8
    
    SUB $1, R3
    CBNZ R3, remainder_loop

done:
    VADDV V0.S4, V0
    VMOV V0.S[0], R6
    ADD R8, R6
    MOVD R6, ret+24(FP)
    RET




// func DotVecSIMD64(a, b *uint8, len uint) int32
TEXT ·DotVecSIMD64(SB), NOSPLIT, $0-32
    MOVD a_base+0(FP), R0
    MOVD b_base+8(FP), R1
    MOVD len+16(FP), R2

    MOVD R2, R4
    AND $~63, R4

	VEOR V0.B16, V0.B16, V0.B16
    VEOR V1.B16, V1.B16, V1.B16
    VEOR V2.B16, V2.B16, V2.B16
    VEOR V3.B16, V3.B16, V3.B16
    MOVD $0, R8

    CBZ R4, remainder

loop:
    

    VLD1.P 64(R0), [V4.B16, V5.B16, V6.B16, V7.B16] 
    VLD1.P 64(R1), [V8.B16, V9.B16, V10.B16, V11.B16]

    // The following instruction is not supported by the go assembler, so use
	// the binary format. It would be the equivalent of the following instruction:
	//
    // SDOT V4.B16, V8.B16, V0.S4
    // SDOT V5.B16, V9.B16, V1.S4
    // SDOT V6.B16, V10.B16, V2.S4
    // SDOT V7.B16, V11.B16, V3.S4
    // this creates a dot product for each 4 bytes in V1 and V2, and stores the sum in V0
	//
	// Generated the binary form of the instruction using this godbolt setup:
	// https://godbolt.org/z/M45roP43Y
	WORD $0x6E889480
    WORD $0x6E8994A1
    WORD $0x6E8A94C2
    WORD $0x6E8B94E3

    SUB $64, R4, R4

    CBNZ R4, loop

    VADD V0.S4, V1.S4, V0.S4
    VADD V2.S4, V3.S4, V2.S4
    VADD V0.S4, V2.S4, V0.S4

remainder:
    AND $63, R2, R3
    CBZ R3, done

remainder_loop:
    
    MOVBU.P 1(R0), R5
    MOVBU.P 1(R1), R6
    MUL R5, R6, R7
    ADD R7, R8
    
    SUB $1, R3
    CBNZ R3, remainder_loop

done:
    VADDV V0.S4, V0
    VMOV V0.S[0], R6
    ADD R8, R6
    MOVD R6, ret+24(FP)
    RET
