/*
# NEON 101

## Introduction
NEON is a SIMD architecture extension for ARM processors. SIMD stands for Single Instruction, Multiple Data. This means that a single 
instruction can operate on multiple data elements in parallel. NEON is a 128-bit SIMD architecture extension for ARM processors. 
It can operate on 16 bytes, 8 halfwords, 4 words, or 2 doublewords in parallel.

## R vs V Registers
R registers are general purpose registers in ARM architecture. V registers are NEON registers. NEON registers are 128-bit wide and can
hold 16 bytes, 8 halfwords, 4 words, or 2 doublewords.

## Byte naming
In NEON, data elements are named as follows:
- B: Byte (8 bits)
- H: Halfword (16 bits)
- S: Single precision floating point (32 bits)
- D: Double precision floating point (64 bits)

### Suffixing Registers
- Since Vectors can hold 128 bits, the most normal suffixes are as follows, since it adds up to 128
    - B16
    - H8
    - S4
    - D2
- To use the bottom half of the register, you can use the half version of the suffix
    - B8
    - H4
    - S2
    - D1

## Vector operations
NEON supports a variety of vector operations. Some of the common vector operations are:
- Add: VADD
- Subtract: VSUB
- Multiply: VMUL
- Divide: VDIV
- Load: VLD
- Store: VST
- Bitwise AND: VAND
- Bitwise OR: VORR
- Bitwise XOR: VEOR
- Shift: VSHL, VSHR
- Compare: VCEQ, VCGT, VCLT
- Select: VSEL
- Dot product: VDOT (only available in ARMv8.3)
 */