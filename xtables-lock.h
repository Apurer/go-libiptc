/*
 * go-libiptc v0.3.1 - libiptc bindings for Go language
 * Copyright (C) 2015~2016 apurer - https://github.com/apurer/go-libiptc/

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
*/

// these headers needs to be here so that Go will import the externs
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

int get_errno();
void reset_errno();

int xtables_lock(bool wait, uint max_seconds_wait);
int xtables_unlock();
