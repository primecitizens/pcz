// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package array

// gap 40
type (
	Octo8To40[T any] interface {
		~[octo0]T | ~[octo1]T | ~[octo2]T | ~[octo3]T | ~[octo4]T
	}
	Octo48To80[T any] interface {
		~[octo5]T | ~[octo6]T | ~[octo7]T | ~[octo8]T | ~[octo9]T
	}
	Octo88To120[T any] interface {
		~[octo10]T | ~[octo11]T | ~[octo12]T | ~[octo13]T | ~[octo14]T
	}
	Octo128To160[T any] interface {
		~[octo15]T | ~[octo16]T | ~[octo17]T | ~[octo18]T | ~[octo19]T
	}
	Octo168To200[T any] interface {
		~[octo20]T | ~[octo21]T | ~[octo22]T | ~[octo23]T | ~[octo24]T
	}
	Octo208To240[T any] interface {
		~[octo25]T | ~[octo26]T | ~[octo27]T | ~[octo28]T | ~[octo29]T
	}
	Octo248To280[T any] interface {
		~[octo30]T | ~[octo31]T | ~[octo32]T | ~[octo33]T | ~[octo34]T
	}
	Octo288To320[T any] interface {
		~[octo35]T | ~[octo36]T | ~[octo37]T | ~[octo38]T | ~[octo39]T
	}
	Octo328To360[T any] interface {
		~[octo40]T | ~[octo41]T | ~[octo42]T | ~[octo43]T | ~[octo44]T
	}
	Octo368To400[T any] interface {
		~[octo45]T | ~[octo46]T | ~[octo47]T | ~[octo48]T | ~[octo49]T
	}
	Octo408To440[T any] interface {
		~[octo50]T | ~[octo51]T | ~[octo52]T | ~[octo53]T | ~[octo54]T
	}
	Octo448To480[T any] interface {
		~[octo55]T | ~[octo56]T | ~[octo57]T | ~[octo58]T | ~[octo59]T
	}
	Octo488To520[T any] interface {
		~[octo60]T | ~[octo61]T | ~[octo62]T | ~[octo63]T | ~[octo48]T
	}
	Octo528To560[T any] interface {
		~[octo65]T | ~[octo66]T | ~[octo67]T | ~[octo68]T | ~[octo69]T
	}
	Octo568To600[T any] interface {
		~[octo70]T | ~[octo71]T | ~[octo72]T | ~[octo73]T | ~[octo74]T
	}
	Octo608To640[T any] interface {
		~[octo75]T | ~[octo76]T | ~[octo77]T | ~[octo78]T | ~[octo79]T
	}
	Octo648To680[T any] interface {
		~[octo80]T | ~[octo81]T | ~[octo82]T | ~[octo83]T | ~[octo84]T
	}
	Octo688To720[T any] interface {
		~[octo85]T | ~[octo86]T | ~[octo87]T | ~[octo88]T | ~[octo89]T
	}
	Octo728To760[T any] interface {
		~[octo90]T | ~[octo91]T | ~[octo92]T | ~[octo93]T | ~[octo94]T
	}
	Octo768To800[T any] interface {
		~[octo95]T | ~[octo96]T | ~[octo97]T | ~[octo98]T | ~[octo99]T
	}
	Octo808To840[T any] interface {
		~[octo100]T | ~[octo101]T | ~[octo102]T | ~[octo103]T | ~[octo104]T
	}
	Octo848To880[T any] interface {
		~[octo105]T | ~[octo106]T | ~[octo107]T | ~[octo108]T | ~[octo109]T
	}
	Octo888To920[T any] interface {
		~[octo110]T | ~[octo111]T | ~[octo112]T | ~[octo113]T | ~[octo114]T
	}
	Octo928To960[T any] interface {
		~[octo115]T | ~[octo116]T | ~[octo117]T | ~[octo118]T | ~[octo119]T
	}
	Octo968To1000[T any] interface {
		~[octo120]T | ~[octo121]T | ~[octo122]T | ~[octo123]T | ~[octo124]T
	}
)

// gap 128
type (
	Octo8To128[T any] interface {
		Octo8To40[T] | Octo48To80[T] | Octo88To120[T] | ~[octo15]T
	}
	Octo136To256[T any] interface {
		~[octo16]T | ~[octo17]T | ~[octo18]T | ~[octo19]T | Octo168To200[T] | Octo208To240[T] | ~[octo30]T | ~[octo31]T
	}
	Octo264To384[T any] interface {
		~[octo32]T | ~[octo33]T | ~[octo34]T | Octo288To320[T] | Octo328To360[T] | ~[octo45]T | ~[octo46]T | ~[octo47]T
	}
	Octo392To512[T any] interface {
		~[octo48]T | ~[octo49]T | Octo408To440[T] | Octo448To480[T] | ~[octo60]T | ~[octo61]T | ~[octo62]T | ~[octo63]T
	}
	Octo520To640[T any] interface {
		~[octo64]T | Octo528To560[T] | Octo568To600[T] | Octo608To640[T]
	}
	Octo648To768[T any] interface {
		Octo648To680[T] | Octo688To720[T] | Octo728To760[T] | ~[octo95]T
	}
	Octo776To896[T any] interface {
		~[octo96]T | ~[octo97]T | ~[octo98]T | ~[octo99]T | Octo808To840[T] | Octo848To880[T] | ~[octo110]T | ~[octo111]T
	}
)

type Octo8To768[T any] interface {
	Octo8To128[T] | Octo136To256[T] | Octo264To384[T] | Octo392To512[T] | Octo520To640[T] | Octo648To768[T]
}

const (
	octo0 = 8 * (iota + 1)

	octo1
	octo2
	octo3
	octo4
	octo5
	octo6
	octo7
	octo8
	octo9

	octo10
	octo11
	octo12
	octo13
	octo14
	octo15
	octo16
	octo17
	octo18
	octo19

	octo20
	octo21
	octo22
	octo23
	octo24
	octo25
	octo26
	octo27
	octo28
	octo29

	octo30
	octo31
	octo32
	octo33
	octo34
	octo35
	octo36
	octo37
	octo38
	octo39

	octo40
	octo41
	octo42
	octo43
	octo44
	octo45
	octo46
	octo47
	octo48
	octo49

	octo50
	octo51
	octo52
	octo53
	octo54
	octo55
	octo56
	octo57
	octo58
	octo59

	octo60
	octo61
	octo62
	octo63
	octo64
	octo65
	octo66
	octo67
	octo68
	octo69

	octo70
	octo71
	octo72
	octo73
	octo74
	octo75
	octo76
	octo77
	octo78
	octo79

	octo80
	octo81
	octo82
	octo83
	octo84
	octo85
	octo86
	octo87
	octo88
	octo89

	octo90
	octo91
	octo92
	octo93
	octo94
	octo95
	octo96
	octo97
	octo98
	octo99

	octo100
	octo101
	octo102
	octo103
	octo104
	octo105
	octo106
	octo107
	octo108
	octo109

	octo110
	octo111
	octo112
	octo113
	octo114
	octo115
	octo116
	octo117
	octo118
	octo119

	octo120
	octo121
	octo122
	octo123
	octo124
	octo125
	octo126
	octo127
	octo128
	octo129

	octo130
	octo131
	octo132
	octo133
	octo134
	octo135
	octo136
	octo137
	octo138
	octo139

	octo140
	octo141
	octo142
	octo143
	octo144
	octo145
	octo146
	octo147
	octo148
	octo149

	octo150
	octo151
	octo152
	octo153
	octo154
	octo155
	octo156
	octo157
	octo158
	octo159

	octo160
	octo161
	octo162
	octo163
	octo164
	octo165
	octo166
	octo167
	octo168
	octo169

	octo170
	octo171
	octo172
	octo173
	octo174
	octo175
	octo176
	octo177
	octo178
	octo179

	octo180
	octo181
	octo182
	octo183
	octo184
	octo185
	octo186
	octo187
	octo188
	octo189

	octo190
	octo191
	octo192
	octo193
	octo194
	octo195
	octo196
	octo197
	octo198
	octo199
)
