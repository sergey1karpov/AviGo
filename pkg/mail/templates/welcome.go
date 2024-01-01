package templates

func Welcome() string {
	template := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Welcome</title>
		</head>
		<body>
			<h1>Welcome to our platform!</h1>
			<p>Dear User,</p>
			<p>Thank you for joining us! We are excited to have you on board.</p>
			<p>Please proceed to our {LINK} to access your account.</p>
			<p>Best regards,</p>
			<p>The Support Team</p>
			<p>{YEAR}</p>
		</body>
		</html>
	`

	return template
}
