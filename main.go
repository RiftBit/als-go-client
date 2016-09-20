package alsgoclient


func New(config Config) (*Client, error) {

	_, err := config.validate()

	client := &Client{
		Config: Config{
			Uri: config.Uri,
			Login: config.Login,
			Password: config.Password,
			IsAsync: config.IsAsync,
			Timeout: config.Timeout,
		},
	}
	return client, err
}

