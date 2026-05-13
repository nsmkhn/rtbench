#include "textflag.h"

// func indexStopByte(b []byte) int
//
// Scans b for the first '"' (0x22) or '\' (0x5C) byte using NEON.
// Processes 16 bytes per iteration in the fast path, falls back to
// byte-by-byte for the tail (< 16 remaining bytes).
// Returns len(b) if neither byte is found.
//
// Frame: 0 locals, 32 bytes args/return
//   b_base+0(FP)  ptr
//   b_len+8(FP)   len
//   b_cap+16(FP)  cap  (unused)
//   ret+24(FP)    int
TEXT ·indexStopByte(SB),NOSPLIT,$0-32
	MOVD	b_base+0(FP), R0	// R0 = &b[0]
	MOVD	b_len+8(FP), R1		// R1 = len(b)
	MOVD	$0, R2			// R2 = current index

	CBZ	R1, notfound

	// Only set up NEON registers if we have >= 16 bytes to scan.
	CMP	$16, R1
	BLT	tail

	MOVD	$0x22, R3
	VDUP	R3, V0.B16		// V0 = ['"' x16]
	MOVD	$0x5C, R3
	VDUP	R3, V1.B16		// V1 = ['\' x16]

loop16:
	// Load 16 bytes; R0 is post-incremented.
	VLD1.P	16(R0), [V2.B16]
	// V3[i] = 0xFF if V2[i] == '"', else 0x00
	VCMEQ	V0.B16, V2.B16, V3.B16
	// V4[i] = 0xFF if V2[i] == '\', else 0x00
	VCMEQ	V1.B16, V2.B16, V4.B16
	// V5[i] = 0xFF if V2[i] is a stop byte.
	VORR	V3.B16, V4.B16, V5.B16
	// Check whether any lane matched: OR the two 64-bit halves into one GP register.
	VMOV	V5.D[0], R3
	VMOV	V5.D[1], R4
	ORR	R4, R3, R3
	CBNZ	R3, found16
	ADD	$16, R2, R2
	SUB	$16, R1, R1
	CMP	$16, R1
	BGE	loop16

tail:
	CBZ	R1, notfound

tailloop:
	// Load one byte; R0 is post-incremented.
	MOVBU.P	1(R0), R3
	CMP	$0x22, R3
	BEQ	found
	CMP	$0x5C, R3
	BEQ	found
	ADD	$1, R2, R2
	SUB	$1, R1, R1
	CBNZ	R1, tailloop

notfound:
	MOVD	b_len+8(FP), R0
	MOVD	R0, ret+24(FP)
	RET

found:
	// R2 holds the index of the matched byte.
	MOVD	R2, ret+24(FP)
	RET

found16:
	// A stop byte exists somewhere in the 16-byte chunk that was just loaded.
	// R0 was post-incremented past the chunk; walk back to find the exact byte.
	SUB	$16, R0, R4		// R4 = start of matched chunk
find16loop:
	MOVBU.P	1(R4), R3
	CMP	$0x22, R3
	BEQ	found
	CMP	$0x5C, R3
	BEQ	found
	ADD	$1, R2, R2
	B	find16loop
