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

### How lanes work
128-bit V register with different lane configurations:

16 Byte lanes (V0.16B): each lane is 8 bits * 16 = 128 bits
[B0][B1][B2][B3][B4][B5][B6][B7][B8][B9][B10][B11][B12][B13][B14][B15]
 └─┘ └─┘ └─┘ └─┘ └─┘ └─┘ └─┘ └─┘ └─┘ └─┘  └─┘  └─┘  └─┘  └─┘  └─┘  └─┘
Lane0 1   2   3   4   5   6   7   8   9   10   11   12   13   14   15

8 Halfword lanes (V0.8H): each lane is 16 bits * 8 = 128 bits
[   H0   ][   H1   ][   H2   ][   H3   ][   H4   ][   H5   ][   H6   ][   H7   ]
 └──────┘  └──────┘  └──────┘  └──────┘  └──────┘  └──────┘  └──────┘  └──────┘
  Lane0     Lane1     Lane2     Lane3     Lane4     Lane5     Lane6     Lane7

4 Word lanes (V0.4S): each lane is 32 bits * 4 = 128 bits
[       S0       ][       S1       ][       S2       ][       S3       ]
 └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘
      Lane0            Lane1            Lane2            Lane3

2 Double lanes (V0.2D):
[                D0                ][                D1                ]
 └────────────────────────────────┘  └────────────────────────────────┘
              Lane0                               Lane1

### Moving data between lanes
You can access lanes using [] format, so if you wanted to access the 0th lane of a 16B register, you would use V0.B[0]
You can use INS to insert data into a lane
INS V0.S[1], V0.S[0]
    └dest─┘  └─src─┘
This would insert the value from the 0th lane of V0 into the 1st lane of V0

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

 */