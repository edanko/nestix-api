package version

import (
	"runtime/debug"
	"sync"
	"time"
)

type Version struct {
	// Version will be the version tag if the binary is built with "go install url/tool@version".
	// If the binary is built some other way, it will be "(devel)".
	Version string `json:"version"`
	// Revision is taken from the vcs.revision tag.
	Revision string `json:"revision"`
	// LastCommit is taken from the vcs.time tag.
	LastCommit time.Time `json:"lastCommit"`
	// DirtyBuild is taken from the vcs.modified tag.
	DirtyBuild bool `json:"dirtyBuild"`
}

var instance *Version
var once sync.Once

func GetVersion() *Version {
	once.Do(func() {
		instance = &Version{
			Version:    "unknown",
			Revision:   "unknown",
			DirtyBuild: true,
		}

		info, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}

		instance.Version = info.Main.Version

		for _, kv := range info.Settings {
			switch kv.Key {
			case "vcs.revision":
				instance.Revision = kv.Value
			case "vcs.time":
				instance.LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
			case "vcs.modified":
				instance.DirtyBuild = kv.Value == "true"
			}
		}
	})
	return instance
}
