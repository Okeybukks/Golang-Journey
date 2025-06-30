/*
This file is name main_internal_test.go because we want to test a file called main, internal
because we are accessing functions not exposed
*/
package main

import "testing"

func TestGreet(t *testing.T){
	type testCase struct{
		lang language
		want string
	}

	tests := map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום עולם",
		},
		"Urdu":{
			lang: "ur",
			want: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang: "vi",
			want: "Xin chào Thế Giới",
		},
	}

	for name, tc := range tests{
		t.Run(name, func(t *testing.T) {
			got := greet(language(tc.lang))
			if got != tc.want{
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}

func TestGreet_English(t *testing.T) {
	want := "Hello World!"
	lang := language("en")

	got := greet(lang)
	if got != want {
		// mark this as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	want := "Bonjour le monde!"
	lang := language("fr")

	got := greet(lang)

	if got != want {
		// mark this as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
