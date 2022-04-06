package main

import (
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
	"strings"
	"sync"
)

var m sync.RWMutex

func bootstrap() *cobra.Command {
	return &cobra.Command{
		Use:   "bootstrap",
		Short: "ipfs bootstrap all node",
		Long:  `ipfs bootstrap all node`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ipfsColony(args)
			return nil
		},
	}
}
func ipfsColony(node []string) {
	// 获取所有的id
	var wg sync.WaitGroup
	addr := make([]string, 0, len(node))
	for _, v := range node {
		wg.Add(1)
		go func(ipAddr string) {
			sh := shell.NewShell(ipAddr + ":5001")
			idOut, _ := sh.ID()
			a := strings.Replace(idOut.Addresses[0], "127.0.0.1", ipAddr, 1)
			m.Lock()
			addr = append(addr, a)
			m.Unlock()
			wg.Done()
		}(v)

	}
	wg.Wait()
	for _, v := range node {
		wg.Add(1)
		go func(ipAddr string) {
			sh := shell.NewShell(ipAddr + ":5001")
			sh.BootstrapRmAll()
			sh.BootstrapAdd(addr)
			fmt.Println(addr)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
