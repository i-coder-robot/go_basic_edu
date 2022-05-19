package book

import "errors"

func ShowBookInfo(bookName, authorName string) (string, error) {
	if bookName == "" {
		return "", errors.New("图书名称不能为空")
	}
	if authorName == "" {
		return "", errors.New("作者名称不能为空")
	}
	return bookName + ",作者:" + authorName, nil
}
