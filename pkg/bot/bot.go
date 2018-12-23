package bot

import (
	"image/draw"
	"log"
	"time"

	"github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev/fb"

	"github.com/miry/ev3dev/pkg/dev"
)

type Bot struct {
	MMotor *dev.Motor
	LMotor *dev.Motor
	RMotor *dev.Motor
}

func New() *Bot {
	bot := &Bot{}
	bot.initScreen()
	err := bot.initMotors()
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func (b *Bot) Run() {
	for i := 0; i < 2; i++ {
		b.Draw(gopher)
		b.Forward(70)
		b.Draw(gopherSquint)
		b.Steering(50)
	}
}

func (b *Bot) initScreen() {
	ev3.LCD.Init(true)
}

func (b *Bot) initMotors() error {
	var err error

	b.MMotor, err = dev.New("outA", "lego-ev3-m-motor")
	if err != nil {
		return err
	}
	b.LMotor, err = dev.New("outB", "lego-ev3-l-motor")
	if err != nil {
		return err
	}
	b.RMotor, err = dev.New("outC", "lego-ev3-l-motor")
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) Forward(power int) {
	b.MoveTank(power, power)
}

func (b *Bot) Steering(power int) {
	b.MoveTank(power, -power)
}

func (b *Bot) MoveTank(right, left int) {
	b.RMotor.SetSpeed(right)
	b.LMotor.SetSpeed(left)
	dev.CheckErrors(b.RMotor, b.LMotor)
	time.Sleep(2 * time.Second)
	b.RMotor.Stop()
	b.LMotor.Stop()
	dev.CheckErrors(b.RMotor, b.LMotor)
}

func (b *Bot) Draw(picture *fb.Monochrome) {
	draw.Draw(ev3.LCD, ev3.LCD.Bounds(), picture, picture.Bounds().Min, draw.Src)
}

func (b *Bot) Exit() {
	b.RMotor.Stop()
	b.LMotor.Stop()
	b.MMotor.Stop()
	ev3.LCD.Close()
}
