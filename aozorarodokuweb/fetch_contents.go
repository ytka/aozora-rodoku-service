package aozorarodokuweb

func FetchContents() ([]Content, error) {
	meta, err := GetAppMeta()
	if err != nil {
		return nil, err
	}

	data, err := GetContentsCSV(meta.Url)
	if err != nil {
		return nil, err
	}

	return ParseContents(data)
}
