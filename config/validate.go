package config

import "errors"

func (c Config) Validate() error {
	if err := c.HTTP.Validate(); err != nil {
		return err
	}

	if err := c.GRPC.Validate(); err != nil {
		return err
	}

	return nil
}
func (c HTTP) Validate() error {
	if c.URL == "" {
		return errors.New("http.url is required")
	}

	if c.Timeout <= 0 {
		return errors.New("http.timeout must be positive")
	}

	return nil
}
func (c GRPC) Validate() error {
	if c.Address == "" {
		return errors.New("grpc.address is required")
	}

	if c.Timeout <= 0 {
		return errors.New("grpc.timeout must be positive")
	}

	return nil
}
