package config

type Package struct {
	Name    string `yaml:"-"`
	Version string `yaml:"version"`
	Enabled bool   `yaml:"enabled"`
	Locked  bool   `yaml:"locked"`
}

func (pkg Package) SetVersion(version string) Package {
	pkg.Version = version
	return pkg
}
