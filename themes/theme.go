/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

type Theme struct {
        Name           string
        Type           string  //system or local theme
        GtkTheme       string
        IconTheme      string
        CursorTheme    string
        FontName       string
        BackgroundFile string
        BasePath       string
        path           string
}

func newTheme(path string) *Theme {
        m := &Theme{}

        m.path = path
        m.setPropName("GtkTheme")
        m.setPropName("IconTheme")
        m.setPropName("CursorTheme")
        m.setPropName("FontName")
        m.setPropName("BackgroundFile")

        m.listenSettingsChanged()

        return m
}