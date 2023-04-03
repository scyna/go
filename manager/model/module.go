package model

import (
	"time"

	proto "github.com/scyna/go/manager/.proto/generated"
)

type Module struct {
	Organization string `db:"org_code"`
	Code         string `db:"code"`
	Description  string `db:"description"`
	Secret       string `db:"secret"`
}

func (m *Module) FromDTO(o *proto.Module) {
	m.Code = o.Code
	m.Organization = o.Organization
	m.Description = o.Description
	m.Secret = o.Secret
}

func (m *Module) ToDTO() *proto.Module {
	return &proto.Module{
		Organization: m.Organization,
		Code:         m.Code,
		Description:  m.Description,
		Secret:       "", // NO RETURN
	}
}

type Session struct {
	ID         uint64    `db:"id"`
	ModuleCode string    `db:"module_code"`
	End        time.Time `db:"end"`
	ExitCode   uint32    `db:"exit_code"`
	LastUpdate time.Time `db:"last_update"`
	Start      time.Time `db:"start"`
}

func (s *Session) ToDTO() *proto.Session {
	raw := proto.Session{
		Id:         s.ID,
		ModuleCode: s.ModuleCode,
		Start:      s.Start.Format(time.RFC3339),
		LastUpdate: s.LastUpdate.Format(time.RFC3339),
		ExitCode:   s.ExitCode,
	}

	if s.End.IsZero() {
		raw.End = ""
		raw.Status = "RUNNING - " + time.Since(s.Start).String() + " ago"
	} else {
		raw.End = s.End.Format(time.RFC3339)
		raw.Status = "STOPPED - " + time.Since(s.LastUpdate).String() + " ago"
	}

	return &raw
}
