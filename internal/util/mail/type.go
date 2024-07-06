package mail

type (
	EmailVerification struct {
		Token string
	}

	PasswordReset struct {
		Token string
	}
)
