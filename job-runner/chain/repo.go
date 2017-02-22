// Copyright 2017, Square, Inc.

package chain

type Repo interface {
	Add(*Chain) error
	Remove(uint) error
	Set(*Chain) error
}

type FakeRepo struct{}

func (f *FakeRepo) Add(chain *Chain) error {
	return nil
}

func (f *FakeRepo) Remove(id uint) error {
	return nil
}

func (f *FakeRepo) Set(chain *Chain) error {
	return nil
}
