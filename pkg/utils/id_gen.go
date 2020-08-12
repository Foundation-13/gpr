package utils

import "github.com/gofrs/uuid"

//go:generate mockery -name IDGen -outpkg utilmocks -output ./utilsmocks -dir .
type IDGen interface {
	StringID() string
}

// impl

type UUIDIDGen struct {}

func (UUIDIDGen) StringID() string {
	return uuid.Must(uuid.NewV4()).String()
}


