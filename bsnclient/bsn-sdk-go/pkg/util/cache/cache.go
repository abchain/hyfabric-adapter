package cache

import (
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient/bsn-sdk-go/pkg/util/cache/proc"
	"time"
)

const (
	Expiration = "0.5h"
	Interval   = "3s"
)

var ca *proc.Cache

func InitCache() {
	// 默认过期时间
	defaultExpiration, _ := time.ParseDuration(Expiration)
	//回收时间,回收已过期
	gcInterval, _ := time.ParseDuration(Interval)
	ca = proc.NewCache(defaultExpiration, gcInterval)

}

//设置一个滑动过期的缓存 @expiration 过期时间
func SetSlideValue(key string, value interface{}, expiration string) {
	if ca == nil {
		InitCache()
	}
	dur, _ := time.ParseDuration(expiration)
	ca.Set(key, value, dur, true)
}

//设置一个带过期时间的缓存
func SetValueByExpiration(key string, value interface{}, expiration string) {
	if ca == nil {
		InitCache()
	}
	expi, _ := time.ParseDuration(expiration)
	ca.Set(key, value, expi, false)
}

//设置一个永不过期的缓存
func SetValue(key string, value interface{}) {
	if ca == nil {
		InitCache()
	}
	ca.Set(key, value, 0, false)
}

//获取设置的值
func GetValue(key string) (interface{}, bool) {
	if ca == nil {
		return nil, false
	}
	return ca.Get(key)

}

func ClearKey(key string) {
	ca.Delete(key)
}
