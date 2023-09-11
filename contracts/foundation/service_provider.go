package foundation

type ServiceProvider interface {
	Boot(app Application)
	Register(app Application)
}
