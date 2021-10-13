package util

const FileName string = ".youtrack-config.json"

type Storage struct {
	Token string `json:"token"`
	Url string `json:"url"`
}

func (s *Storage) Init() (*Storage, error) {
	file, err := HomeFile(FileName)

	if err != nil {
		return nil, err
	}

	if FileExists(file) {
		return s.Load()
	}

	err = WriteJson(file, s.asMap())

	if err != nil { return nil, err }

	return s.Load()
}

func (s *Storage) Load() (*Storage, error) {
	file, err := HomeFile(FileName)

	if err != nil {
		return nil, err
	}

	var data map[string]string
	err = ReadJson(file, &data)

	if err != nil {
		return nil, err
	}

	storage := Storage {
		Token: data["token"],
		Url:   data["url"],
	}

	return &storage, nil
}

func (s Storage) Save() error {
	file, err := HomeFile(FileName)

	if err != nil {
		return err
	}

	return WriteJson(file, s.asMap())
}

func (s Storage) asMap() map[string]string {
	m := make(map[string]string)

	m["token"] = s.Token
	m["url"] = s.Url

	return m
}
