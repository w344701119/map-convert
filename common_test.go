package map_convert

import "testing"

func TestHumpKey(t *testing.T) {
	var strArr = []string{
		"aame_cececdfe_kuuud879er", "AAAame_cececdfe_kuuud879er", "CCvvame_cSSDDcecdfe_kuuud879er", "Bame_cececdfe_kuuud879er", "vvvrrwame_cececdfe_zuuud879er", "yuame_cececdfe_ZZuuud879er", "name_ceceEEcdfe_kuuud879er",
		"aame_xkmce_khdere", "EEEEEE_RRRRR_PPPP", "name_key_value_access",
	}
	for _, value := range strArr {
		k := HumpKey(value)
		t.Log(k)
	}

}

func BenchmarkHumpKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HumpKey("name_xiaom_DFFssr")
	}
}
