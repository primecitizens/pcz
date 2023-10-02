// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elf

// Indexes into the Header.Ident array.
const (
	EI_CLASS      = 4  /* Class of machine. */
	EI_DATA       = 5  /* Data format. */
	EI_VERSION    = 6  /* ELF format version. */
	EI_OSABI      = 7  /* Operating system / ABI identification */
	EI_ABIVERSION = 8  /* ABI version */
	EI_PAD        = 9  /* Start of padding (per SVR4 ABI). */
	EI_NIDENT     = 16 /* Size of e_ident array. */
)

// Initial magic number for ELF files.
const ELFMAG = "\177ELF"

// Version is found in Header.Ident[EI_VERSION] and Header.Version.
type Version byte

const (
	EV_NONE    Version = 0
	EV_CURRENT Version = 1
)

// Class is found in Header.Ident[EI_CLASS] and Header.Class.
type Class byte

const (
	ELFCLASSNONE Class = 0 /* Unknown class. */
	ELFCLASS32   Class = 1 /* 32-bit architecture. */
	ELFCLASS64   Class = 2 /* 64-bit architecture. */
)

// Data is found in Header.Ident[EI_DATA] and Header.Data.
type Data byte

const (
	ELFDATANONE Data = 0 /* Unknown data format. */
	ELFDATA2LSB Data = 1 /* 2's complement little-endian. */
	ELFDATA2MSB Data = 2 /* 2's complement big-endian. */
)

// OSABI is found in Header.Ident[EI_OSABI] and Header.OSABI.
type OSABI byte

const (
	ELFOSABI_NONE       OSABI = 0   /* UNIX System V ABI */
	ELFOSABI_HPUX       OSABI = 1   /* HP-UX operating system */
	ELFOSABI_NETBSD     OSABI = 2   /* NetBSD */
	ELFOSABI_LINUX      OSABI = 3   /* Linux */
	ELFOSABI_HURD       OSABI = 4   /* Hurd */
	ELFOSABI_86OPEN     OSABI = 5   /* 86Open common IA32 ABI */
	ELFOSABI_SOLARIS    OSABI = 6   /* Solaris */
	ELFOSABI_AIX        OSABI = 7   /* AIX */
	ELFOSABI_IRIX       OSABI = 8   /* IRIX */
	ELFOSABI_FREEBSD    OSABI = 9   /* FreeBSD */
	ELFOSABI_TRU64      OSABI = 10  /* TRU64 UNIX */
	ELFOSABI_MODESTO    OSABI = 11  /* Novell Modesto */
	ELFOSABI_OPENBSD    OSABI = 12  /* OpenBSD */
	ELFOSABI_OPENVMS    OSABI = 13  /* Open VMS */
	ELFOSABI_NSK        OSABI = 14  /* HP Non-Stop Kernel */
	ELFOSABI_AROS       OSABI = 15  /* Amiga Research OS */
	ELFOSABI_FENIXOS    OSABI = 16  /* The FenixOS highly scalable multi-core OS */
	ELFOSABI_CLOUDABI   OSABI = 17  /* Nuxi CloudABI */
	ELFOSABI_ARM        OSABI = 97  /* ARM */
	ELFOSABI_STANDALONE OSABI = 255 /* Standalone (embedded) application */
)

// Type is found in Header.Type.
type Type uint16

const (
	ET_NONE   Type = 0      /* Unknown type. */
	ET_REL    Type = 1      /* Relocatable. */
	ET_EXEC   Type = 2      /* Executable. */
	ET_DYN    Type = 3      /* Shared object. */
	ET_CORE   Type = 4      /* Core file. */
	ET_LOOS   Type = 0xfe00 /* First operating system specific. */
	ET_HIOS   Type = 0xfeff /* Last operating system-specific. */
	ET_LOPROC Type = 0xff00 /* First processor-specific. */
	ET_HIPROC Type = 0xffff /* Last processor-specific. */
)

// Machine is found in Header.Machine.
type Machine uint16

const (
	EM_NONE          Machine = 0   /* Unknown machine. */
	EM_M32           Machine = 1   /* AT&T WE32100. */
	EM_SPARC         Machine = 2   /* Sun SPARC. */
	EM_386           Machine = 3   /* Intel i386. */
	EM_68K           Machine = 4   /* Motorola 68000. */
	EM_88K           Machine = 5   /* Motorola 88000. */
	EM_860           Machine = 7   /* Intel i860. */
	EM_MIPS          Machine = 8   /* MIPS R3000 Big-Endian only. */
	EM_S370          Machine = 9   /* IBM System/370. */
	EM_MIPS_RS3_LE   Machine = 10  /* MIPS R3000 Little-Endian. */
	EM_PARISC        Machine = 15  /* HP PA-RISC. */
	EM_VPP500        Machine = 17  /* Fujitsu VPP500. */
	EM_SPARC32PLUS   Machine = 18  /* SPARC v8plus. */
	EM_960           Machine = 19  /* Intel 80960. */
	EM_PPC           Machine = 20  /* PowerPC 32-bit. */
	EM_PPC64         Machine = 21  /* PowerPC 64-bit. */
	EM_S390          Machine = 22  /* IBM System/390. */
	EM_V800          Machine = 36  /* NEC V800. */
	EM_FR20          Machine = 37  /* Fujitsu FR20. */
	EM_RH32          Machine = 38  /* TRW RH-32. */
	EM_RCE           Machine = 39  /* Motorola RCE. */
	EM_ARM           Machine = 40  /* ARM. */
	EM_SH            Machine = 42  /* Hitachi SH. */
	EM_SPARCV9       Machine = 43  /* SPARC v9 64-bit. */
	EM_TRICORE       Machine = 44  /* Siemens TriCore embedded processor. */
	EM_ARC           Machine = 45  /* Argonaut RISC Core. */
	EM_H8_300        Machine = 46  /* Hitachi H8/300. */
	EM_H8_300H       Machine = 47  /* Hitachi H8/300H. */
	EM_H8S           Machine = 48  /* Hitachi H8S. */
	EM_H8_500        Machine = 49  /* Hitachi H8/500. */
	EM_IA_64         Machine = 50  /* Intel IA-64 Processor. */
	EM_MIPS_X        Machine = 51  /* Stanford MIPS-X. */
	EM_COLDFIRE      Machine = 52  /* Motorola ColdFire. */
	EM_68HC12        Machine = 53  /* Motorola M68HC12. */
	EM_MMA           Machine = 54  /* Fujitsu MMA. */
	EM_PCP           Machine = 55  /* Siemens PCP. */
	EM_NCPU          Machine = 56  /* Sony nCPU. */
	EM_NDR1          Machine = 57  /* Denso NDR1 microprocessor. */
	EM_STARCORE      Machine = 58  /* Motorola Star*Core processor. */
	EM_ME16          Machine = 59  /* Toyota ME16 processor. */
	EM_ST100         Machine = 60  /* STMicroelectronics ST100 processor. */
	EM_TINYJ         Machine = 61  /* Advanced Logic Corp. TinyJ processor. */
	EM_X86_64        Machine = 62  /* Advanced Micro Devices x86-64 */
	EM_PDSP          Machine = 63  /* Sony DSP Processor */
	EM_PDP10         Machine = 64  /* Digital Equipment Corp. PDP-10 */
	EM_PDP11         Machine = 65  /* Digital Equipment Corp. PDP-11 */
	EM_FX66          Machine = 66  /* Siemens FX66 microcontroller */
	EM_ST9PLUS       Machine = 67  /* STMicroelectronics ST9+ 8/16 bit microcontroller */
	EM_ST7           Machine = 68  /* STMicroelectronics ST7 8-bit microcontroller */
	EM_68HC16        Machine = 69  /* Motorola MC68HC16 Microcontroller */
	EM_68HC11        Machine = 70  /* Motorola MC68HC11 Microcontroller */
	EM_68HC08        Machine = 71  /* Motorola MC68HC08 Microcontroller */
	EM_68HC05        Machine = 72  /* Motorola MC68HC05 Microcontroller */
	EM_SVX           Machine = 73  /* Silicon Graphics SVx */
	EM_ST19          Machine = 74  /* STMicroelectronics ST19 8-bit microcontroller */
	EM_VAX           Machine = 75  /* Digital VAX */
	EM_CRIS          Machine = 76  /* Axis Communications 32-bit embedded processor */
	EM_JAVELIN       Machine = 77  /* Infineon Technologies 32-bit embedded processor */
	EM_FIREPATH      Machine = 78  /* Element 14 64-bit DSP Processor */
	EM_ZSP           Machine = 79  /* LSI Logic 16-bit DSP Processor */
	EM_MMIX          Machine = 80  /* Donald Knuth's educational 64-bit processor */
	EM_HUANY         Machine = 81  /* Harvard University machine-independent object files */
	EM_PRISM         Machine = 82  /* SiTera Prism */
	EM_AVR           Machine = 83  /* Atmel AVR 8-bit microcontroller */
	EM_FR30          Machine = 84  /* Fujitsu FR30 */
	EM_D10V          Machine = 85  /* Mitsubishi D10V */
	EM_D30V          Machine = 86  /* Mitsubishi D30V */
	EM_V850          Machine = 87  /* NEC v850 */
	EM_M32R          Machine = 88  /* Mitsubishi M32R */
	EM_MN10300       Machine = 89  /* Matsushita MN10300 */
	EM_MN10200       Machine = 90  /* Matsushita MN10200 */
	EM_PJ            Machine = 91  /* picoJava */
	EM_OPENRISC      Machine = 92  /* OpenRISC 32-bit embedded processor */
	EM_ARC_COMPACT   Machine = 93  /* ARC International ARCompact processor (old spelling/synonym: EM_ARC_A5) */
	EM_XTENSA        Machine = 94  /* Tensilica Xtensa Architecture */
	EM_VIDEOCORE     Machine = 95  /* Alphamosaic VideoCore processor */
	EM_TMM_GPP       Machine = 96  /* Thompson Multimedia General Purpose Processor */
	EM_NS32K         Machine = 97  /* National Semiconductor 32000 series */
	EM_TPC           Machine = 98  /* Tenor Network TPC processor */
	EM_SNP1K         Machine = 99  /* Trebia SNP 1000 processor */
	EM_ST200         Machine = 100 /* STMicroelectronics (www.st.com) ST200 microcontroller */
	EM_IP2K          Machine = 101 /* Ubicom IP2xxx microcontroller family */
	EM_MAX           Machine = 102 /* MAX Processor */
	EM_CR            Machine = 103 /* National Semiconductor CompactRISC microprocessor */
	EM_F2MC16        Machine = 104 /* Fujitsu F2MC16 */
	EM_MSP430        Machine = 105 /* Texas Instruments embedded microcontroller msp430 */
	EM_BLACKFIN      Machine = 106 /* Analog Devices Blackfin (DSP) processor */
	EM_SE_C33        Machine = 107 /* S1C33 Family of Seiko Epson processors */
	EM_SEP           Machine = 108 /* Sharp embedded microprocessor */
	EM_ARCA          Machine = 109 /* Arca RISC Microprocessor */
	EM_UNICORE       Machine = 110 /* Microprocessor series from PKU-Unity Ltd. and MPRC of Peking University */
	EM_EXCESS        Machine = 111 /* eXcess: 16/32/64-bit configurable embedded CPU */
	EM_DXP           Machine = 112 /* Icera Semiconductor Inc. Deep Execution Processor */
	EM_ALTERA_NIOS2  Machine = 113 /* Altera Nios II soft-core processor */
	EM_CRX           Machine = 114 /* National Semiconductor CompactRISC CRX microprocessor */
	EM_XGATE         Machine = 115 /* Motorola XGATE embedded processor */
	EM_C166          Machine = 116 /* Infineon C16x/XC16x processor */
	EM_M16C          Machine = 117 /* Renesas M16C series microprocessors */
	EM_DSPIC30F      Machine = 118 /* Microchip Technology dsPIC30F Digital Signal Controller */
	EM_CE            Machine = 119 /* Freescale Communication Engine RISC core */
	EM_M32C          Machine = 120 /* Renesas M32C series microprocessors */
	EM_TSK3000       Machine = 131 /* Altium TSK3000 core */
	EM_RS08          Machine = 132 /* Freescale RS08 embedded processor */
	EM_SHARC         Machine = 133 /* Analog Devices SHARC family of 32-bit DSP processors */
	EM_ECOG2         Machine = 134 /* Cyan Technology eCOG2 microprocessor */
	EM_SCORE7        Machine = 135 /* Sunplus S+core7 RISC processor */
	EM_DSP24         Machine = 136 /* New Japan Radio (NJR) 24-bit DSP Processor */
	EM_VIDEOCORE3    Machine = 137 /* Broadcom VideoCore III processor */
	EM_LATTICEMICO32 Machine = 138 /* RISC processor for Lattice FPGA architecture */
	EM_SE_C17        Machine = 139 /* Seiko Epson C17 family */
	EM_TI_C6000      Machine = 140 /* The Texas Instruments TMS320C6000 DSP family */
	EM_TI_C2000      Machine = 141 /* The Texas Instruments TMS320C2000 DSP family */
	EM_TI_C5500      Machine = 142 /* The Texas Instruments TMS320C55x DSP family */
	EM_TI_ARP32      Machine = 143 /* Texas Instruments Application Specific RISC Processor, 32bit fetch */
	EM_TI_PRU        Machine = 144 /* Texas Instruments Programmable Realtime Unit */
	EM_MMDSP_PLUS    Machine = 160 /* STMicroelectronics 64bit VLIW Data Signal Processor */
	EM_CYPRESS_M8C   Machine = 161 /* Cypress M8C microprocessor */
	EM_R32C          Machine = 162 /* Renesas R32C series microprocessors */
	EM_TRIMEDIA      Machine = 163 /* NXP Semiconductors TriMedia architecture family */
	EM_QDSP6         Machine = 164 /* QUALCOMM DSP6 Processor */
	EM_8051          Machine = 165 /* Intel 8051 and variants */
	EM_STXP7X        Machine = 166 /* STMicroelectronics STxP7x family of configurable and extensible RISC processors */
	EM_NDS32         Machine = 167 /* Andes Technology compact code size embedded RISC processor family */
	EM_ECOG1         Machine = 168 /* Cyan Technology eCOG1X family */
	EM_ECOG1X        Machine = 168 /* Cyan Technology eCOG1X family */
	EM_MAXQ30        Machine = 169 /* Dallas Semiconductor MAXQ30 Core Micro-controllers */
	EM_XIMO16        Machine = 170 /* New Japan Radio (NJR) 16-bit DSP Processor */
	EM_MANIK         Machine = 171 /* M2000 Reconfigurable RISC Microprocessor */
	EM_CRAYNV2       Machine = 172 /* Cray Inc. NV2 vector architecture */
	EM_RX            Machine = 173 /* Renesas RX family */
	EM_METAG         Machine = 174 /* Imagination Technologies META processor architecture */
	EM_MCST_ELBRUS   Machine = 175 /* MCST Elbrus general purpose hardware architecture */
	EM_ECOG16        Machine = 176 /* Cyan Technology eCOG16 family */
	EM_CR16          Machine = 177 /* National Semiconductor CompactRISC CR16 16-bit microprocessor */
	EM_ETPU          Machine = 178 /* Freescale Extended Time Processing Unit */
	EM_SLE9X         Machine = 179 /* Infineon Technologies SLE9X core */
	EM_L10M          Machine = 180 /* Intel L10M */
	EM_K10M          Machine = 181 /* Intel K10M */
	EM_AARCH64       Machine = 183 /* ARM 64-bit Architecture (AArch64) */
	EM_AVR32         Machine = 185 /* Atmel Corporation 32-bit microprocessor family */
	EM_STM8          Machine = 186 /* STMicroeletronics STM8 8-bit microcontroller */
	EM_TILE64        Machine = 187 /* Tilera TILE64 multicore architecture family */
	EM_TILEPRO       Machine = 188 /* Tilera TILEPro multicore architecture family */
	EM_MICROBLAZE    Machine = 189 /* Xilinx MicroBlaze 32-bit RISC soft processor core */
	EM_CUDA          Machine = 190 /* NVIDIA CUDA architecture */
	EM_TILEGX        Machine = 191 /* Tilera TILE-Gx multicore architecture family */
	EM_CLOUDSHIELD   Machine = 192 /* CloudShield architecture family */
	EM_COREA_1ST     Machine = 193 /* KIPO-KAIST Core-A 1st generation processor family */
	EM_COREA_2ND     Machine = 194 /* KIPO-KAIST Core-A 2nd generation processor family */
	EM_ARC_COMPACT2  Machine = 195 /* Synopsys ARCompact V2 */
	EM_OPEN8         Machine = 196 /* Open8 8-bit RISC soft processor core */
	EM_RL78          Machine = 197 /* Renesas RL78 family */
	EM_VIDEOCORE5    Machine = 198 /* Broadcom VideoCore V processor */
	EM_78KOR         Machine = 199 /* Renesas 78KOR family */
	EM_56800EX       Machine = 200 /* Freescale 56800EX Digital Signal Controller (DSC) */
	EM_BA1           Machine = 201 /* Beyond BA1 CPU architecture */
	EM_BA2           Machine = 202 /* Beyond BA2 CPU architecture */
	EM_XCORE         Machine = 203 /* XMOS xCORE processor family */
	EM_MCHP_PIC      Machine = 204 /* Microchip 8-bit PIC(r) family */
	EM_INTEL205      Machine = 205 /* Reserved by Intel */
	EM_INTEL206      Machine = 206 /* Reserved by Intel */
	EM_INTEL207      Machine = 207 /* Reserved by Intel */
	EM_INTEL208      Machine = 208 /* Reserved by Intel */
	EM_INTEL209      Machine = 209 /* Reserved by Intel */
	EM_KM32          Machine = 210 /* KM211 KM32 32-bit processor */
	EM_KMX32         Machine = 211 /* KM211 KMX32 32-bit processor */
	EM_KMX16         Machine = 212 /* KM211 KMX16 16-bit processor */
	EM_KMX8          Machine = 213 /* KM211 KMX8 8-bit processor */
	EM_KVARC         Machine = 214 /* KM211 KVARC processor */
	EM_CDP           Machine = 215 /* Paneve CDP architecture family */
	EM_COGE          Machine = 216 /* Cognitive Smart Memory Processor */
	EM_COOL          Machine = 217 /* Bluechip Systems CoolEngine */
	EM_NORC          Machine = 218 /* Nanoradio Optimized RISC */
	EM_CSR_KALIMBA   Machine = 219 /* CSR Kalimba architecture family */
	EM_Z80           Machine = 220 /* Zilog Z80 */
	EM_VISIUM        Machine = 221 /* Controls and Data Services VISIUMcore processor */
	EM_FT32          Machine = 222 /* FTDI Chip FT32 high performance 32-bit RISC architecture */
	EM_MOXIE         Machine = 223 /* Moxie processor family */
	EM_AMDGPU        Machine = 224 /* AMD GPU architecture */
	EM_RISCV         Machine = 243 /* RISC-V */
	EM_LANAI         Machine = 244 /* Lanai 32-bit processor */
	EM_BPF           Machine = 247 /* Linux BPF – in-kernel virtual machine */

	/* Non-standard or deprecated. */
	EM_486         Machine = 6      /* Intel i486. */
	EM_MIPS_RS4_BE Machine = 10     /* MIPS R4000 Big-Endian */
	EM_ALPHA_STD   Machine = 41     /* Digital Alpha (standard value). */
	EM_ALPHA       Machine = 0x9026 /* Alpha (written in the absence of an ABI) */
)

// Special section indices.
type SectionIndex int

const (
	SHN_UNDEF     SectionIndex = 0      /* Undefined, missing, irrelevant. */
	SHN_LORESERVE SectionIndex = 0xff00 /* First of reserved range. */
	SHN_LOPROC    SectionIndex = 0xff00 /* First processor-specific. */
	SHN_HIPROC    SectionIndex = 0xff1f /* Last processor-specific. */
	SHN_LOOS      SectionIndex = 0xff20 /* First operating system-specific. */
	SHN_HIOS      SectionIndex = 0xff3f /* Last operating system-specific. */
	SHN_ABS       SectionIndex = 0xfff1 /* Absolute values. */
	SHN_COMMON    SectionIndex = 0xfff2 /* Common data. */
	SHN_XINDEX    SectionIndex = 0xffff /* Escape; index stored elsewhere. */
	SHN_HIRESERVE SectionIndex = 0xffff /* Last of reserved range. */
)

// Section type.
type SectionType uint32

const (
	SHT_NULL           SectionType = 0          /* inactive */
	SHT_PROGBITS       SectionType = 1          /* program defined information */
	SHT_SYMTAB         SectionType = 2          /* symbol table section */
	SHT_STRTAB         SectionType = 3          /* string table section */
	SHT_RELA           SectionType = 4          /* relocation section with addends */
	SHT_HASH           SectionType = 5          /* symbol hash table section */
	SHT_DYNAMIC        SectionType = 6          /* dynamic section */
	SHT_NOTE           SectionType = 7          /* note section */
	SHT_NOBITS         SectionType = 8          /* no space section */
	SHT_REL            SectionType = 9          /* relocation section - no addends */
	SHT_SHLIB          SectionType = 10         /* reserved - purpose unknown */
	SHT_DYNSYM         SectionType = 11         /* dynamic symbol table section */
	SHT_INIT_ARRAY     SectionType = 14         /* Initialization function pointers. */
	SHT_FINI_ARRAY     SectionType = 15         /* Termination function pointers. */
	SHT_PREINIT_ARRAY  SectionType = 16         /* Pre-initialization function ptrs. */
	SHT_GROUP          SectionType = 17         /* Section group. */
	SHT_SYMTAB_SHNDX   SectionType = 18         /* Section indexes (see SHN_XINDEX). */
	SHT_LOOS           SectionType = 0x60000000 /* First of OS specific semantics */
	SHT_GNU_ATTRIBUTES SectionType = 0x6ffffff5 /* GNU object attributes */
	SHT_GNU_HASH       SectionType = 0x6ffffff6 /* GNU hash table */
	SHT_GNU_LIBLIST    SectionType = 0x6ffffff7 /* GNU prelink library list */
	SHT_GNU_VERDEF     SectionType = 0x6ffffffd /* GNU version definition section */
	SHT_GNU_VERNEED    SectionType = 0x6ffffffe /* GNU version needs section */
	SHT_GNU_VERSYM     SectionType = 0x6fffffff /* GNU version symbol table */
	SHT_HIOS           SectionType = 0x6fffffff /* Last of OS specific semantics */
	SHT_LOPROC         SectionType = 0x70000000 /* reserved range for processor */
	SHT_MIPS_ABIFLAGS  SectionType = 0x7000002a /* .MIPS.abiflags */
	SHT_HIPROC         SectionType = 0x7fffffff /* specific section header types */
	SHT_LOUSER         SectionType = 0x80000000 /* reserved range for application */
	SHT_HIUSER         SectionType = 0xffffffff /* specific indexes */
)

// Section flags.
type SectionFlag uint32

const (
	SHF_WRITE            SectionFlag = 0x1        /* Section contains writable data. */
	SHF_ALLOC            SectionFlag = 0x2        /* Section occupies memory. */
	SHF_EXECINSTR        SectionFlag = 0x4        /* Section contains instructions. */
	SHF_MERGE            SectionFlag = 0x10       /* Section may be merged. */
	SHF_STRINGS          SectionFlag = 0x20       /* Section contains strings. */
	SHF_INFO_LINK        SectionFlag = 0x40       /* sh_info holds section index. */
	SHF_LINK_ORDER       SectionFlag = 0x80       /* Special ordering requirements. */
	SHF_OS_NONCONFORMING SectionFlag = 0x100      /* OS-specific processing required. */
	SHF_GROUP            SectionFlag = 0x200      /* Member of section group. */
	SHF_TLS              SectionFlag = 0x400      /* Section contains TLS data. */
	SHF_COMPRESSED       SectionFlag = 0x800      /* Section is compressed. */
	SHF_MASKOS           SectionFlag = 0x0ff00000 /* OS-specific semantics. */
	SHF_MASKPROC         SectionFlag = 0xf0000000 /* Processor-specific semantics. */
)

// Section compression type.
type CompressionType int

const (
	COMPRESS_ZLIB   CompressionType = 1          /* ZLIB compression. */
	COMPRESS_ZSTD   CompressionType = 2          /* ZSTD compression. */
	COMPRESS_LOOS   CompressionType = 0x60000000 /* First OS-specific. */
	COMPRESS_HIOS   CompressionType = 0x6fffffff /* Last OS-specific. */
	COMPRESS_LOPROC CompressionType = 0x70000000 /* First processor-specific type. */
	COMPRESS_HIPROC CompressionType = 0x7fffffff /* Last processor-specific type. */
)

// Prog.Type
type ProgType int

const (
	PT_NULL    ProgType = 0 /* Unused entry. */
	PT_LOAD    ProgType = 1 /* Loadable segment. */
	PT_DYNAMIC ProgType = 2 /* Dynamic linking information segment. */
	PT_INTERP  ProgType = 3 /* Pathname of interpreter. */
	PT_NOTE    ProgType = 4 /* Auxiliary information. */
	PT_SHLIB   ProgType = 5 /* Reserved (not used). */
	PT_PHDR    ProgType = 6 /* Location of program header itself. */
	PT_TLS     ProgType = 7 /* Thread local storage segment */

	PT_LOOS ProgType = 0x60000000 /* First OS-specific. */

	PT_GNU_EH_FRAME ProgType = 0x6474e550 /* Frame unwind information */
	PT_GNU_STACK    ProgType = 0x6474e551 /* Stack flags */
	PT_GNU_RELRO    ProgType = 0x6474e552 /* Read only after relocs */
	PT_GNU_PROPERTY ProgType = 0x6474e553 /* GNU property */
	PT_GNU_MBIND_LO ProgType = 0x6474e555 /* Mbind segments start */
	PT_GNU_MBIND_HI ProgType = 0x6474f554 /* Mbind segments finish */

	PT_PAX_FLAGS ProgType = 0x65041580 /* PAX flags */

	PT_OPENBSD_RANDOMIZE ProgType = 0x65a3dbe6 /* Random data */
	PT_OPENBSD_WXNEEDED  ProgType = 0x65a3dbe7 /* W^X violations */
	PT_OPENBSD_BOOTDATA  ProgType = 0x65a41be6 /* Boot arguments */

	PT_SUNW_EH_FRAME ProgType = 0x6474e550 /* Frame unwind information */
	PT_SUNWSTACK     ProgType = 0x6ffffffb /* Stack segment */

	PT_HIOS ProgType = 0x6fffffff /* Last OS-specific. */

	PT_LOPROC ProgType = 0x70000000 /* First processor-specific type. */

	PT_ARM_ARCHEXT ProgType = 0x70000000 /* Architecture compatibility */
	PT_ARM_EXIDX   ProgType = 0x70000001 /* Exception unwind tables */

	PT_AARCH64_ARCHEXT ProgType = 0x70000000 /* Architecture compatibility */
	PT_AARCH64_UNWIND  ProgType = 0x70000001 /* Exception unwind tables */

	PT_MIPS_REGINFO  ProgType = 0x70000000 /* Register usage */
	PT_MIPS_RTPROC   ProgType = 0x70000001 /* Runtime procedures */
	PT_MIPS_OPTIONS  ProgType = 0x70000002 /* Options */
	PT_MIPS_ABIFLAGS ProgType = 0x70000003 /* ABI flags */

	PT_S390_PGSTE ProgType = 0x70000000 /* 4k page table size */

	PT_HIPROC ProgType = 0x7fffffff /* Last processor-specific type. */
)

// Prog.Flag
type ProgFlag uint32

const (
	PF_X        ProgFlag = 0x1        /* Executable. */
	PF_W        ProgFlag = 0x2        /* Writable. */
	PF_R        ProgFlag = 0x4        /* Readable. */
	PF_MASKOS   ProgFlag = 0x0ff00000 /* Operating system-specific. */
	PF_MASKPROC ProgFlag = 0xf0000000 /* Processor-specific. */
)

// Dyn.Tag
type DynTag int

const (
	DT_NULL         DynTag = 0  /* Terminating entry. */
	DT_NEEDED       DynTag = 1  /* String table offset of a needed shared library. */
	DT_PLTRELSZ     DynTag = 2  /* Total size in bytes of PLT relocations. */
	DT_PLTGOT       DynTag = 3  /* Processor-dependent address. */
	DT_HASH         DynTag = 4  /* Address of symbol hash table. */
	DT_STRTAB       DynTag = 5  /* Address of string table. */
	DT_SYMTAB       DynTag = 6  /* Address of symbol table. */
	DT_RELA         DynTag = 7  /* Address of ElfNN_Rela relocations. */
	DT_RELASZ       DynTag = 8  /* Total size of ElfNN_Rela relocations. */
	DT_RELAENT      DynTag = 9  /* Size of each ElfNN_Rela relocation entry. */
	DT_STRSZ        DynTag = 10 /* Size of string table. */
	DT_SYMENT       DynTag = 11 /* Size of each symbol table entry. */
	DT_INIT         DynTag = 12 /* Address of initialization function. */
	DT_FINI         DynTag = 13 /* Address of finalization function. */
	DT_SONAME       DynTag = 14 /* String table offset of shared object name. */
	DT_RPATH        DynTag = 15 /* String table offset of library path. [sup] */
	DT_SYMBOLIC     DynTag = 16 /* Indicates "symbolic" linking. [sup] */
	DT_REL          DynTag = 17 /* Address of ElfNN_Rel relocations. */
	DT_RELSZ        DynTag = 18 /* Total size of ElfNN_Rel relocations. */
	DT_RELENT       DynTag = 19 /* Size of each ElfNN_Rel relocation. */
	DT_PLTREL       DynTag = 20 /* Type of relocation used for PLT. */
	DT_DEBUG        DynTag = 21 /* Reserved (not used). */
	DT_TEXTREL      DynTag = 22 /* Indicates there may be relocations in non-writable segments. [sup] */
	DT_JMPREL       DynTag = 23 /* Address of PLT relocations. */
	DT_BIND_NOW     DynTag = 24 /* [sup] */
	DT_INIT_ARRAY   DynTag = 25 /* Address of the array of pointers to initialization functions */
	DT_FINI_ARRAY   DynTag = 26 /* Address of the array of pointers to termination functions */
	DT_INIT_ARRAYSZ DynTag = 27 /* Size in bytes of the array of initialization functions. */
	DT_FINI_ARRAYSZ DynTag = 28 /* Size in bytes of the array of termination functions. */
	DT_RUNPATH      DynTag = 29 /* String table offset of a null-terminated library search path string. */
	DT_FLAGS        DynTag = 30 /* Object specific flag values. */
	DT_ENCODING     DynTag = 32 /* Values greater than or equal to DT_ENCODING
	   and less than DT_LOOS follow the rules for
	   the interpretation of the d_un union
	   as follows: even == 'd_ptr', even == 'd_val'
	   or none */
	DT_PREINIT_ARRAY   DynTag = 32 /* Address of the array of pointers to pre-initialization functions. */
	DT_PREINIT_ARRAYSZ DynTag = 33 /* Size in bytes of the array of pre-initialization functions. */
	DT_SYMTAB_SHNDX    DynTag = 34 /* Address of SHT_SYMTAB_SHNDX section. */

	DT_LOOS DynTag = 0x6000000d /* First OS-specific */
	DT_HIOS DynTag = 0x6ffff000 /* Last OS-specific */

	DT_VALRNGLO       DynTag = 0x6ffffd00
	DT_GNU_PRELINKED  DynTag = 0x6ffffdf5
	DT_GNU_CONFLICTSZ DynTag = 0x6ffffdf6
	DT_GNU_LIBLISTSZ  DynTag = 0x6ffffdf7
	DT_CHECKSUM       DynTag = 0x6ffffdf8
	DT_PLTPADSZ       DynTag = 0x6ffffdf9
	DT_MOVEENT        DynTag = 0x6ffffdfa
	DT_MOVESZ         DynTag = 0x6ffffdfb
	DT_FEATURE        DynTag = 0x6ffffdfc
	DT_POSFLAG_1      DynTag = 0x6ffffdfd
	DT_SYMINSZ        DynTag = 0x6ffffdfe
	DT_SYMINENT       DynTag = 0x6ffffdff
	DT_VALRNGHI       DynTag = 0x6ffffdff

	DT_ADDRRNGLO    DynTag = 0x6ffffe00
	DT_GNU_HASH     DynTag = 0x6ffffef5
	DT_TLSDESC_PLT  DynTag = 0x6ffffef6
	DT_TLSDESC_GOT  DynTag = 0x6ffffef7
	DT_GNU_CONFLICT DynTag = 0x6ffffef8
	DT_GNU_LIBLIST  DynTag = 0x6ffffef9
	DT_CONFIG       DynTag = 0x6ffffefa
	DT_DEPAUDIT     DynTag = 0x6ffffefb
	DT_AUDIT        DynTag = 0x6ffffefc
	DT_PLTPAD       DynTag = 0x6ffffefd
	DT_MOVETAB      DynTag = 0x6ffffefe
	DT_SYMINFO      DynTag = 0x6ffffeff
	DT_ADDRRNGHI    DynTag = 0x6ffffeff

	DT_VERSYM     DynTag = 0x6ffffff0
	DT_RELACOUNT  DynTag = 0x6ffffff9
	DT_RELCOUNT   DynTag = 0x6ffffffa
	DT_FLAGS_1    DynTag = 0x6ffffffb
	DT_VERDEF     DynTag = 0x6ffffffc
	DT_VERDEFNUM  DynTag = 0x6ffffffd
	DT_VERNEED    DynTag = 0x6ffffffe
	DT_VERNEEDNUM DynTag = 0x6fffffff

	DT_LOPROC DynTag = 0x70000000 /* First processor-specific type. */

	DT_MIPS_RLD_VERSION           DynTag = 0x70000001
	DT_MIPS_TIME_STAMP            DynTag = 0x70000002
	DT_MIPS_ICHECKSUM             DynTag = 0x70000003
	DT_MIPS_IVERSION              DynTag = 0x70000004
	DT_MIPS_FLAGS                 DynTag = 0x70000005
	DT_MIPS_BASE_ADDRESS          DynTag = 0x70000006
	DT_MIPS_MSYM                  DynTag = 0x70000007
	DT_MIPS_CONFLICT              DynTag = 0x70000008
	DT_MIPS_LIBLIST               DynTag = 0x70000009
	DT_MIPS_LOCAL_GOTNO           DynTag = 0x7000000a
	DT_MIPS_CONFLICTNO            DynTag = 0x7000000b
	DT_MIPS_LIBLISTNO             DynTag = 0x70000010
	DT_MIPS_SYMTABNO              DynTag = 0x70000011
	DT_MIPS_UNREFEXTNO            DynTag = 0x70000012
	DT_MIPS_GOTSYM                DynTag = 0x70000013
	DT_MIPS_HIPAGENO              DynTag = 0x70000014
	DT_MIPS_RLD_MAP               DynTag = 0x70000016
	DT_MIPS_DELTA_CLASS           DynTag = 0x70000017
	DT_MIPS_DELTA_CLASS_NO        DynTag = 0x70000018
	DT_MIPS_DELTA_INSTANCE        DynTag = 0x70000019
	DT_MIPS_DELTA_INSTANCE_NO     DynTag = 0x7000001a
	DT_MIPS_DELTA_RELOC           DynTag = 0x7000001b
	DT_MIPS_DELTA_RELOC_NO        DynTag = 0x7000001c
	DT_MIPS_DELTA_SYM             DynTag = 0x7000001d
	DT_MIPS_DELTA_SYM_NO          DynTag = 0x7000001e
	DT_MIPS_DELTA_CLASSSYM        DynTag = 0x70000020
	DT_MIPS_DELTA_CLASSSYM_NO     DynTag = 0x70000021
	DT_MIPS_CXX_FLAGS             DynTag = 0x70000022
	DT_MIPS_PIXIE_INIT            DynTag = 0x70000023
	DT_MIPS_SYMBOL_LIB            DynTag = 0x70000024
	DT_MIPS_LOCALPAGE_GOTIDX      DynTag = 0x70000025
	DT_MIPS_LOCAL_GOTIDX          DynTag = 0x70000026
	DT_MIPS_HIDDEN_GOTIDX         DynTag = 0x70000027
	DT_MIPS_PROTECTED_GOTIDX      DynTag = 0x70000028
	DT_MIPS_OPTIONS               DynTag = 0x70000029
	DT_MIPS_INTERFACE             DynTag = 0x7000002a
	DT_MIPS_DYNSTR_ALIGN          DynTag = 0x7000002b
	DT_MIPS_INTERFACE_SIZE        DynTag = 0x7000002c
	DT_MIPS_RLD_TEXT_RESOLVE_ADDR DynTag = 0x7000002d
	DT_MIPS_PERF_SUFFIX           DynTag = 0x7000002e
	DT_MIPS_COMPACT_SIZE          DynTag = 0x7000002f
	DT_MIPS_GP_VALUE              DynTag = 0x70000030
	DT_MIPS_AUX_DYNAMIC           DynTag = 0x70000031
	DT_MIPS_PLTGOT                DynTag = 0x70000032
	DT_MIPS_RWPLT                 DynTag = 0x70000034
	DT_MIPS_RLD_MAP_REL           DynTag = 0x70000035

	DT_PPC_GOT DynTag = 0x70000000
	DT_PPC_OPT DynTag = 0x70000001

	DT_PPC64_GLINK DynTag = 0x70000000
	DT_PPC64_OPD   DynTag = 0x70000001
	DT_PPC64_OPDSZ DynTag = 0x70000002
	DT_PPC64_OPT   DynTag = 0x70000003

	DT_SPARC_REGISTER DynTag = 0x70000001

	DT_AUXILIARY DynTag = 0x7ffffffd
	DT_USED      DynTag = 0x7ffffffe
	DT_FILTER    DynTag = 0x7fffffff

	DT_HIPROC DynTag = 0x7fffffff /* Last processor-specific type. */
)

// DT_FLAGS values.
type DynFlag int

const (
	DF_ORIGIN DynFlag = 0x0001 /* Indicates that the object being loaded may
	   make reference to the
	   $ORIGIN substitution string */
	DF_SYMBOLIC DynFlag = 0x0002 /* Indicates "symbolic" linking. */
	DF_TEXTREL  DynFlag = 0x0004 /* Indicates there may be relocations in non-writable segments. */
	DF_BIND_NOW DynFlag = 0x0008 /* Indicates that the dynamic linker should
	   process all relocations for the object
	   containing this entry before transferring
	   control to the program. */
	DF_STATIC_TLS DynFlag = 0x0010 /* Indicates that the shared object or
	   executable contains code using a static
	   thread-local storage scheme. */
)

// DT_FLAGS_1 values.
type DynFlag1 uint32

const (
	// Indicates that all relocations for this object must be processed before
	// returning control to the program.
	DF_1_NOW DynFlag1 = 0x00000001
	// Unused.
	DF_1_GLOBAL DynFlag1 = 0x00000002
	// Indicates that the object is a member of a group.
	DF_1_GROUP DynFlag1 = 0x00000004
	// Indicates that the object cannot be deleted from a process.
	DF_1_NODELETE DynFlag1 = 0x00000008
	// Meaningful only for filters. Indicates that all associated filtees be
	// processed immediately.
	DF_1_LOADFLTR DynFlag1 = 0x00000010
	// Indicates that this object's initialization section be run before any other
	// objects loaded.
	DF_1_INITFIRST DynFlag1 = 0x00000020
	// Indicates that the object cannot be added to a running process with dlopen.
	DF_1_NOOPEN DynFlag1 = 0x00000040
	// Indicates the object requires $ORIGIN processing.
	DF_1_ORIGIN DynFlag1 = 0x00000080
	// Indicates that the object should use direct binding information.
	DF_1_DIRECT DynFlag1 = 0x00000100
	// Unused.
	DF_1_TRANS DynFlag1 = 0x00000200
	// Indicates that the objects symbol table is to interpose before all symbols
	// except the primary load object, which is typically the executable.
	DF_1_INTERPOSE DynFlag1 = 0x00000400
	// Indicates that the search for dependencies of this object ignores any
	// default library search paths.
	DF_1_NODEFLIB DynFlag1 = 0x00000800
	// Indicates that this object is not dumped by dldump. Candidates are objects
	// with no relocations that might get included when generating alternative
	// objects using.
	DF_1_NODUMP DynFlag1 = 0x00001000
	// Identifies this object as a configuration alternative object generated by
	// crle. Triggers the runtime linker to search for a configuration file $ORIGIN/ld.config.app-name.
	DF_1_CONFALT DynFlag1 = 0x00002000
	// Meaningful only for filtees. Terminates a filters search for any
	// further filtees.
	DF_1_ENDFILTEE DynFlag1 = 0x00004000
	// Indicates that this object has displacement relocations applied.
	DF_1_DISPRELDNE DynFlag1 = 0x00008000
	// Indicates that this object has displacement relocations pending.
	DF_1_DISPRELPND DynFlag1 = 0x00010000
	// Indicates that this object contains symbols that cannot be directly
	// bound to.
	DF_1_NODIRECT DynFlag1 = 0x00020000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_IGNMULDEF DynFlag1 = 0x00040000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_NOKSYMS DynFlag1 = 0x00080000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_NOHDR DynFlag1 = 0x00100000
	// Indicates that this object has been edited or has been modified since the
	// objects original construction by the link-editor.
	DF_1_EDITED DynFlag1 = 0x00200000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_NORELOC DynFlag1 = 0x00400000
	// Indicates that the object contains individual symbols that should interpose
	// before all symbols except the primary load object, which is typically the
	// executable.
	DF_1_SYMINTPOSE DynFlag1 = 0x00800000
	// Indicates that the executable requires global auditing.
	DF_1_GLOBAUDIT DynFlag1 = 0x01000000
	// Indicates that the object defines, or makes reference to singleton symbols.
	DF_1_SINGLETON DynFlag1 = 0x02000000
	// Indicates that the object is a stub.
	DF_1_STUB DynFlag1 = 0x04000000
	// Indicates that the object is a position-independent executable.
	DF_1_PIE DynFlag1 = 0x08000000
	// Indicates that the object is a kernel module.
	DF_1_KMOD DynFlag1 = 0x10000000
	// Indicates that the object is a weak standard filter.
	DF_1_WEAKFILTER DynFlag1 = 0x20000000
	// Unused.
	DF_1_NOCOMMON DynFlag1 = 0x40000000
)

// NType values; used in core files.
type NType int

const (
	NT_PRSTATUS NType = 1 /* Process status. */
	NT_FPREGSET NType = 2 /* Floating point registers. */
	NT_PRPSINFO NType = 3 /* Process state info. */
)

/* Symbol Binding - ELFNN_ST_BIND - st_info */
type SymBind int

const (
	STB_LOCAL  SymBind = 0  /* Local symbol */
	STB_GLOBAL SymBind = 1  /* Global symbol */
	STB_WEAK   SymBind = 2  /* like global - lower precedence */
	STB_LOOS   SymBind = 10 /* Reserved range for operating system */
	STB_HIOS   SymBind = 12 /*   specific semantics. */
	STB_LOPROC SymBind = 13 /* reserved range for processor */
	STB_HIPROC SymBind = 15 /*   specific semantics. */
)

/* Symbol type - ELFNN_ST_TYPE - st_info */
type SymType int

const (
	STT_NOTYPE  SymType = 0  /* Unspecified type. */
	STT_OBJECT  SymType = 1  /* Data object. */
	STT_FUNC    SymType = 2  /* Function. */
	STT_SECTION SymType = 3  /* Section. */
	STT_FILE    SymType = 4  /* Source file. */
	STT_COMMON  SymType = 5  /* Uninitialized common block. */
	STT_TLS     SymType = 6  /* TLS object. */
	STT_LOOS    SymType = 10 /* Reserved range for operating system */
	STT_HIOS    SymType = 12 /*   specific semantics. */
	STT_LOPROC  SymType = 13 /* reserved range for processor */
	STT_HIPROC  SymType = 15 /*   specific semantics. */
)

/* Symbol visibility - ELFNN_ST_VISIBILITY - st_other */
type SymVis int

const (
	STV_DEFAULT   SymVis = 0x0 /* Default visibility (see binding). */
	STV_INTERNAL  SymVis = 0x1 /* Special meaning in relocatable objects. */
	STV_HIDDEN    SymVis = 0x2 /* Not visible. */
	STV_PROTECTED SymVis = 0x3 /* Visible but not preemptible. */
)

// Magic number for the elf trampoline, chosen wisely to be an immediate value.
const ARM_MAGIC_TRAMP_NUMBER = 0x5c000003
