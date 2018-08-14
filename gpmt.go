/* Greenplum magic tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

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

package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pivotal-gss/gpmt2/cmd"
	"io"
	"os"
	"time"
)

// TODO: maybe clean this up by adding a logrus file hook
// https://github.com/Sirupsen/logrus/issues/230
func init() {
	logFilename := fmt.Sprintf("/tmp/gpmt_log_%s", time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.WithField("file", logFilename).Error("could not open file - please please ensure user has write access")
		log.Error("continuing with stdout only")
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}

func main() {
	cmd.Execute()
}
