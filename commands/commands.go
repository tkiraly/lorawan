package commands

type Fopter interface {
	ByteArray() []byte
	Len() uint8
	String() string
}

const (
	//LinkCheckReqCommand is the CID of LinkCheckReq
	LinkCheckReqCommand byte = iota + 2
	//LinkADRAnsCommand is the CID of LinkADRAns
	LinkADRAnsCommand
	//DutyCycleAnsCommand is the CID of DutyCycleAns
	DutyCycleAnsCommand
	//RXParamSetupAnsCommand is the CID of RXParamSetupAns
	RXParamSetupAnsCommand
	//DevStatusAnsCommand is the CID of DevStatusAns
	DevStatusAnsCommand
	//NewChannelAnsCommand is the CID of NewChannelAns
	NewChannelAnsCommand
	//RXTimingSetupAnsCommand is the CID of RXTimingSetupAns
	RxTimingSetupAnsCommand
	//TxParamSetupAnsCommand is the CID of TxParamSetupAns
	TxParamSetupAnsCommand
)

const (
	//LinkCheckAnsCommand is the CID of LinkCheckAns
	LinkCheckAnsCommand byte = iota + 2
	//LinkADRReqCommand is the CID of LinkADRReq
	LinkADRReqCommand
	//DutyCycleReqCommand is the CID of DutyCycleReq
	DutyCycleReqCommand
	//RXParamSetupReqCommand is the CID of RXParamSetupReq
	RXParamSetupReqCommand
	//DevStatusReqCommand is the CID of DevStatusReq
	DevStatusReqCommand
	//NewChannelReqCommand is the CID of NewChannelReq
	NewChannelReqCommand
	//RxTimingSetupReqCommand is the CID of TXTimingSetupReq
	RxTimingSetupReqCommand
	//TxParamSetupReqCommand is the CID of TxParamSetupReq
	TxParamSetupReqCommand
)
