package main

import o "design_pattern/structure_pattern/adapter/a"

func main() {
	client := &o.Client{}
	windows := &o.Window{}
	client.InsertSquereUSBinComputer(windows)
	macadapter := &o.MacAdapter{
		M: o.Mac{},
	}
	client.InsertSquereUSBinComputer(macadapter)
}