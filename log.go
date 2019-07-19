package good

type dummyLogger struct{}

func (d *dummyLogger) Debug(args ...interface{})                 {}
func (d *dummyLogger) Debugf(format string, args ...interface{}) {}

func (d *dummyLogger) Info(args ...interface{})                 {}
func (d *dummyLogger) Infof(format string, args ...interface{}) {}

func (d *dummyLogger) Warn(args ...interface{})                 {}
func (d *dummyLogger) Warnf(format string, args ...interface{}) {}

func (d *dummyLogger) Error(args ...interface{})                 {}
func (d *dummyLogger) Errorf(format string, args ...interface{}) {}

// DefaultLogger is a dummy logger.
var DefaultLogger Logger = &dummyLogger{}

// Logger represents the standard logging interface
// used by the library. It matches the interface of
// the logrus logger.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}
