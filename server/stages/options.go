package stages

import (
	"gitlab.iyorhe.com/wfgz/reverseproxy/service/stores"
	"gitlab.iyorhe.com/wfgz/reverseproxy/service/stores/file"
)

const (
	//DIR 目录路径
	DIR = "data/stages"
)

//Options ...
type Options struct {
	Store stores.Store
}

//Option ...
type Option func(*Options)

func defaultOptions() Options {
	opts := Options{}
	DefaultStore(&opts)
	return opts
}

//DefaultStore 默认存储
func DefaultStore(opts *Options) {
	opts.Store = file.NewFileStore(file.Dir(DIR))
}
