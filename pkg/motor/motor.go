package motor

import (
	"fmt"
	"log"
	"time"

	"github.com/ev3go/ev3dev"
)

type Motor struct {
	*ev3dev.TachoMotor
	MaxSpeed int
}

func New(port, deviceName string) (*Motor, error) {
	var err error
	var motor *ev3dev.TachoMotor
	// Get the handle for the medium motor on outA.
	motor, err = ev3dev.TachoMotorFor("ev3-ports:"+port, deviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to find %s motor on %s: %v", deviceName, port, err)
	}
	err = motor.SetStopAction("brake").Err()
	if err != nil {
		return nil, fmt.Errorf("failed to set brake stop for %s motor on %s: %v", deviceName, port, err)
	}
	speed := motor.MaxSpeed()

	result := &Motor{
		motor,
		speed,
	}
	return result, nil
}

func (m *Motor) Run(speed int) {
	m.SetSpeed(speed)
	time.Sleep(time.Second / 2)
	m.Stop()
	CheckErrors(m)
}

func (m *Motor) SetSpeed(speed int) {
	m.SetSpeedSetpoint(speed * m.MaxSpeed / 100).Command("run-forever")
}

func (m *Motor) Stop() {
	m.Command("stop")
}

func CheckErrors(devs ...*Motor) {
	for _, d := range devs {
		err := d.TachoMotor.Err()
		if err != nil {
			drv, dErr := ev3dev.DriverFor(d)
			if dErr != nil {
				drv = fmt.Sprintf("(missing driver name: %v)", dErr)
			}
			addr, aErr := ev3dev.AddressOf(d)
			if aErr != nil {
				drv = fmt.Sprintf("(missing port address: %v)", aErr)
			}
			log.Fatalf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
		}
	}
}
