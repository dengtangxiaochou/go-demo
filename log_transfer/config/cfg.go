package config

type LogTransfer struct {
	KafkaCfg `ini:"kafka"`
	EsCfg    `ini:"es"`
}

//KafkaCfg
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

//EsCfg
type EsCfg struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chanSize"`
	Nums     int    `ini:"nums"`
}
