package metainfo

type InfoHasher interface {
	InfoHash() string //SHA-1 Hash of the bencoded Info field of a metainfo file.
}
