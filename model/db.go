package model

import (
	"errors"

	"xorm.io/xorm"
)

type databaseModel struct {
	engine *xorm.Engine
	// mu     sync.Mutex
}

func (d databaseModel) setEngine(engine *xorm.Engine) {
	// d.mu.Lock()
	// defer d.mu.Unlock()
	d.engine = engine
}

func (d databaseModel) getEngine() (*xorm.Engine, error) {
	// d.mu.Lock()
	// defer d.mu.Unlock()
	if d.engine == nil {
		return nil, errors.New("DB engine is not set")
	}
	return d.engine, nil
}

func (d databaseModel) getSession() (*xorm.Session, error) {
	// d.mu.Lock()
	// defer d.mu.Unlock()
	if d.engine == nil {
		return nil, errors.New("DB engine is not set")
	}
	return d.engine.NewSession(), nil
}

func SetDatabase(engine *xorm.Engine) {
	if engine == nil {
		return
	}
	// Set engine for each model
}

func SyncDatabase(engine *xorm.Engine) {
	if engine == nil {
		return
	}
	// Sync database for each model
}
