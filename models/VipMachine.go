package models

import (
	//"fmt"
	//"github.com/go-xorm/xorm"
	//_ "github.com/lunny/godbc"
	//"os"
	"time"
)

/*
USE [QPAccountsDB]
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[VipMachine](
	[MachineSerial] [nvarchar](32) NOT NULL,
	[OverDate] [datetime] NULL,
	[Memo] [ntext] NULL,
	[AddTime] [datetime] NULL,
	[IsBlack] [int] NULL,
	[UserID] [int] NULL,
 CONSTRAINT [PK_VipMachine_MachineSerial] PRIMARY KEY CLUSTERED
(
	[MachineSerial] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]

GO

EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'机器序列' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'VipMachine', @level2type=N'COLUMN',@level2name=N'MachineSerial'
GO

EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'过期时间' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'VipMachine', @level2type=N'COLUMN',@level2name=N'OverDate'
GO
*/
type VipMachine struct {
	MachineSerial string    `xorm:"pk not null 'MachineSerial'"`
	UserID        int       `xorm:"not null UserID"`
	OverDate      time.Time `xorm:"not null OverDate"`
	Memo          string    `xorm:"not null Memo"`
	IsBlack       int       `xorm:"not null IsBlack"`
	UpdateUser    string    `xorm:"not null UpdateUser"`
	AddTime       time.Time `xorm:"AddTime"`
	UpdateTime    time.Time `xorm:"UpdateTime"`
}

func GetMachineByMachineSerial(MachineSerial string) (bool, VipMachine) {
	var machine VipMachine

	return false, machine
}

func GetVipMachineByUid(uid int64) (bool, VipMachine) {

	var machine VipMachine

	err, _ := engine.Alias("vm").Where("vm.UserID = ?", uid).Get(&machine)
	return err, machine
}

func GetAll(username string) (bool, VipMachine) {

	var machine VipMachine

	return false, machine
}
