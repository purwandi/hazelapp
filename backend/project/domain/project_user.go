package domain

const (
	MaintainerAccess = "MAINTAINER"
	DeveloperAccess  = "DEVELOPER"
	ReporterAccess   = "REPORTER"
	GuestAccess      = "GUEST"
)

type ProjectUser struct {
	ProjectID   int
	UserID      int
	AccessLevel string
}

func (p *Project) AttachMaintainer(userID int) (*ProjectUser, error) {
	return &ProjectUser{
		ProjectID:   p.ID,
		UserID:      userID,
		AccessLevel: MaintainerAccess,
	}, nil
}

func (p *Project) AttachDeveloper(userID int) (*ProjectUser, error) {
	return &ProjectUser{
		ProjectID:   p.ID,
		UserID:      userID,
		AccessLevel: DeveloperAccess,
	}, nil
}

func (p *Project) AttachReporter(userID int) (*ProjectUser, error) {
	return &ProjectUser{
		ProjectID:   p.ID,
		UserID:      userID,
		AccessLevel: ReporterAccess,
	}, nil
}

func (p *Project) AttachGuest(userID int) (*ProjectUser, error) {
	return &ProjectUser{
		ProjectID:   p.ID,
		UserID:      userID,
		AccessLevel: GuestAccess,
	}, nil
}
