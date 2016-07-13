package locky

import "time"

// https://github.com/baderj/domain_generation_algorithms/blob/master/locky/dgav3.py

// Config is the Locky ransomware config holder
type Config struct {
	Seed  uint32
	Shift uint32
	TLDs  []string
}

var configs = []Config{
	{
		// md5: 81e85dcaf482aba2f8ea047145490493
		// sha256: 9afb127e733b01952f00a8174291d39f740b6df2c90d0422b4d6f2e2e6bc7d1a
		// sample: https://virustotal.com/en/file/9afb127e733b01952f00a8174291d39f740b6df2c90d0422b4d6f2e2e6bc7d1a/analysis/
		Seed:  7077,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 5fb8f8f75342ff68ed8c79cc375f0cd8
		// sha256: bc7c45b5a05f3f0deea162578e45d5fb64c9aa72a81395083509c0f78b6ae1de
		// sample: https://malwr.com/analysis/NzFlYzRkOWZhZDliNDZmMThkMzkzMjU2ZmE5ODUzMjE/
		Seed:  5566,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 19079496f6abfafd9a99d02098556a79
		// sha256: 5dd6188efe13268bb9ac20ecdb257085e7d62163
		// sample: https://malwr.com/analysis/Yzg5ZWMyZDZmNGFhNGU0YWJjMzY3YjBmMjY4Y2JlMDM/
		Seed:  9106,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 0e223d578eaddec361498591ec8c1a19
		// sha256:
		// sample:
		Seed:  5579,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 6cb11f4066f74556dd14d27008d867b4
		// sha256: 353ea18baa6c9c796a7b48bcccbf4c9c3c6aa63f8b4dd55d796c65e22028b77b
		// sample: https://malwr.com/analysis/N2Q4YWUyM2I3Y2VlNGQwYzllMjczNzE2Njc1ZTFhZWI/
		Seed:  111,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: bfb5fec661318f07b0eca8520bb8c53f
		// sha256: 92a1d459194d0bf86ff26103a9c92d64059e1caa9d98e1ed9002058a0da8f53d
		// sample: https://malwr.com/analysis/Mzc0MTU0MTQ1YjRlNGVhYzgzMmQ0MGQ3YWY1NWUzZDg/
		Seed:  9099,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: a075610a69e196ab74f79508dbcf5eef
		// sha256: caa6e59e98c22a3f9159412a612ad170d2683640e1845afb6f533f279f5e6577
		// sample: https://malwr.com/analysis/OWJhZjQ5ZmE4YjY4NDFhYmFhNjIxZDcyYmFkYzlhYTM/
		Seed:  9047,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 6c3e68307d01e4340c83fac94f237237
		// sha256:
		// sample:
		Seed:  9133,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 17bf0d1776de896315cb2d63118f9667
		// sha256: 98d6ebd37c861beaf7420494aa8dfd43e4904145bac62c607965b3a8d92967c1
		// sample: https://malwr.com/analysis/NWJiNjJhZWU2YjU2NGRlYjg5Njk2ZGJlMzZkZDcxYWI/
		Seed:  9199,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5:
		// sha256:
		// sample:
		Seed:  9190,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5:
		// sha256:
		// sample: https://malwr.com/analysis/NzY4YmRjZDA1MTYwNGEzMzg2MWZkNmUyODIzMWRhMDM/
		Seed:  9088,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 6eb8865bf055ba30cc9e2843f16ee461
		// sha256:
		// sample:
		Seed:  9998,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: b5660f65369abc67cfa4a65e7d0d65d9
		// sha256: 478ab3b1f465dc1088b0d1e7cef8cab1f3b736856f6be279d4e7a8113ad065d5
		// sample: https://www.virustotal.com/en/file/478ab3b1f465dc1088b0d1e7cef8cab1f3b736856f6be279d4e7a8113ad065d5/analysis/
		Seed:  9992,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: aceec3d6334e925297efc8d4232473c2
		// sha256: 5c18ab258a3a89980aaa9d673a07851fcab4443733a00c4fbf14d21906b65c9e
		// sample: https://www.virustotal.com/en/file/5c18ab258a3a89980aaa9d673a07851fcab4443733a00c4fbf14d21906b65c9e/analysis/1463993646/
		Seed:  1511,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5:
		// sha256:
		// sample: https://malwr.com/analysis/ODU3OWM4ZDMxMmE2NDllZWE4MWQ3ZGQ2ZTBjZTc4MWI/
		Seed:  1513,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: 89f35a5d22088e3ca8664697e895b7a5
		// sha256:
		// sample:
		Seed:  1517,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: a9d09406e0982d652b6db291558df71a
		// sha256:
		// sample:
		Seed:  9056,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
	{
		// md5: a9d09406e0982d652b6db291558df71a
		// sha256:
		// sample:
		Seed:  7773,
		Shift: 7,
		TLDs: []string{
			"ru", "info", "biz", "click", "su", "work", "pl", "org", "pw", "xyz",
		},
	},
}

// -----------------------------------------------------------------------------

func ror32(v, s uint32) uint32 {
	v &= 0xFFFFFFFF
	return ((v >> s) | (v << (32 - s))) & 0xFFFFFFFF
}

func rol32(v, s uint32) uint32 {
	v &= 0xFFFFFFFF
	return ((v << s) | (v >> (32 - s))) & 0xFFFFFFFF
}

func dga(date time.Time, configNr, domainNr uint32) string {
	c := configs[configNr]
	seed_shifted := rol32(c.Seed, 17)
	dnr_shifted := rol32(domainNr, 21)

	k := uint32(0)
	year := uint32(date.Year())
	month := uint32(date.Month())
	day := uint32(date.Day())

	for i := 0; i < 7; i++ {
		t_0 := ror32(0xB11924E1*(year+k+0x1BF5), c.Shift) & 0xFFFFFFFF
		t_1 := ((t_0 + 0x27100001) ^ k) & 0xFFFFFFFF
		t_2 := (ror32(0xB11924E1*(t_1+c.Seed), c.Shift)) & 0xFFFFFFFF
		t_3 := ((t_2 + 0x27100001) ^ t_1) & 0xFFFFFFFF
		t_4 := (ror32(0xB11924E1*(uint32(day/2)+t_3), c.Shift)) & 0xFFFFFFFF
		t_5 := (0xD8EFFFFF - t_4 + t_3) & 0xFFFFFFFF
		t_6 := (ror32(0xB11924E1*(month+t_5-0x65CAD), c.Shift)) & 0xFFFFFFFF
		t_7 := (t_5 + t_6 + 0x27100001) & 0xFFFFFFFF
		t_8 := (ror32(0xB11924E1*(t_7+seed_shifted+dnr_shifted), c.Shift)) & 0xFFFFFFFF
		k = ((t_8 + 0x27100001) ^ t_7) & 0xFFFFFFFF
		year++
	}

	length := (k % 11) + 7
	domain := ""
	for i := uint32(0); i < length; i++ {
		k = (ror32(0xB11924E1*rol32(k, i), c.Shift) + 0x27100001) & 0xFFFFFFFF
		domain += string(k%25 + uint32('a'))
	}

	domain += "."

	k = ror32(k*0xB11924E1, c.Shift)
	tlds := c.TLDs
	tld_i := ((k + 0x27100001) & 0xFFFFFFFF) % uint32(len(tlds))
	domain += tlds[tld_i]

	return domain
}
