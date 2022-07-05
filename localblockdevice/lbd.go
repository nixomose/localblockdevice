// SPDX-License-Identifier: LGPL-2.1
// Copyright (C) 2021-2022 stu mark

package main

import (
	"container/list"
	"os"

	blockdevicelib "github.com/nixomose/blockdevicelib/blockdevicelib"
	"github.com/nixomose/nixomosegotools/tools"
	"github.com/spf13/cobra"
)

const TXT_APPLICATION_NAME = "lbd"

const TXT_DEFAULT_CONFIG_FILE = "/etc/localblockdevice/localblockdevice.cf"
const TXT_DEFAULT_LOG_FILE = "/var/log/localblockdevice/localblockdevice.log"
const TXT_DEFAULT_CATALOG_FILE = "/etc/localblockdevice/catalog.toml"

/* this is the main binary entry point for local block device that uses stree_v to store the block device data
on a local disk or file. */

func main() {

	// var params = os.Args[1:]

	var ret tools.Ret
	var l *blockdevicelib.Lbd_lib
	ret, l = blockdevicelib.New_blockdevicelib(TXT_APPLICATION_NAME)
	if ret != nil {
		os.Exit(1)
		return
	}
	var root_cmd *cobra.Command
	ret, root_cmd = l.Startup(TXT_DEFAULT_CONFIG_FILE, TXT_DEFAULT_LOG_FILE,
		TXT_DEFAULT_CATALOG_FILE) // start configuring and make log and stuff.
	if ret != nil {
		os.Exit(1)
		return
	}
	if ret = l.Run(root_cmd, &list.List{}); ret != nil {
		os.Exit(1)
		return
	}
}
