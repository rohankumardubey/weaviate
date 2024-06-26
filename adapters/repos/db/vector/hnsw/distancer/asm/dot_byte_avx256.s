//go:build !noasm && amd64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

TEXT ·dot_byte_256(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ res+16(FP), DX
	MOVQ len+24(FP), CX
	BYTE $0x55               // pushq	%rbp
	WORD $0x8948; BYTE $0xe5 // movq	%rsp, %rbp
	BYTE $0x53               // pushq	%rbx
	LONG $0xf8e48348         // andq	$-8, %rsp
	WORD $0x8b4c; BYTE $0x11 // movq	(%rcx), %r10
	WORD $0x8945; BYTE $0xd0 // movl	%r10d, %r8d
	LONG $0x20fa8341         // cmpl	$32, %r10d
	JGE  LBB0_1
	WORD $0x8545; BYTE $0xd2 // testl	%r10d, %r10d
	JLE  LBB0_7
	LONG $0x10f88341         // cmpl	$16, %r8d
	JAE  LBB0_10
	WORD $0xc031             // xorl	%eax, %eax
	WORD $0x3145; BYTE $0xc9 // xorl	%r9d, %r9d
	JMP  LBB0_13

LBB0_1:
	WORD $0x634d; BYTE $0xc8 // movslq	%r8d, %r9
	LONG $0xc0eff9c5         // vpxor	%xmm0, %xmm0, %xmm0
	WORD $0xdb31             // xorl	%ebx, %ebx
	WORD $0x894c; BYTE $0xc0 // movq	%r8, %rax

LBB0_2:
	WORD $0x8948; BYTE $0xd9       // movq	%rbx, %rcx
	LONG $0x0c6ffec5; BYTE $0x1f   // vmovdqu	(%rdi,%rbx), %ymm1
	LONG $0xe171edc5; BYTE $0x08   // vpsraw	$8, %ymm1, %ymm2
	LONG $0xf171f5c5; BYTE $0x08   // vpsllw	$8, %ymm1, %ymm1
	LONG $0xe171f5c5; BYTE $0x08   // vpsraw	$8, %ymm1, %ymm1
	LONG $0x1c6ffec5; BYTE $0x1e   // vmovdqu	(%rsi,%rbx), %ymm3
	LONG $0xe371ddc5; BYTE $0x08   // vpsraw	$8, %ymm3, %ymm4
	LONG $0xd4f5edc5               // vpmaddwd	%ymm4, %ymm2, %ymm2
	LONG $0xc0feedc5               // vpaddd	%ymm0, %ymm2, %ymm0
	LONG $0xf371edc5; BYTE $0x08   // vpsllw	$8, %ymm3, %ymm2
	LONG $0xe271edc5; BYTE $0x08   // vpsraw	$8, %ymm2, %ymm2
	LONG $0xcaf5f5c5               // vpmaddwd	%ymm2, %ymm1, %ymm1
	LONG $0xc1fefdc5               // vpaddd	%ymm1, %ymm0, %ymm0
	LONG $0x20c38348               // addq	$32, %rbx
	LONG $0xe0c08348               // addq	$-32, %rax
	LONG $0x3fc18348               // addq	$63, %rcx
	WORD $0x394c; BYTE $0xc9       // cmpq	%r9, %rcx
	JL   LBB0_2
	LONG $0x397de3c4; WORD $0x01c1 // vextracti128	$1, %ymm0, %xmm1
	LONG $0xc0fef1c5               // vpaddd	%xmm0, %xmm1, %xmm0
	LONG $0xc870f9c5; BYTE $0x1b   // vpshufd	$27, %xmm0, %xmm1               # xmm1 = xmm0[3,2,1,0]
	LONG $0xc0fef1c5               // vpaddd	%xmm0, %xmm1, %xmm0
	LONG $0xc870f9c5; BYTE $0x55   // vpshufd	$85, %xmm0, %xmm1               # xmm1 = xmm0[1,1,1,1]
	LONG $0xc0fef1c5               // vpaddd	%xmm0, %xmm1, %xmm0
	LONG $0x7e79c1c4; BYTE $0xc1   // vmovd	%xmm0, %r9d
	WORD $0x3944; BYTE $0xd3       // cmpl	%r10d, %ebx
	JGE  LBB0_18
	WORD $0x894d; BYTE $0xc2       // movq	%r8, %r10
	WORD $0x2949; BYTE $0xda       // subq	%rbx, %r10
	LONG $0x20fa8349               // cmpq	$32, %r10
	JAE  LBB0_14
	WORD $0x8948; BYTE $0xd8       // movq	%rbx, %rax
	JMP  LBB0_17

LBB0_7:
	WORD $0x3145; BYTE $0xc9 // xorl	%r9d, %r9d
	JMP  LBB0_18

LBB0_10:
	WORD $0x8945; BYTE $0xc2 // movl	%r8d, %r10d
	LONG $0x0fe28341         // andl	$15, %r10d
	WORD $0x8944; BYTE $0xc0 // movl	%r8d, %eax
	WORD $0xe083; BYTE $0xf0 // andl	$-16, %eax
	LONG $0xc0eff9c5         // vpxor	%xmm0, %xmm0, %xmm0
	WORD $0xc931             // xorl	%ecx, %ecx
	LONG $0xc9eff1c5         // vpxor	%xmm1, %xmm1, %xmm1
	LONG $0xd2efe9c5         // vpxor	%xmm2, %xmm2, %xmm2
	LONG $0xdbefe1c5         // vpxor	%xmm3, %xmm3, %xmm3

LBB0_11:
	LONG $0x327de2c4; WORD $0x0f24             // vpmovzxbq	(%rdi,%rcx), %ymm4      # ymm4 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x327de2c4; WORD $0x0f6c; BYTE $0x04 // vpmovzxbq	4(%rdi,%rcx), %ymm5     # ymm5 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x327de2c4; WORD $0x0f74; BYTE $0x08 // vpmovzxbq	8(%rdi,%rcx), %ymm6     # ymm6 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x327de2c4; WORD $0x0f7c; BYTE $0x0c // vpmovzxbq	12(%rdi,%rcx), %ymm7    # ymm7 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x327d62c4; WORD $0x0e04             // vpmovzxbq	(%rsi,%rcx), %ymm8      # ymm8 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x283de2c4; BYTE $0xe4               // vpmuldq	%ymm4, %ymm8, %ymm4
	LONG $0xc0d4ddc5                           // vpaddq	%ymm0, %ymm4, %ymm0
	LONG $0x327de2c4; WORD $0x0e64; BYTE $0x04 // vpmovzxbq	4(%rsi,%rcx), %ymm4     # ymm4 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x285de2c4; BYTE $0xe5               // vpmuldq	%ymm5, %ymm4, %ymm4
	LONG $0xc9d4ddc5                           // vpaddq	%ymm1, %ymm4, %ymm1
	LONG $0x327de2c4; WORD $0x0e64; BYTE $0x08 // vpmovzxbq	8(%rsi,%rcx), %ymm4     # ymm4 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x285de2c4; BYTE $0xe6               // vpmuldq	%ymm6, %ymm4, %ymm4
	LONG $0xd2d4ddc5                           // vpaddq	%ymm2, %ymm4, %ymm2
	LONG $0x327de2c4; WORD $0x0e64; BYTE $0x0c // vpmovzxbq	12(%rsi,%rcx), %ymm4    # ymm4 = mem[0],zero,zero,zero,zero,zero,zero,zero,mem[1],zero,zero,zero,zero,zero,zero,zero,mem[2],zero,zero,zero,zero,zero,zero,zero,mem[3],zero,zero,zero,zero,zero,zero,zero
	LONG $0x285de2c4; BYTE $0xe7               // vpmuldq	%ymm7, %ymm4, %ymm4
	LONG $0xdbd4ddc5                           // vpaddq	%ymm3, %ymm4, %ymm3
	LONG $0x10c18348                           // addq	$16, %rcx
	WORD $0x3948; BYTE $0xc8                   // cmpq	%rcx, %rax
	JNE  LBB0_11
	LONG $0xc0d4f5c5                           // vpaddq	%ymm0, %ymm1, %ymm0
	LONG $0xc0d4edc5                           // vpaddq	%ymm0, %ymm2, %ymm0
	LONG $0xc0d4e5c5                           // vpaddq	%ymm0, %ymm3, %ymm0
	LONG $0x397de3c4; WORD $0x01c1             // vextracti128	$1, %ymm0, %xmm1
	LONG $0xc1d4f9c5                           // vpaddq	%xmm1, %xmm0, %xmm0
	LONG $0xc870f9c5; BYTE $0xee               // vpshufd	$238, %xmm0, %xmm1              # xmm1 = xmm0[2,3,2,3]
	LONG $0xc1d4f9c5                           // vpaddq	%xmm1, %xmm0, %xmm0
	LONG $0x7ef9c1c4; BYTE $0xc1               // vmovq	%xmm0, %r9
	WORD $0x854d; BYTE $0xd2                   // testq	%r10, %r10
	JE   LBB0_18

LBB0_13:
	LONG $0x14b60f44; BYTE $0x07 // movzbl	(%rdi,%rax), %r10d
	LONG $0x060cb60f             // movzbl	(%rsi,%rax), %ecx
	LONG $0xcaaf0f49             // imulq	%r10, %rcx
	WORD $0x0149; BYTE $0xc9     // addq	%rcx, %r9
	LONG $0x01c08348             // addq	$1, %rax
	WORD $0x3949; BYTE $0xc0     // cmpq	%rax, %r8
	JNE  LBB0_13
	JMP  LBB0_18

LBB0_14:
	WORD $0x894d; BYTE $0xd3     // movq	%r10, %r11
	LONG $0xe0e38349             // andq	$-32, %r11
	LONG $0xe0e08348             // andq	$-32, %rax
	WORD $0x0148; BYTE $0xd8     // addq	%rbx, %rax
	LONG $0x18c38348             // addq	$24, %rbx
	LONG $0x6e79c1c4; BYTE $0xc1 // vmovd	%r9d, %xmm0
	LONG $0xc9eff1c5             // vpxor	%xmm1, %xmm1, %xmm1
	WORD $0x894d; BYTE $0xd9     // movq	%r11, %r9
	LONG $0xd2efe9c5             // vpxor	%xmm2, %xmm2, %xmm2
	LONG $0xdbefe1c5             // vpxor	%xmm3, %xmm3, %xmm3

LBB0_15:
	LONG $0x317de2c4; WORD $0x1f64; BYTE $0xe8 // vpmovzxbd	-24(%rdi,%rbx), %ymm4   # ymm4 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0x317de2c4; WORD $0x1f6c; BYTE $0xf0 // vpmovzxbd	-16(%rdi,%rbx), %ymm5   # ymm5 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0x317de2c4; WORD $0x1f74; BYTE $0xf8 // vpmovzxbd	-8(%rdi,%rbx), %ymm6    # ymm6 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0x317de2c4; WORD $0x1f3c             // vpmovzxbd	(%rdi,%rbx), %ymm7      # ymm7 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0x317d62c4; WORD $0x1e44; BYTE $0xe8 // vpmovzxbd	-24(%rsi,%rbx), %ymm8   # ymm8 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0xe4f5bdc5                           // vpmaddwd	%ymm4, %ymm8, %ymm4
	LONG $0xc0feddc5                           // vpaddd	%ymm0, %ymm4, %ymm0
	LONG $0x317de2c4; WORD $0x1e64; BYTE $0xf0 // vpmovzxbd	-16(%rsi,%rbx), %ymm4   # ymm4 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0xe5f5ddc5                           // vpmaddwd	%ymm5, %ymm4, %ymm4
	LONG $0xc9feddc5                           // vpaddd	%ymm1, %ymm4, %ymm1
	LONG $0x317de2c4; WORD $0x1e64; BYTE $0xf8 // vpmovzxbd	-8(%rsi,%rbx), %ymm4    # ymm4 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0xe6f5ddc5                           // vpmaddwd	%ymm6, %ymm4, %ymm4
	LONG $0xd2feddc5                           // vpaddd	%ymm2, %ymm4, %ymm2
	LONG $0x317de2c4; WORD $0x1e24             // vpmovzxbd	(%rsi,%rbx), %ymm4      # ymm4 = mem[0],zero,zero,zero,mem[1],zero,zero,zero,mem[2],zero,zero,zero,mem[3],zero,zero,zero,mem[4],zero,zero,zero,mem[5],zero,zero,zero,mem[6],zero,zero,zero,mem[7],zero,zero,zero
	LONG $0xe7f5ddc5                           // vpmaddwd	%ymm7, %ymm4, %ymm4
	LONG $0xdbfeddc5                           // vpaddd	%ymm3, %ymm4, %ymm3
	LONG $0x20c38348                           // addq	$32, %rbx
	LONG $0xe0c18349                           // addq	$-32, %r9
	JNE  LBB0_15
	LONG $0xc0fef5c5                           // vpaddd	%ymm0, %ymm1, %ymm0
	LONG $0xc0feedc5                           // vpaddd	%ymm0, %ymm2, %ymm0
	LONG $0xc0fee5c5                           // vpaddd	%ymm0, %ymm3, %ymm0
	LONG $0x397de3c4; WORD $0x01c1             // vextracti128	$1, %ymm0, %xmm1
	LONG $0xc1fef9c5                           // vpaddd	%xmm1, %xmm0, %xmm0
	LONG $0xc870f9c5; BYTE $0xee               // vpshufd	$238, %xmm0, %xmm1              # xmm1 = xmm0[2,3,2,3]
	LONG $0xc1fef9c5                           // vpaddd	%xmm1, %xmm0, %xmm0
	LONG $0xc870f9c5; BYTE $0x55               // vpshufd	$85, %xmm0, %xmm1               # xmm1 = xmm0[1,1,1,1]
	LONG $0xc1fef9c5                           // vpaddd	%xmm1, %xmm0, %xmm0
	LONG $0x7e79c1c4; BYTE $0xc1               // vmovd	%xmm0, %r9d
	WORD $0x394d; BYTE $0xda                   // cmpq	%r11, %r10
	JE   LBB0_18

LBB0_17:
	LONG $0x070cb60f         // movzbl	(%rdi,%rax), %ecx
	LONG $0x061cb60f         // movzbl	(%rsi,%rax), %ebx
	WORD $0xaf0f; BYTE $0xd9 // imull	%ecx, %ebx
	WORD $0x0141; BYTE $0xd9 // addl	%ebx, %r9d
	LONG $0x01c08348         // addq	$1, %rax
	WORD $0x3949; BYTE $0xc0 // cmpq	%rax, %r8
	JNE  LBB0_17

LBB0_18:
	WORD $0x8944; BYTE $0x0a // movl	%r9d, (%rdx)
	LONG $0xf8658d48         // leaq	-8(%rbp), %rsp
	BYTE $0x5b               // popq	%rbx
	BYTE $0x5d               // popq	%rbp
	WORD $0xf8c5; BYTE $0x77 // vzeroupper
	BYTE $0xc3               // retq
