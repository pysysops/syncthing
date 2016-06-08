// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package db

import (
	"fmt"

	"github.com/syncthing/syncthing/lib/protocol"
)

func (f FileInfoTruncated) String() string {
	return fmt.Sprintf("File{Name:%q, Permissions:0%o, Modified:%d, Version:%v, Length:%d, Deleted:%v, Invalid:%v, NoPermissions:%v}",
		f.Name, f.Permissions, f.Modified, f.Version, f.Size, f.Deleted, f.Invalid, f.NoPermissions)
}

func (f FileInfoTruncated) IsDeleted() bool {
	return f.Deleted
}

func (f FileInfoTruncated) IsInvalid() bool {
	return f.Invalid
}

func (f FileInfoTruncated) IsDirectory() bool {
	return f.Type == protocol.FileInfoTypeDirectory
}

func (f FileInfoTruncated) IsSymlink() bool {
	switch f.Type {
	case protocol.FileInfoTypeSymlinkDirectory, protocol.FileInfoTypeSymlinkFile, protocol.FileInfoTypeSymlinkUnknown:
		return true
	default:
		return false
	}
}

func (f FileInfoTruncated) HasPermissionBits() bool {
	return !f.NoPermissions
}

func (f FileInfoTruncated) FileLength() int64 {
	if f.IsDirectory() || f.IsDeleted() {
		return 128
	}
	return f.Size
}

func (f FileInfoTruncated) FileName() string {
	return f.Name
}
