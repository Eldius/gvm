//go:build linux && amd64

package versions

func (v *GoVersion) GetURL() string {
	return v.LinuxAmd64
}
