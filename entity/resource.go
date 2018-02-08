package entity

type Resource struct {
	Id          string       `toml:"id"`
	Path        string       `toml:"path"`
	Integration *Integration `toml:"integration"`
	UserDefined bool         `toml:"user_defined"`
}

func NewResource(id, path string) *Resource {
	return &Resource{
		Id:          id,
		Path:        formatPath(path),
		Integration: nil,
		UserDefined: false,
	}
}
