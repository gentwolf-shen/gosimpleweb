package dao

import (
	"github.com/gentwolf-shen/gohelper/gomybatis"
	"github.com/gentwolf-shen/gohelper/logger"
)

type BaseDao struct{}

func (this *BaseDao) List(selector string, p map[string]interface{}) []map[string]string {
	rows, err := gomybatis.Query(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return rows
}

func (this *BaseDao) Row(selector string, p map[string]interface{}) map[string]string {
	row, err := gomybatis.QueryRow(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return row
}

func (this *BaseDao) Update(selector string, p map[string]interface{}) error {
	_, err := gomybatis.Update(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return err
}

func (this *BaseDao) Delete(selector string, p map[string]interface{}) error {
	_, err := gomybatis.Delete(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return err
}

func (this *BaseDao) Scalar(selector string, p map[string]interface{}) string {
	str, err := gomybatis.QueryScalar(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return str
}

func (this *BaseDao) Create(selector string, p map[string]interface{}) int64 {
	id, err := gomybatis.Insert(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}
	return id
}

func (this *BaseDao) UpdateExt(selector string, p map[string]interface{}) int64 {
	n, err := gomybatis.Update(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return n
}

func (this *BaseDao) DeleteExt(selector string, p map[string]interface{}) int64 {
	n, err := gomybatis.Delete(selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return n
}

func (this *BaseDao) ListObject(value interface{}, selector string, p map[string]interface{}) error {
	err := gomybatis.QueryObjects(value, selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}

	return err
}

func (this *BaseDao) RowObject(value interface{}, selector string, p map[string]interface{}) error {
	err := gomybatis.QueryObject(value, selector, p)
	if err != nil {
		logger.Error(selector)
		logger.Error(err)
	}
	return err
}
