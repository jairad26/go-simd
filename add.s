#include "textflag.h"

// The function below is inspired from https://github.com/camdencheek/simd_blog/blob/main/dot_arm64.s
// Thank you @camdencheek for the great article https://sourcegraph.com/blog/slow-to-simd
// func dotUint8Vec(a, b *uint8, len int) uint32
TEXT ·dotUint8Vec(SB), NOSPLIT, $0-32
    MOVD a_base+0(FP), R0
    MOVD b_base+8(FP), R1
    MOVD len+16(FP), R2

    MOVD R2, R4
    AND $~15, R4

    // Zero V0, which will store 4 packed 32-bit sums
	VEOR V0.B16, V0.B16, V0.B16

    CBZ R4, remainder

    MOVD R4, R5

loop:
    CBZ R5, remainder

    VLD1.P 16(R0), [V1.B16]  // from a
    VLD1.P 16(R1), [V2.B16]  // from b

    // The following instruction is not supported by the go assembler, so use
	// the binary format. It would be the equivalent of the following instruction:
	//
    // SDOT V1.B16, V2.B16, V0.S4
	//
	// I generated the binary form of the instruction using this godbolt setup:
	// https://godbolt.org/z/3jPohn4dn
	WORD $0x4E829420

    SUB $16, R5, R5

    JMP loop

remainder:
    // Calculate remaining elements
    SUB R4, R2, R3  // R3 = len - (len & ~15)
    CBZ R3, done    // Skip if no remainder

remainder_loop:
    CBZ R3, done
    
    // Load single bytes and multiply
    MOVBU.P 1(R0), R5
    MOVBU.P 1(R1), R6
    MUL R5, R6, R7
    ADD R7, R8      // Accumulate in R8
    
    SUB $1, R3
    JMP remainder_loop

done:
    // Add remainder sum to vector sum
    VADDV V0.S4, V0
    VMOV V0.S[0], R6
    ADD R8, R6      // Add remainder sum
    MOVD R6, ret+24(FP)
    RET



// func addUint8Vec(a, b, result *uint8, n int)
TEXT ·addUint8Vec(SB), NOSPLIT, $0-32
    MOVD result+0(FP), R0  // pointer to result slice data
    MOVD a+8(FP), R1       // pointer to first slice data
    MOVD b+16(FP), R2      // pointer to second slice data
    MOVD len+24(FP), R3    // length of slices
    
    MOVD R3, R4            // Copy length for SIMD processing
    AND $~15, R4           // Round down to nearest multiple of 16
    
    CBZ R4, remainder      // If no 16-byte chunks, skip to remainder
    
    MOVD R4, R5           // Counter for SIMD loop

loop:
    CBZ R5, remainder     // If R5 == 0 go to remainder
    
    // Load 16 bytes from each slice - note R1 and R2 now
    VLD1.P 16(R1), [V0.B16]  // from a
    VLD1.P 16(R2), [V1.B16]  // from b
    
    VADD V1.B16, V0.B16, V0.B16
    
    // Store to result - note R0 now
    VST1.P [V0.B16], 16(R0)
    
    SUB $16, R5, R5
    B loop

remainder:
    AND $15, R3, R5
    CBZ R5, done
    
remainder_loop:
    // Load bytes from a and b - note R1 and R2 now
    MOVBU (R1), R6
    MOVBU (R2), R7
    ADD R7, R6, R6
    // Store to result - note R0 now
    MOVB R6, (R0)
    
    // Increment all pointers - note new register order
    ADD $1, R1
    ADD $1, R2
    ADD $1, R0
    
    SUB $1, R5, R5
    CBNZ R5, remainder_loop

done:
    RET
