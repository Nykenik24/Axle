package config

type ConfigRoot struct {
	PackageManagers map[string]PackageManager `yaml:"package_managers"`
	Settings        Settings                  `yaml:"global"`
	Graph           GraphSettings             `yaml:"graph"`
}

type PackageManager struct {
	Name     string             `yaml:"-"`
	Commands map[string]string  `yaml:"commands"`
	Packages map[string]Package `yaml:"packages"`
}

type Package struct {
	Name    string `yaml:"-"`
	Version string `yaml:"version"`
	Enabled bool   `yaml:"enabled"`
	Locked  bool   `yaml:"locked"`
}

type Settings struct {
	Parallel bool `yaml:"parallel"`
	Cache    bool `yaml:"cache"`
	Lockfile bool `yaml:"lockfile"`
}

type GraphSettings struct {
	Format     string `yaml:"format"`
	OutputFile string `yaml:"output_file"`
}
