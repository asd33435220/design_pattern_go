package main

func main() {
	myClient := &client{}
	myMac := &mac{}
	myClient.insertLightingConnectorIntoComputer(myMac)
	myWinodws := &windows{}
	myWindowsAdapter := &windowsAdapter{windowsMachine: myWinodws}
	myClient.insertLightingConnectorIntoComputer(myWindowsAdapter)
}
