package handlers

import (
	"errors"
	"ssug/internal/utils"
	"ssug/modules/base"
	"ssug/modules/data"
)

// AddMappingHandler 处理添加短链请求
func AddMappingHandler(originalURLs []string) ([]data.Mapping, error) {

	var mappings []data.Mapping
	for _, originalURL := range originalURLs {
		if !utils.IsURL(originalURL) {
			return nil, errors.New("添加映射失败，URL非法")
		}

		ok := data.Redirect.HasOriginalURL(originalURL)
		if ok {
			m, _ := data.Redirect.GetMappingFO(originalURL)
			mappings = append(mappings, m)
			// utils.Logger.Info(fmt.Sprintf("添加映射%s -> %s失败，映射已存在", m.ShortURL, m.OriginalURL))
			// return m, errors.New("添加失败，映射已存在")
		} else {
			shortURL := base.GenValue(originalURL)
			m, err := data.Redirect.AddMapping(originalURL, shortURL)
			if err == nil {
				// utils.Logger.Info(fmt.Sprintf("成功添加映射%s -> %s", m.ShortURL, m.OriginalURL))
				mappings = append(mappings, m)
				// return m, nil
			} else {
				utils.Logger.Warning(err)
				return nil, err
			}
		}

	}
	return mappings, nil
}

// GetMappingHandler 处理原始链接获取请求
func GetMappingHandler(shortURL string) (string, error) {
	if !utils.IsLegalValue(shortURL) {
		return "", errors.New("查询映射失败，短链非法")
	}
	k, err := data.Redirect.GetMappingO(shortURL)
	if err == nil {
		return k, nil
	} else {
		return "", errors.New("查询失败，映射不存在")
	}
}
