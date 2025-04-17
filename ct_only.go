package tessera

import (
	"crypto/sha256"
	"fmt"

	"github.com/transparency-dev/merkle/rfc6962"
	"github.com/transparency-dev/trillian-tessera/api/layout"
	"golang.org/x/crypto/cryptobyte"
)

// WithCTLayout instructs the underlying storage to use a Static CT API compatible scheme for layout.
func (o *AppendOptions) WithCTLayout() *AppendOptions {
	o.entriesPath = ctEntriesPath
	o.bundleIDHasher = ctBundleIDHasher
	return o
}

// WithCTLayout instructs the underlying storage to use a Static CT API compatible scheme for layout.
func (o *MigrationOptions) WithCTLayout() *MigrationOptions {
	o.entriesPath = ctEntriesPath
	o.bundleIDHasher = ctBundleIDHasher
	o.bundleLeafHasher = ctMerkleLeafHasher
	return o
}

func ctEntriesPath(n uint64, p uint8) string {
	return fmt.Sprintf("tile/data/%s", layout.NWithSuffix(0, n, p))
}

// ctBundleIDHasher knows how to calculate antispam identity hashes for entries in a Static-CT formatted entry bundle.
func ctBundleIDHasher(bundle []byte) ([][]byte, error) {
	r := make([][]byte, 0, layout.EntryBundleWidth)
	b := cryptobyte.String(bundle)
	for i := 0; i < layout.EntryBundleWidth && !b.Empty(); i++ {
		// Timestamp
		if !b.Skip(8) {
			return nil, fmt.Errorf("failed to read timestamp of entry index %d of bundle", i)
		}

		var entryType uint16
		if !b.ReadUint16(&entryType) {
			return nil, fmt.Errorf("failed to read entry type of entry index %d of bundle", i)
		}

		switch entryType {
		case 0: // X509 entry
			cert := cryptobyte.String{}
			if !b.ReadUint24LengthPrefixed(&cert) {
				return nil, fmt.Errorf("failed to read certificate at entry index %d of bundle", i)
			}

			// For x509 entries we hash (just) the x509 certificate for identity.
			r = append(r, identityHash(cert))

			// Must continue below to consume all the remaining bytes in the entry.

		case 1: // Precert entry
			// IssuerKeyHash
			if !b.Skip(sha256.Size) {
				return nil, fmt.Errorf("failed to read issuer key hash at entry index %d of bundle", i)
			}
			tbs := cryptobyte.String{}
			if !b.ReadUint24LengthPrefixed(&tbs) {
				return nil, fmt.Errorf("failed to read precert tbs at entry index %d of bundle", i)
			}

		default:
			return nil, fmt.Errorf("unknown entry type at entry index %d of bundle", i)
		}

		ignore := cryptobyte.String{}
		if !b.ReadUint16LengthPrefixed(&ignore) {
			return nil, fmt.Errorf("failed to read SCT extensions at entry index %d of bundle", i)
		}

		if entryType == 1 {
			precert := cryptobyte.String{}
			if !b.ReadUint24LengthPrefixed(&precert) {
				return nil, fmt.Errorf("failed to read precert at entry index %d of bundle", i)
			}
			// For Precert entries we hash (just) the full precertificate for identity.
			r = append(r, identityHash(precert))

		}
		if !b.ReadUint16LengthPrefixed(&ignore) {
			return nil, fmt.Errorf("failed to read chain fingerprints at entry index %d of bundle", i)
		}
	}
	if !b.Empty() {
		return nil, fmt.Errorf("unexpected %d bytes of trailing data in entry bundle", len(b))
	}
	return r, nil
}

// copyBytes copies N bytes between from and to.
func copyBytes(from *cryptobyte.String, to *cryptobyte.Builder, N int) bool {
	b := make([]byte, N)
	if !from.ReadBytes(&b, N) {
		return false
	}
	to.AddBytes(b)
	return true
}

// copyUint16LengthPrefixed copies a uint16 length and value between from and to.
func copyUint16LengthPrefixed(from *cryptobyte.String, to *cryptobyte.Builder) bool {
	b := cryptobyte.String{}
	if !from.ReadUint16LengthPrefixed(&b) {
		return false
	}
	to.AddUint16LengthPrefixed(func(c *cryptobyte.Builder) {
		c.AddBytes(b)
	})
	return true
}

// copyUint24LengthPrefixed copies a uint24 length and value between from and to.
func copyUint24LengthPrefixed(from *cryptobyte.String, to *cryptobyte.Builder) bool {
	b := cryptobyte.String{}
	if !from.ReadUint24LengthPrefixed(&b) {
		return false
	}
	to.AddUint24LengthPrefixed(func(c *cryptobyte.Builder) {
		c.AddBytes(b)
	})
	return true
}

// ctMerkleLeafHasher knows how to calculate RFC6962 Merkle leaf hashes for entries in a Static-CT formatted entry bundle.
func ctMerkleLeafHasher(bundle []byte) ([][]byte, error) {
	r := make([][]byte, 0, layout.EntryBundleWidth)
	b := cryptobyte.String(bundle)
	for i := 0; i < layout.EntryBundleWidth && !b.Empty(); i++ {
		preimage := &cryptobyte.Builder{}
		preimage.AddUint8(0 /* version = v1 */)
		preimage.AddUint8(0 /* leaf_type = timestamped_entry */)

		// Timestamp
		if !copyBytes(&b, preimage, 8) {
			return nil, fmt.Errorf("failed to copy timestamp of entry index %d of bundle", i)
		}

		var entryType uint16
		if !b.ReadUint16(&entryType) {
			return nil, fmt.Errorf("failed to read entry type of entry index %d of bundle", i)
		}
		preimage.AddUint16(entryType)

		switch entryType {
		case 0: // X509 entry
			if !copyUint24LengthPrefixed(&b, preimage) {
				return nil, fmt.Errorf("failed to copy certificate at entry index %d of bundle", i)
			}

		case 1: // Precert entry
			// IssuerKeyHash
			if !copyBytes(&b, preimage, sha256.Size) {
				return nil, fmt.Errorf("failed to copy issuer key hash at entry index %d of bundle", i)
			}

			if !copyUint24LengthPrefixed(&b, preimage) {
				return nil, fmt.Errorf("failed to copy precert tbs at entry index %d of bundle", i)
			}

		default:
			return nil, fmt.Errorf("unknown entry type 0x%x at entry index %d of bundle", entryType, i)
		}

		if !copyUint16LengthPrefixed(&b, preimage) {
			return nil, fmt.Errorf("failed to copy SCT extensions at entry index %d of bundle", i)
		}

		ignore := cryptobyte.String{}
		if entryType == 1 {
			if !b.ReadUint24LengthPrefixed(&ignore) {
				return nil, fmt.Errorf("failed to read precert at entry index %d of bundle", i)
			}
		}
		if !b.ReadUint16LengthPrefixed(&ignore) {
			return nil, fmt.Errorf("failed to read chain fingerprints at entry index %d of bundle", i)
		}

		h := rfc6962.DefaultHasher.HashLeaf(preimage.BytesOrPanic())
		r = append(r, h)
	}
	if !b.Empty() {
		return nil, fmt.Errorf("unexpected %d bytes of trailing data in entry bundle", len(b))
	}
	return r, nil
}
