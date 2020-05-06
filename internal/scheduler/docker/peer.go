/*
 * Copyright 2020. The Alkaid Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 *
 * Alkaid is a BaaS service based on Hyperledger Fabric.
 *
 */

package docker

import (
	"github.com/yakumioto/alkaid/internal/vm"
	dockervm "github.com/yakumioto/alkaid/internal/vm/docker"
)

type peerNode struct {
	cli *dockervm.Controller
}

func (p *peerNode) CreatePeer(crs ...*vm.CreateRequest) error {
	for _, cr := range crs {
		if err := p.cli.Create(cr); err != nil {
			logger.Errof("Create peer error: %s", err)
			return err
		}
	}
	return nil
}

func (p *peerNode) RestartPeer(ids ...string) error {
	for _, id := range ids {
		if err := p.cli.Restart(id); err != nil {
			logger.Errof("Restart peer error: %s", err)
			return err
		}
	}
	return nil
}

func (p *peerNode) StopPeer(ids ...string) error {
	for _, id := range ids {
		if err := p.cli.Stop(id); err != nil {
			logger.Errof("Stop peer error: %s", err)
			return err
		}
	}
	return nil
}

func (p *peerNode) DeletePeer(ids ...string) error {
	for _, id := range ids {
		if err := p.cli.Delete(id); err != nil {
			logger.Errof("Delete peer error: %s", err)
			return err
		}
	}
	return nil
}
