package provision

import (
	"github.com/docker/machine/drivers"
)

const (
	// TODO: eventually the RPM install process will be integrated
	// into the get.docker.com install script; for now
	// we install via vendored RPMs
	dockerCentosRPMPath = "https://get.docker.com/rpm/1.7.0/centos-7/RPMS/x86_64/docker-engine-1.7.0-1.el7.centos.x86_64.rpm"
)

func init() {
	Register("Centos", &RegisteredProvisioner{
		New: NewCentosProvisioner,
	})
}

func NewCentosProvisioner(d drivers.Driver) Provisioner {
	g := GenericProvisioner{
		DockerOptionsDir:  "/etc/docker",
		DaemonOptionsFile: "/etc/systemd/system/docker.service",
		OsReleaseId:       "centos",
		Packages:          []string{},
		Driver:            d,
	}
	p := &CentosProvisioner{
		RedHatProvisioner{
			GenericProvisioner: g,
			DockerRPMPath:      dockerCentosRPMPath,
		},
	}
	return p
}

type CentosProvisioner struct {
	RedHatProvisioner
}
