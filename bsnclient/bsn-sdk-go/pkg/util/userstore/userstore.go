package userstore

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/core/entity/msp"
)

type UserStore interface {
	Load(user *msp.UserData) error
	LoadAll(appCode string) []*msp.UserData
	Store(user *msp.UserData) error
}
