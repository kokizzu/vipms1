package model

import (
	"testing"
	"vipms1/app/config"
)

func TestVip(t *testing.T) {
	db, err := config.ConnectMysql()
	if err != nil {
		t.Error(err)
	}
	
	// test insert
	vip := Vip{
		Name: `test`,
		Country: `test2`,
	}
	id, err := vip.Insert(db)
	if id == 0 {
		t.Error(`failed to insert`)
	}
	if err != nil {
		t.Error(err)
	}
	
	// test select
	vip2 := Vip{}
	err = vip2.GetById(db, id)
	if err != nil {
		t.Error(err)
	}
	
	if vip2.Name != vip.Name {
		t.Error(`inserted vip name not equal`)
	}
	if vip2.Country != vip.Country {
		t.Error(`inserted vip country not equal`)
	}
	
	// test update
	err = vip2.SetArrived(db, id)
	if err != nil {
		t.Error(err)
	}
	
	// test really arrived
	err = vip.GetById(db, vip2.ID)
	if err != nil {
		t.Error(err)
	}
	
	if !vip.Arrived {
		t.Error(`failed set arrived`)
	} 
}
