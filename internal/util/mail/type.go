package mail

type (
	EmailVerification struct {
		OTP string
	}

	PasswordReset struct {
		OTP string
	}
)
