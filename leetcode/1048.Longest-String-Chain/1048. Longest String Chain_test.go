package leetcode

import (
	"fmt"
	"testing"
)

type question1048 struct {
	para1048
	ans1048
}

// para 是参数
// one 代表第一个参数
type para1048 struct {
	words []string
}

// ans 是答案
// one 代表第一个答案
type ans1048 struct {
	one int
}

func Test_Problem1048(t *testing.T) {

	qs := []question1048{

		{
			para1048{[]string{"a", "b", "ab", "bac"}},
			ans1048{2},
		},

		{
			para1048{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}},
			ans1048{5},
		},

		{
			para1048{[]string{"a", "b", "ba", "bca", "bda", "bdca"}},
			ans1048{4},
		},

		{
			para1048{[]string{"qjcaeymang", "bqiq", "bcntqiqulagurhz", "lyctmomvis", "bdnhym", "crxrdlv", "wo", "kijftxssyqmui", "abtcrjs", "rceecupq", "crxrdclv", "tvwkxrev", "oc", "lrzzcl", "snpzuykyobci", "abbtczrjs", "rpqojpmv",
				"kbfbcjxgvnb", "uqvhuucupu", "fwoquoih", "ezsuqxunx", "biq", "crxrwdclv", "qoyfqhytzxfp", "aryqceercpaupqm", "tvwxrev", "gchusjxz", "uls", "whb", "natdmc", "jvidsf", "yhyz", "smvsitdbutamn", "gcfsghusjsxiz", "ijpyhk", "tzvqwkmzxruevs",
				"fwvjxaxrvmfm", "wscxklqmxhn", "velgcy", "lyctomvi", "smvsitbutam", "hfosz", "fuzubrpo", "dfdeidcepshvjn", "twqol", "rpqjpmv", "ijftxssyqmi", "dzuzsainzbsx", "qyzxfp", "tvwkmzxruev", "farfm", "bbwkizqhicip", "wqobtmamvpgluh", "rytspgy",
				"uqvheuucdupuw", "jcmang", "h", "kijfhtxssyqmui", "twqgolksq", "rtkgopofnykkrl", "smvstbutam", "xkbfbbcjxgvnbq", "feyq", "oyfqhytzxfp", "velgcby", "dmnioxbf", "kbx", "zx", "wscxkqlqmxbhn", "efvcjtgoiga", "jumttwxv", "zux", "z", "smvsitbutamn",
				"jftxssyqmi", "wnlhsaj", "bbwizqhcp", "yctomv", "oyqyzxfp", "wqhc", "jnnwp", "bcntqiquagurz", "qzx", "kbfbjxn", "dmnixbf", "ukqs", "fey", "ryqceecaupq", "smvlsitzdbutamn", "bdnhm", "lrhtwfosrzq", "nkptknldw", "crxrwdclvx", "abbtcwzrjs",
				"uqvheuucupu", "abjbtcwbzrjs", "nkmptknldw", "wnulhsbaj", "wnlhsbaj", "wqobtmamvgluh", "jvis", "pcd", "s", "kjuannelajc", "valas", "lrrzzcl", "kjuannelajct", "snyyoc", "jwp", "vbum", "ezuunx", "bcntqiquagur", "vals", "cov", "dfdidcepshvjn",
				"vvamlasl", "budnhym", "h", "fwxarvfm", "lrhwfosrz", "nkptnldw", "vjhse", "zzeb", "fubrpo", "fkla", "qjulm", "xpdcdxqia", "ucwxwdm", "jvidsfr", "exhc", "kbfbjx", "bcntqiquaur", "fwnoqjuoihe", "ezsruqxuinrpxc", "ec", "dzuzstuacinzbvsx",
				"cxkqmxhn", "egpveohyvq", "bkcv", "dzuzsaizbx", "jftxssymi", "ycov", "zbvbeai", "ch", "atcrjs", "qjcemang", "tvjhsed", "vamlas", "bundnhym", "li", "wnulfhsbaj", "o", "ijhtpyhkrif", "nyoc", "ov", "ryceecupq", "wjcrjnszipc", "lrhtwfosrz",
				"tbzngeqcz", "awfotfiqni", "azbw", "o", "gcfghusjsxiz", "uqvheuucdupu", "rypgy", "snpuykyobc", "ckhn", "kbfbcjxgnb", "xkeow", "jvids", "ubnsnusvgmqog", "endjbkjere", "fwarfm", "wvhb", "fwnoqtjuozihe", "jnwp", "awfotfmiyqni", "iv", "ryqceecupq",
				"y", "qjuelm", "qyzxp", "vsbum", "dnh", "fam", "snpuyyobc", "wqobtmamvglu", "gjpw", "jcemang", "ukqso", "evhlfz", "nad", "bucwxwdm", "xkabfbbcjxgvnbq", "fwnoqjuozihe", "smvsitzdbutamn", "vec", "fos", "abbtcwbzrjs", "uyifxeyq", "kbfbjxgnb",
				"nyyoc", "kcv", "fbundnhym", "tbzngpeqcz", "yekfvcjtgoiga", "rgjpw", "ezgvhsalfz", "yoc", "ezvhsalfz", "crxdlv", "chusjz", "fwxaxrvfm", "dzuzstuacinzbsx", "bwizqhc", "pdcdx", "dmnioxbmf", "zuunx", "oqyzxfp", "ezsruqxuinxc", "qjcaemang", "gcghusjsxiz", "nktnldw",
				"qoyfqhytxzxfp", "bwqhc", "btkcvj", "qxpdcdxqia", "kijofhtxssyqmui", "rypy", "helmi", "zkrlexhxbwt", "qobtmamgu", "vhlfz", "rqjpmv", "yhy", "zzembhy", "rjpmv", "jhse", "fosz", "twol", "qbtamu", "nawxxbslyhucqxb", "dzzsaizbx", "dmnijgoxsbmf",
				"ijhtpyhkr", "yp", "awfotfgmiyqni", "yctov", "hse", "azabw", "aryqceercaupqm", "fuzubrpoa", "ubnswnusvgmqog", "fafm", "i", "ezvhalfz", "aryxqceercpaupqm", "bwizqhcp", "pdcdxq", "wscxkqlqmxhn", "fuubrpo", "fwvxaxrvmfm", "abjbtcwbzrjas", "zx",
				"rxmiirbxemog", "dfdeidcepshvvjqn", "az", "velc", "zkrlexnhxbwt", "nawxbslyhucqxb", "qjugelm", "ijhtpdyhkrif", "dmixbf", "gcfsghtusjsxiz", "juannlajc", "uqvheuucdupmuw", "rpqojpmgxv", "rpqojpmxv", "xppph", "kjuannellajct", "lrhfosrz", "dmnijoxsbmf",
				"ckmxhn", "tvijhsed", "dzuzstuainzbsx", "exhvc", "tvwkxruev", "rxmiirbemog", "lhfosz", "fkyla", "tlwqgolksq", "velgcbpy", "bcqiqaur", "xkhfejow", "ezsuqunx", "dmnioxsbmf", "bqiqu", "ijhtpudyhkrif", "xpdcdxqi", "ckh", "nwxbslyhucqxb", "bbwkizqhcip", "pcdx",
				"dzuzsuainzbsx", "xkbfbcjxgvnbq", "smvsbutm", "ezsruqxuinrxc", "smvlsitzdbutamdn", "am", "tvwkzxruev", "scxkqmxhn", "snpzuykyobc", "ekfvcjtgoiga", "fuzsubrpoa", "trtkgopofnykkrl", "oyqhytzxfp", "kbjx", "ifeyq", "vhl", "xkfeow", "ezgvhsialfz", "velgc", "hb",
				"zbveai", "gcghusjxz", "twqgolkq", "btkcv", "ryqceercaupq", "bi", "vvamlas", "awfotfmiqni", "abbtcrjs", "jutkqesoh", "xkbfbcjxgvnb", "hli", "ryspgy", "endjjjbkjere", "mvsbum", "ckqmxhn", "ezsruqxunxc", "zzeby", "xhc", "ezvhlfz", "ezsruqxunx", "tzvwkmzxruev",
				"hlmi", "kbbjx", "uqvhuuupu", "scxklqmxhn", "wqobtmamglu", "xpdcdxq", "qjugelym", "ifxeyq", "bcnqiquaur", "qobtmamglu", "xkabfbbcjxbgvnbq", "fuuzsubrpoa", "tvibjhsed", "oyqhyzxfp", "ijhpyhk", "c", "gcghusjxiz", "exhvoc", "awfotfini", "vhlz", "rtgopofykkrl",
				"yh", "ypy", "azb", "bwiqhc", "fla", "dmnijgioxsbmf", "chusjxz", "jvjidsfr", "natddmc", "uifxeyq", "x", "tzvqwkmzxruev", "bucwxwdwm", "ckmhn", "zzemby", "rpmv", "bcntqiqulagurz", "fwoqjuoihe", "dzuzsainzbx", "zkrlehxbwt", "kv", "ucwxwm", "ubnswnusvgmdqog",
				"wol", "endjjbkjere", "natyddmc", "vl", "ukqsoh", "ezuqunx", "exhvovc", "bqiqau", "bqiqaur", "zunx", "pc", "snuyyoc", "a", "lrhfosz", "kbfbjxgn", "rtgopofnykkrl", "hehszegkvse", "smvsbum", "ijhpyhkr", "ijftxssyqmui", "lyctomvis", "juanlajc", "jukqesoh",
				"xptpph", "fwarvfm", "qbtmamu", "twqgolq", "aryqceercaupq", "qbtmamgu", "rtgopofykkr", "snpuyyoc", "qyzx", "fwvxaxrvfm", "juannelajc", "fwoquoihe", "nadmc", "jumttwxvx", "ijhtpyhkrf", "twqolq", "rpv", "hehszegkuvse", "ls", "tvjhse", "rxmiirbemg",
				"dfdeidcepshvvjn", "dnhm", "egpeohyvq", "rgnjpw", "bbwkizqhcp", "nadc", "bcqiquaur", "xkhfeow", "smvstbutm", "ukqesoh", "yctomvi"}},
			ans1048{15},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1048------------------------\n")

	for _, q := range qs {
		_, p := q.ans1048, q.para1048
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestStrChain(p.words))
	}
	fmt.Printf("\n\n\n")
}
