package bosh

import yaml "gopkg.in/yaml.v2"

type manifest struct {
	DirectorUUID interface{} `yaml:"director_uuid"`
	Name         interface{} `yaml:"name"`
	Compilation  interface{} `yaml:"compilation,omitempty"`
	Stemcells    []struct {
		Alias   string `yaml:"alias,omitempty"`
		OS      string `yaml:"os,omitempty"`
		Version string `yaml:"version,omitempty"`
		Name    string `yaml:"name,omitempty"`
	} `yaml:"stemcells,omitempty"`
	Update         interface{} `yaml:"update,omitempty"`
	Networks       interface{} `yaml:"networks,omitempty"`
	InstanceGroups interface{} `yaml:"instance_groups,omitempty"`
	ResourcePools  []struct {
		Name            interface{} `yaml:"name"`
		Network         interface{} `yaml:"network"`
		Size            interface{} `yaml:"size,omitempty"`
		CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
		Env             interface{} `yaml:"env,omitempty"`
		Stemcell        struct {
			Name    string `yaml:"name,omitempty"`
			Version string `yaml:"version,omitempty"`
			Alias   string `yaml:"alias,omitempty"`
			OS      string `yaml:"os,omitempty"`
		} `yaml:"stemcell"`
	} `yaml:"resource_pools,omitempty"`
	Jobs       interface{} `yaml:"jobs,omitempty"`
	Properties interface{} `yaml:"properties,omitempty"`
	Releases   []struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"releases"`
}

func (c Client) ResolveManifestVersions(manifestYAML []byte) ([]byte, error) {
	m := manifest{}
	err := yaml.Unmarshal(manifestYAML, &m)
	if err != nil {
		return nil, err
	}

	for i, r := range m.Releases {
		if r.Version == "latest" {
			release, err := c.Release(r.Name)
			if err != nil {
				return nil, err
			}
			r.Version = release.Latest()
			m.Releases[i] = r
		}
	}

	for i, pool := range m.ResourcePools {
		if pool.Stemcell.Version == "latest" {
			stemcell, err := c.Stemcell(pool.Stemcell.Name)
			if err != nil {
				return nil, err
			}
			pool.Stemcell.Version, err = stemcell.Latest()
			if err != nil {
				return nil, err
			}
			m.ResourcePools[i] = pool
		}
	}

	for i, stemcell := range m.Stemcells {
		if stemcell.Version == "latest" {
			stemcell, err := c.Stemcell(stemcell.Name)
			if err != nil {
				return nil, err
			}
			stemcellVersion, err := stemcell.Latest()
			if err != nil {
				return nil, err
			}
			m.Stemcells[i].Version = stemcellVersion
		}
	}
	return yaml.Marshal(m)
}
