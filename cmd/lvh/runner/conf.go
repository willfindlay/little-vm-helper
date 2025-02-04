package runner

import (
	"github.com/sirupsen/logrus"
)

type RunConf struct {
	// Image filename
	Image string
	// kernel filename to boot with. (if empty no -kernel option will be passed to qemu)
	KernelFname string
	// Do not run the qemu command, just print it
	QemuPrint bool
	// Do not use KVM acceleration, even if /dev/kvm exists
	DisableKVM bool
	// Daemonize QEMU after initializing
	Daemonize bool

	// Disable the network connection to the VM
	DisableNetwork bool
	ForwardedPorts []PortForward

	Logger *logrus.Logger

	HostMount string

	SerialPort int
}

func (rc *RunConf) testImageFname() string {
	return rc.Image
}

type PortForward struct {
	HostPort int
	VMPort   int
	Protocol string
}
