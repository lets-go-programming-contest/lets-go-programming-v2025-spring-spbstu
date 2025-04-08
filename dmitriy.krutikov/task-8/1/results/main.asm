"".bubbleSort STEXT nosplit size=135 args=0x18 locals=0x18 funcid=0x0 align=0x0
	0x0000 00000 (main.go:5)	TEXT	"".bubbleSort(SB), NOSPLIT|ABIInternal, $24-24
	0x0000 00000 (main.go:5)	SUBQ	$24, SP
	0x0004 00004 (main.go:5)	MOVQ	BP, 16(SP)
	0x0009 00009 (main.go:5)	LEAQ	16(SP), BP
	0x000e 00014 (main.go:5)	MOVQ	AX, "".arr+32(FP)
	0x0013 00019 (main.go:5)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0013 00019 (main.go:5)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0013 00019 (main.go:5)	FUNCDATA	$5, "".bubbleSort.arginfo1(SB)
	0x0013 00019 (main.go:5)	FUNCDATA	$6, "".bubbleSort.argliveinfo(SB)
	0x0013 00019 (main.go:5)	PCDATA	$3, $1
	0x0013 00019 (main.go:6)	XORL	CX, CX
	0x0015 00021 (main.go:7)	JMP	29
	0x0017 00023 (main.go:7)	INCQ	CX
	0x001a 00026 (main.go:7)	MOVQ	SI, BX
	0x001d 00029 (main.go:7)	LEAQ	-1(BX), DX
	0x0021 00033 (main.go:7)	CMPQ	CX, DX
	0x0024 00036 (main.go:7)	JGE	42
	0x0026 00038 (main.go:7)	XORL	DX, DX
	0x0028 00040 (main.go:8)	JMP	58
	0x002a 00042 (main.go:14)	MOVQ	16(SP), BP
	0x002f 00047 (main.go:14)	ADDQ	$24, SP
	0x0033 00051 (main.go:14)	RET
	0x0034 00052 (main.go:8)	MOVQ	SI, BX
	0x0037 00055 (main.go:8)	MOVQ	DI, DX
	0x003a 00058 (main.go:8)	MOVQ	BX, SI
	0x003d 00061 (main.go:8)	SUBQ	CX, BX
	0x0040 00064 (main.go:8)	DECQ	BX
	0x0043 00067 (main.go:8)	CMPQ	DX, BX
	0x0046 00070 (main.go:8)	JGE	23
	0x0048 00072 (main.go:9)	CMPQ	SI, DX
	0x004b 00075 (main.go:9)	JLS	123
	0x004d 00077 (main.go:9)	MOVQ	(AX)(DX*8), BX
	0x0051 00081 (main.go:9)	LEAQ	1(DX), DI
	0x0055 00085 (main.go:9)	CMPQ	SI, DI
	0x0058 00088 (main.go:9)	JLS	112
	0x005a 00090 (main.go:9)	MOVQ	8(AX)(DX*8), R8
	0x005f 00095 (main.go:9)	NOP
	0x0060 00096 (main.go:9)	CMPQ	R8, BX
	0x0063 00099 (main.go:9)	JGE	52
	0x0065 00101 (main.go:10)	MOVQ	R8, (AX)(DX*8)
	0x0069 00105 (main.go:10)	MOVQ	BX, 8(AX)(DX*8)
	0x006e 00110 (main.go:10)	JMP	52
	0x0070 00112 (main.go:9)	MOVQ	DI, AX
	0x0073 00115 (main.go:9)	MOVQ	SI, CX
	0x0076 00118 (main.go:9)	PCDATA	$1, $1
	0x0076 00118 (main.go:9)	CALL	runtime.panicIndex(SB)
	0x007b 00123 (main.go:9)	MOVQ	DX, AX
	0x007e 00126 (main.go:9)	MOVQ	SI, CX
	0x0081 00129 (main.go:9)	CALL	runtime.panicIndex(SB)
	0x0086 00134 (main.go:9)	XCHGL	AX, AX
	0x0000 48 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 89  H...H.l$.H.l$.H.
	0x0010 44 24 20 31 c9 eb 06 48 ff c1 48 89 f3 48 8d 53  D$ 1...H..H..H.S
	0x0020 ff 48 39 d1 7d 04 31 d2 eb 10 48 8b 6c 24 10 48  .H9.}.1...H.l$.H
	0x0030 83 c4 18 c3 48 89 f3 48 89 fa 48 89 de 48 29 cb  ....H..H..H..H).
	0x0040 48 ff cb 48 39 da 7d cf 48 39 d6 76 2e 48 8b 1c  H..H9.}.H9.v.H..
	0x0050 d0 48 8d 7a 01 48 39 fe 76 16 4c 8b 44 d0 08 90  .H.z.H9.v.L.D...
	0x0060 49 39 d8 7d cf 4c 89 04 d0 48 89 5c d0 08 eb c4  I9.}.L...H.\....
	0x0070 48 89 f8 48 89 f1 e8 00 00 00 00 48 89 d0 48 89  H..H.......H..H.
	0x0080 f1 e8 00 00 00 00 90                             .......
	rel 119+4 t=7 runtime.panicIndex+0
	rel 130+4 t=7 runtime.panicIndex+0
"".main STEXT size=439 args=0x0 locals=0x78 funcid=0x0 align=0x0
	0x0000 00000 (main.go:16)	TEXT	"".main(SB), ABIInternal, $120-0
	0x0000 00000 (main.go:16)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:16)	PCDATA	$0, $-2
	0x0004 00004 (main.go:16)	JLS	429
	0x000a 00010 (main.go:16)	PCDATA	$0, $-1
	0x000a 00010 (main.go:16)	SUBQ	$120, SP
	0x000e 00014 (main.go:16)	MOVQ	BP, 112(SP)
	0x0013 00019 (main.go:16)	LEAQ	112(SP), BP
	0x0018 00024 (main.go:16)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x0018 00024 (main.go:16)	FUNCDATA	$1, gclocals·804d9f4c78bfde951f5b4bad61ed1894(SB)
	0x0018 00024 (main.go:16)	FUNCDATA	$2, "".main.stkobj(SB)
	0x0018 00024 (main.go:17)	LEAQ	type.[7]int(SB), AX
	0x001f 00031 (main.go:17)	PCDATA	$1, $0
	0x001f 00031 (main.go:17)	NOP
	0x0020 00032 (main.go:17)	CALL	runtime.newobject(SB)
	0x0025 00037 (main.go:17)	MOVQ	AX, ""..autotmp_38+40(SP)
	0x002a 00042 (main.go:17)	MOVQ	$5, (AX)
	0x0031 00049 (main.go:17)	MOVQ	$1, 8(AX)
	0x0039 00057 (main.go:17)	MOVQ	$5, 16(AX)
	0x0041 00065 (main.go:17)	MOVQ	$6666, 24(AX)
	0x0049 00073 (main.go:17)	MOVQ	$22, 32(AX)
	0x0051 00081 (main.go:17)	MOVQ	$11, 40(AX)
	0x0059 00089 (main.go:17)	MOVQ	$90, 48(AX)
	0x0061 00097 (main.go:18)	MOVUPS	X15, ""..autotmp_19+80(SP)
	0x0067 00103 (main.go:18)	MOVUPS	X15, ""..autotmp_19+96(SP)
	0x006d 00109 (main.go:18)	LEAQ	type.string(SB), CX
	0x0074 00116 (main.go:18)	MOVQ	CX, ""..autotmp_19+80(SP)
	0x0079 00121 (main.go:18)	LEAQ	""..stmp_0(SB), DX
	0x0080 00128 (main.go:18)	MOVQ	DX, ""..autotmp_19+88(SP)
	0x0085 00133 (main.go:18)	MOVL	$7, BX
	0x008a 00138 (main.go:18)	MOVQ	BX, CX
	0x008d 00141 (main.go:18)	PCDATA	$1, $1
	0x008d 00141 (main.go:18)	CALL	runtime.convTslice(SB)
	0x0092 00146 (main.go:18)	LEAQ	type.[]int(SB), CX
	0x0099 00153 (main.go:18)	MOVQ	CX, ""..autotmp_19+96(SP)
	0x009e 00158 (main.go:18)	MOVQ	AX, ""..autotmp_19+104(SP)
	0x00a3 00163 (<unknown line number>)	NOP
	0x00a3 00163 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x00aa 00170 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x00b1 00177 ($GOROOT/src/fmt/print.go:274)	MOVL	$2, DI
	0x00b6 00182 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x00b9 00185 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_19+80(SP), CX
	0x00be 00190 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $2
	0x00be 00190 ($GOROOT/src/fmt/print.go:274)	NOP
	0x00c0 00192 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00c5 00197 (main.go:20)	XCHGL	AX, AX
	0x00c6 00198 (main.go:7)	MOVQ	""..autotmp_38+40(SP), CX
	0x00cb 00203 (main.go:7)	XORL	AX, AX
	0x00cd 00205 (main.go:7)	JMP	210
	0x00cf 00207 (main.go:7)	INCQ	AX
	0x00d2 00210 (main.go:7)	CMPQ	AX, $6
	0x00d6 00214 (main.go:7)	JGE	220
	0x00d8 00216 (main.go:7)	XORL	DX, DX
	0x00da 00218 (main.go:8)	JMP	338
	0x00dc 00220 (main.go:22)	MOVUPS	X15, ""..autotmp_24+48(SP)
	0x00e2 00226 (main.go:22)	MOVUPS	X15, ""..autotmp_24+64(SP)
	0x00e8 00232 (main.go:22)	LEAQ	type.string(SB), DX
	0x00ef 00239 (main.go:22)	MOVQ	DX, ""..autotmp_24+48(SP)
	0x00f4 00244 (main.go:22)	LEAQ	""..stmp_1(SB), DX
	0x00fb 00251 (main.go:22)	MOVQ	DX, ""..autotmp_24+56(SP)
	0x0100 00256 (main.go:22)	MOVQ	CX, AX
	0x0103 00259 (main.go:22)	MOVL	$7, BX
	0x0108 00264 (main.go:22)	MOVQ	BX, CX
	0x010b 00267 (main.go:22)	PCDATA	$1, $3
	0x010b 00267 (main.go:22)	CALL	runtime.convTslice(SB)
	0x0110 00272 (main.go:22)	LEAQ	type.[]int(SB), DX
	0x0117 00279 (main.go:22)	MOVQ	DX, ""..autotmp_24+64(SP)
	0x011c 00284 (main.go:22)	MOVQ	AX, ""..autotmp_24+72(SP)
	0x0121 00289 (<unknown line number>)	NOP
	0x0121 00289 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), BX
	0x0128 00296 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x012f 00303 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_24+48(SP), CX
	0x0134 00308 ($GOROOT/src/fmt/print.go:274)	MOVL	$2, DI
	0x0139 00313 ($GOROOT/src/fmt/print.go:274)	MOVQ	DI, SI
	0x013c 00316 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $0
	0x013c 00316 ($GOROOT/src/fmt/print.go:274)	NOP
	0x0140 00320 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0145 00325 (main.go:23)	MOVQ	112(SP), BP
	0x014a 00330 (main.go:23)	ADDQ	$120, SP
	0x014e 00334 (main.go:23)	RET
	0x014f 00335 (main.go:8)	MOVQ	DI, DX
	0x0152 00338 (main.go:8)	LEAQ	-6(AX), SI
	0x0156 00342 (main.go:8)	NEGQ	SI
	0x0159 00345 (main.go:8)	NOP
	0x0160 00352 (main.go:8)	CMPQ	DX, SI
	0x0163 00355 (main.go:8)	JGE	207
	0x0169 00361 (main.go:9)	CMPQ	DX, $7
	0x016d 00365 (main.go:9)	JCC	415
	0x016f 00367 (main.go:9)	MOVQ	(CX)(DX*8), SI
	0x0173 00371 (main.go:9)	LEAQ	1(DX), DI
	0x0177 00375 (main.go:9)	CMPQ	DI, $7
	0x017b 00379 (main.go:9)	JCC	402
	0x017d 00381 (main.go:9)	MOVQ	8(CX)(DX*8), R8
	0x0182 00386 (main.go:9)	CMPQ	R8, SI
	0x0185 00389 (main.go:9)	JGE	335
	0x0187 00391 (main.go:10)	MOVQ	R8, (CX)(DX*8)
	0x018b 00395 (main.go:10)	MOVQ	SI, 8(CX)(DX*8)
	0x0190 00400 (main.go:10)	JMP	335
	0x0192 00402 (main.go:9)	MOVQ	DI, AX
	0x0195 00405 (main.go:9)	MOVL	$7, CX
	0x019a 00410 (main.go:9)	CALL	runtime.panicIndex(SB)
	0x019f 00415 (main.go:9)	MOVQ	DX, AX
	0x01a2 00418 (main.go:9)	MOVL	$7, CX
	0x01a7 00423 (main.go:9)	CALL	runtime.panicIndex(SB)
	0x01ac 00428 (main.go:9)	XCHGL	AX, AX
	0x01ad 00429 (main.go:9)	NOP
	0x01ad 00429 (main.go:16)	PCDATA	$1, $-1
	0x01ad 00429 (main.go:16)	PCDATA	$0, $-2
	0x01ad 00429 (main.go:16)	CALL	runtime.morestack_noctxt(SB)
	0x01b2 00434 (main.go:16)	PCDATA	$0, $-1
	0x01b2 00434 (main.go:16)	JMP	0
	0x0000 49 3b 66 10 0f 86 a3 01 00 00 48 83 ec 78 48 89  I;f.......H..xH.
	0x0010 6c 24 70 48 8d 6c 24 70 48 8d 05 00 00 00 00 90  l$pH.l$pH.......
	0x0020 e8 00 00 00 00 48 89 44 24 28 48 c7 00 05 00 00  .....H.D$(H.....
	0x0030 00 48 c7 40 08 01 00 00 00 48 c7 40 10 05 00 00  .H.@.....H.@....
	0x0040 00 48 c7 40 18 0a 1a 00 00 48 c7 40 20 16 00 00  .H.@.....H.@ ...
	0x0050 00 48 c7 40 28 0b 00 00 00 48 c7 40 30 5a 00 00  .H.@(....H.@0Z..
	0x0060 00 44 0f 11 7c 24 50 44 0f 11 7c 24 60 48 8d 0d  .D..|$PD..|$`H..
	0x0070 00 00 00 00 48 89 4c 24 50 48 8d 15 00 00 00 00  ....H.L$PH......
	0x0080 48 89 54 24 58 bb 07 00 00 00 48 89 d9 e8 00 00  H.T$X.....H.....
	0x0090 00 00 48 8d 0d 00 00 00 00 48 89 4c 24 60 48 89  ..H......H.L$`H.
	0x00a0 44 24 68 48 8b 1d 00 00 00 00 48 8d 05 00 00 00  D$hH......H.....
	0x00b0 00 bf 02 00 00 00 48 89 fe 48 8d 4c 24 50 66 90  ......H..H.L$Pf.
	0x00c0 e8 00 00 00 00 90 48 8b 4c 24 28 31 c0 eb 03 48  ......H.L$(1...H
	0x00d0 ff c0 48 83 f8 06 7d 04 31 d2 eb 76 44 0f 11 7c  ..H...}.1..vD..|
	0x00e0 24 30 44 0f 11 7c 24 40 48 8d 15 00 00 00 00 48  $0D..|$@H......H
	0x00f0 89 54 24 30 48 8d 15 00 00 00 00 48 89 54 24 38  .T$0H......H.T$8
	0x0100 48 89 c8 bb 07 00 00 00 48 89 d9 e8 00 00 00 00  H.......H.......
	0x0110 48 8d 15 00 00 00 00 48 89 54 24 40 48 89 44 24  H......H.T$@H.D$
	0x0120 48 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00 48  HH......H......H
	0x0130 8d 4c 24 30 bf 02 00 00 00 48 89 fe 0f 1f 40 00  .L$0.....H....@.
	0x0140 e8 00 00 00 00 48 8b 6c 24 70 48 83 c4 78 c3 48  .....H.l$pH..x.H
	0x0150 89 fa 48 8d 70 fa 48 f7 de 0f 1f 80 00 00 00 00  ..H.p.H.........
	0x0160 48 39 f2 0f 8d 66 ff ff ff 48 83 fa 07 73 30 48  H9...f...H...s0H
	0x0170 8b 34 d1 48 8d 7a 01 48 83 ff 07 73 15 4c 8b 44  .4.H.z.H...s.L.D
	0x0180 d1 08 49 39 f0 7d c8 4c 89 04 d1 48 89 74 d1 08  ..I9.}.L...H.t..
	0x0190 eb bd 48 89 f8 b9 07 00 00 00 e8 00 00 00 00 48  ..H............H
	0x01a0 89 d0 b9 07 00 00 00 e8 00 00 00 00 90 e8 00 00  ................
	0x01b0 00 00 e9 49 fe ff ff                             ...I...
	rel 3+0 t=23 type.string+0
	rel 3+0 t=23 type.[]int+0
	rel 3+0 t=23 type.*os.File+0
	rel 3+0 t=23 type.string+0
	rel 3+0 t=23 type.[]int+0
	rel 3+0 t=23 type.*os.File+0
	rel 27+4 t=14 type.[7]int+0
	rel 33+4 t=7 runtime.newobject+0
	rel 112+4 t=14 type.string+0
	rel 124+4 t=14 ""..stmp_0+0
	rel 142+4 t=7 runtime.convTslice+0
	rel 149+4 t=14 type.[]int+0
	rel 166+4 t=14 os.Stdout+0
	rel 173+4 t=14 go.itab.*os.File,io.Writer+0
	rel 193+4 t=7 fmt.Fprintln+0
	rel 235+4 t=14 type.string+0
	rel 247+4 t=14 ""..stmp_1+0
	rel 268+4 t=7 runtime.convTslice+0
	rel 275+4 t=14 type.[]int+0
	rel 292+4 t=14 os.Stdout+0
	rel 299+4 t=14 go.itab.*os.File,io.Writer+0
	rel 321+4 t=7 fmt.Fprintln+0
	rel 411+4 t=7 runtime.panicIndex+0
	rel 424+4 t=7 runtime.panicIndex+0
	rel 430+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
go.info."".bubbleSort$abstract SDWARFABSFCN dupok size=50
	0x0000 05 2e 62 75 62 62 6c 65 53 6f 72 74 00 01 01 13  ..bubbleSort....
	0x0010 61 72 72 00 00 00 00 00 00 0e 6e 00 06 00 00 00  arr.......n.....
	0x0020 00 0e 69 00 07 00 00 00 00 0e 6a 00 08 00 00 00  ..i.......j.....
	0x0030 00 00                                            ..
	rel 0+0 t=22 type.[]int+0
	rel 0+0 t=22 type.int+0
	rel 21+4 t=31 go.info.[]int+0
	rel 29+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.int+0
	rel 45+4 t=31 go.info.int+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.string."Original array:" SRODATA dupok size=15
	0x0000 4f 72 69 67 69 6e 61 6c 20 61 72 72 61 79 3a     Original array:
go.string."Sorted array:  " SRODATA dupok size=15
	0x0000 53 6f 72 74 65 64 20 61 72 72 61 79 3a 20 20     Sorted array:  
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
""..stmp_0 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 0f 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."Original array:"+0
""..stmp_1 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 0f 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."Sorted array:  "+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]int- SRODATA dupok size=8
	0x0000 00 06 2a 5b 5d 69 6e 74                          ..*[]int
type.*[]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 44+4 t=-32763 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..eqfunc56 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 38 00 00 00 00 00 00 00  ........8.......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[7]int- SRODATA dupok size=9
	0x0000 00 07 2a 5b 37 5d 69 6e 74                       ..*[7]int
type.*[7]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 0d 1e 1e a7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[7]int-+0
	rel 48+8 t=1 type.[7]int+0
runtime.gcbits. SRODATA dupok size=0
type.[7]int SRODATA dupok size=72
	0x0000 38 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  8...............
	0x0010 ab 61 c0 7f 0a 08 08 11 00 00 00 00 00 00 00 00  .a..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 07 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc56+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[7]int-+0
	rel 44+4 t=-32763 type.*[7]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
	0x0000 0a                                               .
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
"".bubbleSort.arginfo1 SRODATA static dupok size=9
	0x0000 fe 00 08 08 08 10 08 fd ff                       .........
"".bubbleSort.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·f6bd6b3389b872033d462029172c8612 SRODATA dupok size=8
	0x0000 04 00 00 00 00 00 00 00                          ........
gclocals·804d9f4c78bfde951f5b4bad61ed1894 SRODATA dupok size=16
	0x0000 04 00 00 00 09 00 00 00 00 00 41 01 01 00 14 00  ..........A.....
"".main.stkobj SRODATA static size=40
	0x0000 02 00 00 00 00 00 00 00 c0 ff ff ff 20 00 00 00  ............ ...
	0x0010 20 00 00 00 00 00 00 00 e0 ff ff ff 20 00 00 00   ........... ...
	0x0020 20 00 00 00 00 00 00 00                           .......
	rel 20+4 t=5 runtime.gcbits.0a+0
	rel 36+4 t=5 runtime.gcbits.0a+0
