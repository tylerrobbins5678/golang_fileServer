package service

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const fileroot = "C://served_files"

func GetNames(c *gin.Context) {
	// get path and check for files and folders at path
	path := c.Param("path")
	var dirArray []string
	var fileArray []string

	dirStat, err := os.Stat(fileroot + path)
	// check if path or file exist
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": path + " not found"})
		return
	}

	if dirStat.IsDir() {
		dirs, _ := ioutil.ReadDir(fileroot + path)

		for _, obj := range dirs {
			if obj.IsDir() {
				dirArray = append(dirArray, obj.Name())
			} else {
				fileArray = append(fileArray, obj.Name())
			}
		}
		// add dirs to response
		c.JSON(http.StatusOK, gin.H{
			"dirs":  dirArray,
			"files": fileArray,
		})

		// return file
	} else {
		completePathSplit := strings.Split(path, "/")

		path = strings.Join(completePathSplit[:len(completePathSplit)-1], "/") + "/"
		file := completePathSplit[len(completePathSplit)-1]
		c.Header("Content-Type", "application/octet-stream")
		//Force browser download
		c.Header("Content-Disposition", "attachment; filename="+file)
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Cache-Control", "no-cache")
		c.File(fileroot + path + file)
		return
	}

}

func GetByName(c *gin.Context) {

}

func Create(c *gin.Context) {
	file, err := c.FormFile("file")
	path := c.Param("path")
	if err != nil {
		// path not found
		return
	}
	// if path does not exist
	if _, err := os.Stat(path); err != nil {
		os.MkdirAll("/served_files/"+path, os.ModePerm)
	}
	if err = c.SaveUploadedFile(file, fileroot+path+"/"+file.Filename); err != nil {
		// generic save error
		return
	}

}
