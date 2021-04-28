package service

import (
	"testing"
)

func TestRetrieveDeck(t *testing.T) {
	type args struct {
		id   string
		lang string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"CotA", args{"67913327-8b9d-47f7-ad6b-cabfa166b30c", "pt"}, "Raimundo “Quebra-Galho”, o Voltaico", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetrieveDeck(tt.args.id, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrieveDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || got.Data.Name != tt.want {
				t.Errorf("RetrieveDeck() got = %v, want %v", got, tt.want)
			}
		})
	}
}
