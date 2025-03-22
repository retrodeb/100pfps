package model

import (
	"net"

	"github.com/jackc/pgtype"
	"github.com/retrodeb/100pfps/db"
)

type Ban struct {
	Ip pgtype.Inet `gorm:"primaryKey;type:inet;not null"`
}

func BanIP(ip net.IP) (err error) {
	var ban Ban
	ban.Ip.Set(ip)
	return db.GetDB().Create(&ban).Error
}

func RemoveBan(ip net.IP) (err error) {
	var ban Ban
	ban.Ip.Set(ip)
	return db.GetDB().Delete(&ban).Error
}

func IsBanned(ip net.IP) (banned bool) {
	var netip net.IPNet
	netip.IP = ip
	if ip.To4() != nil {
		netip.Mask = createMask(32)
	} else {
		netip.Mask = createMask(70)
	}
	var count int64
	db.GetDB().Model(&Ban{}).Where("ip <<= ?", netip.String()).Limit(1).Count(&count)
	return count > 0
}

func createMask(bits uint8) (mask net.IPMask) {
	var size int
	if bits <= 32 {
		size = 4
	} else {
		size = 16
	}

	mask = make(net.IPMask, size)
	for i := uint8(0); i < bits; i++ {
		mask[i/8] |= 1 << (7 - (i % 8))
	}
	return
}
