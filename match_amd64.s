//+build !noasm !appengine


// func find4SSE4(needle, haystack []byte, dst []uint16)
TEXT ·find4SSE4(SB), 7, $0
    MOVQ    needle+0(FP),R8             // R8: &needle
    MOVQ    haystack+24(FP),SI          // SI: &haystack
    MOVQ    haystack_len+32(FP), R10    // R10: len(haystack)
    MOVQ    dst+48(FP),DX               // DX: &dst
    MOVD    (R8), X0                    // X0: needle
    PXOR    X4, X4                      // X4: Zero

    SHRQ   $3, R10                      // len(haystack)/8
    CMPQ    R10 ,$0
    JEQ     done_find4
loopback_find4:
    MOVOU (SI),X1                       // haystack[x]

    // MPSADBW $0, X0, X1               // Compare lower part X1[0:12] to X0[0:4], store in X1
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xc8; BYTE $0x00

    PCMPEQW X4, X1                      // if result == 0 {set word to 0xffff}
    PACKSSWB X1, X1                     // Words->bytes
    PMOVMSKB X1, R9                     // Transfer to bits

    ADDQ $8, SI
    MOVB R9, (DX)
    ADDQ $1, DX
    SUBQ $1, R10
    JNZ loopback_find4

done_find4:    
    RET

// func find4SSE4s(needle, haystack string, dst []uint16)
TEXT ·find4SSE4s(SB), 7, $0
    MOVQ    needle+0(FP),R8             // R8: &needle
    MOVQ    haystack+16(FP),SI          // SI: &haystack
    MOVQ    haystack_len+24(FP), R10    // R10: len(haystack)
    MOVQ    dst+32(FP),DX               // DX: &dst
    MOVD    (R8), X0                    // X0: needle
    PXOR    X4, X4                      // X4: Zero

    SHRQ   $3, R10                      // len(haystack)/8
    CMPQ    R10 ,$0
    JEQ     done_find4s
loopback_find4s:
    MOVOU (SI),X1                       // haystack[x]

    // MPSADBW $0, X0, X1               // Compare lower part X1[0:12] to X0[0:4], store in X1
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xc8; BYTE $0x00

    PCMPEQW X4, X1                      // if result == 0 {set word to 0xffff}
    PACKSSWB X1, X1                     // Words->bytes
    PMOVMSKB X1, R9                     // Transfer to bits

    ADDQ $8, SI
    MOVB R9, (DX)
    ADDQ $1, DX
    SUBQ $1, R10
    JNZ loopback_find4s

done_find4s:    
    RET


// func find8SSE4(needle, haystack []byte, dst []uint32)
TEXT ·find8SSE4(SB), 7, $0
    MOVQ    needle+0(FP),R8             // R8: &needle
    MOVQ    haystack+24(FP),SI          // SI: &haystack
    MOVQ    haystack_len+32(FP), R10    // R10: len(haystack)
    MOVQ    dst+48(FP),DX               // DX: &dst

    MOVQ    (R8), X0                    // X0: needle
    PXOR    X7, X7                      // X4: Zero
    PCMPEQW X5, X5                      
    PCMPEQW X6, X6
    PSRLW  $8, X5                       //  0xffff >> 8 = 0x00ff, lower byte mask per word (shifts in zeros)
    PSLLW  $8, X6                       //  0xffff << 8  = 0xff00, upper byte mask per word.

    SHRQ   $4, R10                      // len(haystack)/16
    CMPQ    R10 ,$0
    JEQ     done_find8
loopback_find8:
    MOVOU (SI),X1                       // haystack[x]
    MOVOA X1, X2
    MOVOU 8(SI),X3                      // haystack[x+8]
    MOVOA X3, X4

    // MPSADBW $0, X0, X1               // Compare lower part X1[0:12] to X0[0:4], store in X1
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xc8; BYTE $0x00

    // MPSADBW $5, X0, X2               // Compare lower part X2[4:16] to X0[4:8], store in X2
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xd0; BYTE $0x05

    // MPSADBW $0, X0, X3               // Compare lower part X3[0:12] to X0[0:4], store in X3
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xd8; BYTE $0x00

    // MPSADBW $5, X0, X4               // Compare lower part X4[4:16] to X0[4:8], store in X4
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xe0; BYTE $0x05

    PCMPEQW X7, X1                      // if result == 0 {set word to 0xffff}
    PCMPEQW X7, X2                      // if result == 0 {set word to 0xffff}
    PCMPEQW X7, X3                      // if result == 0 {set word to 0xffff}
    PCMPEQW X7, X4                      // if result == 0 {set word to 0xffff}
    PAND    X5, X1                      // Lower result as bytes
    PAND    X6, X2                      // upper result as bytes
    PAND    X5, X3                      // Lower result as bytes
    PAND    X6, X4                      // upper result as bytes
    POR     X2, X1
    POR     X4, X3
    PMOVMSKB X1, R9                     // Transfer to bits
    PMOVMSKB X3, R8                     // Transfer to bits

    ADDQ $16, SI
    MOVW R9, (DX)
    MOVW R8, 2(DX)
    ADDQ $4, DX
    SUBQ $1, R10
    JNZ loopback_find8

done_find8:    
    RET



// func find8SSE4s(needle, haystack string, dst []uint32)
TEXT ·find8SSE4s(SB), 7, $0
    MOVQ    needle+0(FP),R8             // R8: &needle
    MOVQ    haystack+16(FP),SI          // SI: &haystack
    MOVQ    haystack_len+24(FP), R10    // R10: len(haystack)
    MOVQ    dst+32(FP),DX               // DX: &dst

    MOVQ    (R8), X0                    // X0: needle
    PXOR    X4, X4                      // X4: Zero
    PCMPEQW X5, X5                      
    PCMPEQW X6, X6
    PSRLW  $8, X5                       //  0xffff >> 8 = 0x00ff, lower byte mask per word (shifts in zeros)
    PSLLW  $8, X6                       //  0xffff << 8  = 0xff00, upper byte mask per word.

    SHRQ   $3, R10                      // len(haystack)/8
    CMPQ    R10 ,$0
    JEQ     done_find8s
loopback_find8s:
    MOVOU (SI),X1                       // haystack[x]
    MOVOA X1, X2

    // MPSADBW $0, X0, X1               // Compare lower part X1[0:12] to X0[0:4], store in X1
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xc8; BYTE $0x00

    // MPSADBW $5, X0, X2               // Compare lower part X2[4:16] to X0[4:8], store in X2
    BYTE $0x66; BYTE $0x0f; BYTE $0x3a
    BYTE $0x42; BYTE $0xd0; BYTE $0x05

    PCMPEQW X4, X1                      // if result == 0 {set word to 0xffff}
    PCMPEQW X4, X2                      // if result == 0 {set word to 0xffff}
    PAND    X5, X1                      // Lower result as bytes
    PAND    X6, X2                      // upper result as bytes
    POR     X2, X1
    PMOVMSKB X1, R9                     // Transfer to bits

    ADDQ $8, SI
    MOVW R9, (DX)
    ADDQ $2, DX
    SUBQ $1, R10
    JNZ loopback_find8s

done_find8s:    
    RET
