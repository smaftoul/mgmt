// Mgmt
// Copyright (C) 2013-2018+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gapi

import (
	"io/ioutil"

	"github.com/purpleidea/mgmt/resources"
	"github.com/purpleidea/mgmt/util"

	errwrap "github.com/pkg/errors"
)

// Umask is the default to use when none has been specified.
// TODO: apparently using 0666 is equivalent to respecting the current umask
const Umask = 0666

// CopyFileToFs copies a file from src path on the local fs to a dst path on fs.
func CopyFileToFs(fs resources.Fs, src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return errwrap.Wrapf(err, "can't read from file `%s`", src)
	}
	if err := fs.WriteFile(dst, data, Umask); err != nil {
		return errwrap.Wrapf(err, "can't write to file `%s`", dst)
	}
	return nil
}

// CopyStringToFs copies a file from src path on the local fs to a dst path on
// fs.
func CopyStringToFs(fs resources.Fs, str, dst string) error {
	if err := fs.WriteFile(dst, []byte(str), Umask); err != nil {
		return errwrap.Wrapf(err, "can't write to file `%s`", dst)
	}
	return nil
}

// CopyDirToFs copies a dir from src path on the local fs to a dst path on fs.
func CopyDirToFs(fs resources.Fs, src, dst string) error {
	return util.CopyDiskToFs(fs, src, dst, false)
}
