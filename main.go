package alsgoclient

func New(config Config) (*Client, error) {
	_, err := config.validate()
	if err != nil {
		return nil, err
	}
	client := &Client{config}
	return client, nil
}
