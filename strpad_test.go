package strpad

import "testing"

func TestPad(t *testing.T) {
    ans := Pad("test")
    if ans != " test " {
        t.Errorf("Pad(\"test\") = %s; want \" test \"")
    }
}


func TestPadLeft(t *testing.T) {
    ans := PadLeft("test")
    if ans != " test" {
        t.Errorf("PadLeft(\"test\") = %s; want \" test\"")
    }
}

func TestPadRight(t *testing.T) {
    ans := PadRight("test")
    if ans != "test " {
        t.Errorf("PadLeft(\"test\") = %s; want \"test \"")
    }
}