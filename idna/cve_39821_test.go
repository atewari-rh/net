package idna

import "testing"

func TestRejectASCIIOnlyPunycodeLabel(t *testing.T) {
    cases := []struct {
        name string
        fn   func(string) (string, error)
        in   string
    }{
        {"ToUnicode", ToUnicode, "xn--example-.com"},
        {"ToASCII", ToASCII, "xn--example-.com"},
        {"ProfileToUnicode", Lookup.ToUnicode, "xn--example-.com"},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            _, err := tc.fn(tc.in)
            if err == nil {
                t.Fatal("expected error for ASCII-only Punycode label")
            }
        })
    }
}