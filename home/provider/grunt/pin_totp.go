package grunt

import (
	"crypto/subtle"
	"math"
	"time"

	"github.com/Gophercraft/core/grunt"
	"github.com/pquerna/otp/hotp"
	"github.com/pquerna/otp/totp"
)

func totp_validate_pin(server_pin_info *grunt.ServerPINInfo, client_salt []byte, hashed_passcode []byte, secret string, t time.Time, opts totp.ValidateOpts) (bool, error) {
	if opts.Period == 0 {
		opts.Period = 30
	}

	counters := []uint64{}
	counter := int64(math.Floor(float64(t.Unix()) / float64(opts.Period)))

	counters = append(counters, uint64(counter))
	for i := 1; i <= int(opts.Skew); i++ {
		counters = append(counters, uint64(counter+int64(i)))
		counters = append(counters, uint64(counter-int64(i)))
	}

	for _, counter := range counters {
		rv, err := hotp_validate_pin(server_pin_info, client_salt, hashed_passcode, counter, secret, hotp.ValidateOpts{
			Digits:    opts.Digits,
			Algorithm: opts.Algorithm,
		})

		if err != nil {
			return false, err
		}

		if rv == true {
			return true, nil
		}
	}

	return false, nil
}

func hotp_validate_pin(server_pin_info *grunt.ServerPINInfo, client_salt []byte, hashed_passcode []byte, counter uint64, secret string, opts hotp.ValidateOpts) (bool, error) {
	otpstr, err := hotp.GenerateCodeCustom(secret, counter, opts)
	if err != nil {
		return false, err
	}
	converted_digits := grunt.ScramblePINNumber(server_pin_info.GridSeed, []byte(otpstr))
	hashed_otpstr := grunt.GetPINProof(server_pin_info.Salt[:], converted_digits, client_salt)

	if subtle.ConstantTimeCompare(hashed_otpstr, hashed_passcode) == 1 {
		return true, nil
	}

	return false, nil
}
