package metainfo

import "testing"

func TestTorrent_InfoHash(t1 *testing.T) {
	type fields struct {
		Announce     string
		Comment      string
		CreationDate int64
		CreatedBy    string
		Encoding     string
		Info         Info
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Empty info field",
			fields: fields{
				Announce:     "",
				Comment:      "",
				CreationDate: 0,
				CreatedBy:    "",
				Encoding:     "",
				Info: Info{
					Name:        "",
					Pieces:      "",
					PieceLength: 0,
					Private:     false,
					Length:      0,
					Files:       nil,
				},
			},
			want: string([]byte{128, 107, 163, 83, 219, 45, 121, 187, 48, 7, 64, 98, 105, 65, 58, 190, 195, 141, 110, 178}),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Torrent{
				Announce:     tt.fields.Announce,
				Comment:      tt.fields.Comment,
				CreationDate: tt.fields.CreationDate,
				CreatedBy:    tt.fields.CreatedBy,
				Encoding:     tt.fields.Encoding,
				Info:         tt.fields.Info,
			}
			if got := t.InfoHash(); got != tt.want {
				t1.Errorf("InfoHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTorrent_IsSingleFile(t1 *testing.T) {
	type fields struct {
		Announce     string
		Comment      string
		CreationDate int64
		CreatedBy    string
		Encoding     string
		Info         Info
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{{
		name: "Test a multiple file torrent",
		fields: fields{
			Announce:     "",
			Comment:      "",
			CreationDate: 0,
			CreatedBy:    "",
			Encoding:     "",
			Info: Info{
				Name:        "",
				Pieces:      "",
				PieceLength: 0,
				Private:     false,
				Length:      0,
				Files: []Files{{
					Length: 0,
					Md5sum: "",
					Path:   nil,
				}},
			},
		},
		want: false,
	}, {
		name: "Test a single file torrent",
		fields: fields{
			Announce:     "",
			Comment:      "",
			CreationDate: 0,
			CreatedBy:    "",
			Encoding:     "",
			Info: Info{
				Name:        "",
				Pieces:      "",
				PieceLength: 0,
				Private:     false,
				Length:      0,
				Files:       nil,
			},
		},
		want: true,
	}}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Torrent{
				Announce:     tt.fields.Announce,
				Comment:      tt.fields.Comment,
				CreationDate: tt.fields.CreationDate,
				CreatedBy:    tt.fields.CreatedBy,
				Encoding:     tt.fields.Encoding,
				Info:         tt.fields.Info,
			}
			if got := t.IsSingleFile(); got != tt.want {
				t1.Errorf("IsSingleFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
