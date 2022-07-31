package main

import (
	kontroler "knjige-servis/kontroler"
	util "knjige-servis/util"
)

func main() {

	util.KonektujPopuniDB()
	kontroler.OtkrijEndpointe()
}
