package email

func (c *Client) SendWelcomeEmail(to, firstName string) error {
	data := map[string]string{
		"UserFirstName": firstName,
	}

	return c.sendEmail(
		to,
		"Welcome to Boilerplate",
		TemplateWelcome,
		data,
	)
}
