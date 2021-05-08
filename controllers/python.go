package controller

import (
	model "go-run-python/models"
	util "go-run-python/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

type PythonController struct{}

// RunProgram :
func (pc PythonController) RunProgram(c *gin.Context) {
	// get data
	type RequestData struct {
		Code  string `json:"code"`
		Input string `json:"in"`
	}
	req := RequestData{}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.UnprocessableLog(c, err)
		return
	}

	// declare model
	m := new(model.PythonModel)

	// run program
	out, time, err := m.RunProgram(req.Code, req.Input)

	// return output
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"output":   out,
			"timeused": -1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"output":   out,
		"timeused": time,
	})
}

// EvaluateProgram :
func (pc PythonController) EvaluateProgram(c *gin.Context) {
	// get data
	type RequestData struct {
		Code      string `json:"code"`
		ProblemID uint64 `json:"problem_id"`
	}
	req := RequestData{}
	if err := c.ShouldBindJSON(&req); err != nil {
		util.UnprocessableLog(c, err)
		return
	}

	// declare model
	m := new(model.PythonModel)

	// get test cases
	cases, err := m.GetTestCases(req.ProblemID)
	if err != nil {
		util.UnprocessableLog(c, err)
		return
	}

	// evaluate program
	out, err := m.EvaluateProgram(req.Code, cases)

	// return output
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"output": out,
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"output": out,
		"status": true,
	})
}
