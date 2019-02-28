// date: 2019-02-28
package sys

import (
	log "github.com/jeanphorn/log4go"
)

func init() {
	log.LoadConfiguration("./log4go.json")
}
