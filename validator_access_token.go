package main

func ensureAccessTokenAvailableAndValid() (string, error) {
	accessToken, err := ensureAccessTokenAvailable()
	if err != nil {
		return "", err
	}
	err = ensureAccessTokenValid(accessToken)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func ensureAccessTokenAvailable() (string, error) {
	return getAccessToken()
}

func ensureAccessTokenValid(accessToken string) error {
	_, err := getRepositories(accessToken)
	if err != nil {
		return err
	}
	return nil
}
