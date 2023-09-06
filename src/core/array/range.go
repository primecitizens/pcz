// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package array

type Zero[T any] interface{ ~[0]T }

type (
	Range1To100[T any] interface {
		Range1To8[T] | Range9To16[T] | Range17To24[T] | Range25To32[T] |
			Range33To40[T] | Range41To48[T] | Range49To56[T] | Range57To64[T] |
			Range65To72[T] | Range73To80[T] | Range81To88[T] | Range89To96[T] |
			~[97]T | ~[98]T | ~[99]T | ~[100]T
	}

	Range100To200[T any] interface {
		~[101]T | ~[102]T | ~[103]T | ~[104]T |
			Range105To112[T] | Range113To120[T] | Range121To128[T] | Range129To136[T] |
			Range137To144[T] | Range145To152[T] | Range153To160[T] | Range161To168[T] |
			Range169To176[T] | Range177To184[T] | Range185To192[T] | Range193To200[T]
	}

	Range1To8[T any] interface {
		~[1]T | ~[2]T | ~[3]T | ~[4]T | ~[5]T | ~[6]T | ~[7]T | ~[8]T
	}

	Range9To16[T any] interface {
		~[9]T | ~[10]T | ~[11]T | ~[12]T | ~[13]T | ~[14]T | ~[15]T | ~[16]T
	}

	Range17To24[T any] interface {
		~[17]T | ~[18]T | ~[19]T | ~[20]T | ~[21]T | ~[22]T | ~[23]T | ~[24]T
	}

	Range25To32[T any] interface {
		~[25]T | ~[26]T | ~[27]T | ~[28]T | ~[29]T | ~[30]T | ~[31]T | ~[32]T
	}

	Range33To40[T any] interface {
		~[33]T | ~[34]T | ~[35]T | ~[36]T | ~[37]T | ~[38]T | ~[39]T | ~[40]T
	}

	Range41To48[T any] interface {
		~[41]T | ~[42]T | ~[43]T | ~[44]T | ~[45]T | ~[46]T | ~[47]T | ~[48]T
	}

	Range49To56[T any] interface {
		~[49]T | ~[50]T | ~[51]T | ~[52]T | ~[53]T | ~[54]T | ~[55]T | ~[56]T
	}

	Range57To64[T any] interface {
		~[57]T | ~[58]T | ~[59]T | ~[60]T | ~[61]T | ~[62]T | ~[63]T | ~[64]T
	}

	Range65To72[T any] interface {
		~[65]T | ~[66]T | ~[67]T | ~[68]T | ~[69]T | ~[70]T | ~[71]T | ~[72]T
	}

	Range73To80[T any] interface {
		~[73]T | ~[74]T | ~[75]T | ~[76]T | ~[77]T | ~[78]T | ~[79]T | ~[80]T
	}

	Range81To88[T any] interface {
		~[81]T | ~[82]T | ~[83]T | ~[84]T | ~[85]T | ~[86]T | ~[87]T | ~[88]T
	}

	Range89To96[T any] interface {
		~[89]T | ~[90]T | ~[91]T | ~[92]T | ~[93]T | ~[94]T | ~[95]T | ~[96]T
	}

	Range97To104[T any] interface {
		~[97]T | ~[98]T | ~[99]T | ~[100]T | ~[101]T | ~[102]T | ~[103]T | ~[104]T
	}

	Range105To112[T any] interface {
		~[105]T | ~[106]T | ~[107]T | ~[108]T | ~[109]T | ~[110]T | ~[111]T | ~[112]T
	}

	Range113To120[T any] interface {
		~[113]T | ~[114]T | ~[115]T | ~[116]T | ~[117]T | ~[118]T | ~[119]T | ~[120]T
	}

	Range121To128[T any] interface {
		~[121]T | ~[122]T | ~[123]T | ~[124]T | ~[125]T | ~[126]T | ~[127]T | ~[128]T
	}

	Range129To136[T any] interface {
		~[129]T | ~[130]T | ~[131]T | ~[132]T | ~[133]T | ~[134]T | ~[135]T | ~[136]T
	}

	Range137To144[T any] interface {
		~[137]T | ~[138]T | ~[139]T | ~[140]T | ~[141]T | ~[142]T | ~[143]T | ~[144]T
	}

	Range145To152[T any] interface {
		~[145]T | ~[146]T | ~[147]T | ~[148]T | ~[149]T | ~[150]T | ~[151]T | ~[152]T
	}

	Range153To160[T any] interface {
		~[153]T | ~[154]T | ~[155]T | ~[156]T | ~[157]T | ~[158]T | ~[159]T | ~[160]T
	}

	Range161To168[T any] interface {
		~[161]T | ~[162]T | ~[163]T | ~[164]T | ~[165]T | ~[166]T | ~[167]T | ~[168]T
	}

	Range169To176[T any] interface {
		~[169]T | ~[170]T | ~[171]T | ~[172]T | ~[173]T | ~[174]T | ~[175]T | ~[176]T
	}

	Range177To184[T any] interface {
		~[177]T | ~[178]T | ~[179]T | ~[180]T | ~[181]T | ~[182]T | ~[183]T | ~[184]T
	}

	Range185To192[T any] interface {
		~[185]T | ~[186]T | ~[187]T | ~[188]T | ~[189]T | ~[190]T | ~[191]T | ~[192]T
	}

	Range193To200[T any] interface {
		~[193]T | ~[194]T | ~[195]T | ~[196]T | ~[197]T | ~[198]T | ~[199]T | ~[200]T
	}
)
