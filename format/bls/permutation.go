package bls

import "fmt"

type Permutation struct {
	Unk1 uint32
	Unk2 uint32
	Unk3 uint16
	Unk4 uint16
	Code string
}

func (r *reader) readPermutations1_3() error {
	r.shader.Permutations = make([]Permutation, r.shader.PermutationCount)

	for i := 0; i < len(r.shader.Permutations); i++ {
		permute := &r.shader.Permutations[i]

		if err := r.readU32(&permute.Unk1); err != nil {
			return err
		}
		if err := r.readU32(&permute.Unk2); err != nil {
			return err
		}
		if err := r.readU16(&permute.Unk3); err != nil {
			return err
		}
		if err := r.readU16(&permute.Unk4); err != nil {
			return err
		}
		var codeLen uint32
		if err := r.readU32(&codeLen); err != nil {
			return err
		}
		codeSizeIsReasonable := codeLen < (1024 * 1024 * 10)
		if !codeSizeIsReasonable {
			return fmt.Errorf("bls: unreasonably large code length in permutation %d (length %d)", i, codeLen)
		}
		code := make([]byte, codeLen)
		if _, err := r.file.Read(code[:]); err != nil {
			return err
		}
		permute.Code = string(code)

		var pad [4]byte
		if _, err := r.file.Read(pad[:((codeLen+3)&0xFFFFFFFC)-codeLen]); err != nil {
			return err
		}
	}

	return nil
}

func (r *reader) readPermutations() error {
	switch r.shader.Version {
	case Version1_3:
		return r.readPermutations1_3()
	default:
		return fmt.Errorf("bls: cannot read permutations for version 0x%016X", r.shader.Version)
	}
}
