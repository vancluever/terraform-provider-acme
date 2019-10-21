// Copyright 2018. Akamai Technologies, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edgegrid 

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logBuffer *bufio.Writer
var LogFile *os.File

func SetupLogging(log *logrus.Logger) {
	log.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation:    true,
		EnvironmentOverrideColors: true,
	})
        // Log file destination specified? If not, use default stdout
        if logFileName := os.Getenv("AKAMAI_LOG_FILE"); logFileName != "" {
		// If the file doesn't exist, create it, or append to the file
		LogFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
                log.SetOutput(LogFile)
	}

	log.SetLevel(logrus.PanicLevel)
	if logLevel := os.Getenv("AKAMAI_LOG"); logLevel != "" {
		level, err := logrus.ParseLevel(logLevel)
		if err == nil {
			log.SetLevel(level)
		} else {
			log.Warningln("[WARN] Unknown AKAMAI_LOG value. Allowed values: panic, fatal, error, warn, info, debug, trace")

		}
	}
}

func LogMultiline(f func(args ...interface{}), args ...string) {
	for _, str := range args {
		for _, str := range strings.Split(strings.Trim(str, "\n"), "\n") {
			f(str)
		}
	}
}

func LogMultilineln(f func(args ...interface{}), args ...string) {
	LogMultiline(f, args...)
}

func LogMultilinef(f func(formatter string, args ...interface{}), formatter string, args ...interface{}) {
	str := fmt.Sprintf(formatter, args...)
	for _, str := range strings.Split(strings.Trim(str, "\n"), "\n") {
		f(str)
	}
}
