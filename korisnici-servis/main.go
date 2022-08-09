package main

import (
	kontroler "korisnici-servis/kontroler"
	util "korisnici-servis/util"
)

func main() {
	util.KonektujPopuniDB()
	kontroler.OtkrijEndpointe()

}
