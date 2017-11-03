// +build appengine

package service

import (
	"errors"
	"os/exec"
	"time"
)

var ErrNotAvailable = errors.New("this is not available in App Engine")

type Service struct {
	URLTemplate string
	CmdTemplate []string
	url         string
	command     *exec.Cmd
}

type addressInfo struct {
	Address string
	Host    string
	Port    string
}

func (s *Service) URL() string {
	return s.url
}

func (s *Service) Start(debug bool) error { return ErrNotAvailable }

func (s *Service) Stop() error { return ErrNotAvailable }

func freeAddress() (addressInfo, error) { return addressInfo{}, ErrNotAvailable }

func (s *Service) WaitForBoot(timeout time.Duration) error { return ErrNotAvailable }

func (s *Service) checkStatus() bool { return false }
