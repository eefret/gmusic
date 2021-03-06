/*
Copyright 2014 Kaissersoft Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gmusic

import (
	"errors"
)

//=======================================================================
//                          Errors
//=======================================================================

var (
	//ErrInvalidCred is launched when an invalid Email or Password is given.
	ErrInvalidCred = errors.New("Invalid Email or Password please try again..")
	//ErrNoTranslation is launched when there's no translation available
	ErrNoTranslation = errors.New("theres no translation")
)
