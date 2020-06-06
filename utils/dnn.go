package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

type ParamsJSON struct {
	Parameters map[string][][]float64
	Scales     map[string][][]float64
}

func sigmoid(z *mat.Dense) *mat.Dense {
	wr, _ := z.Dims()
	a := mat.NewDense(wr, 1, nil)

	// all the element is 1
	o := make([]float64, wr)
	for i := range o {
		o[i] = 1
	}
	ones := mat.NewDense(wr, 1, o)

	// all the element is -1
	no := make([]float64, wr)
	for i := range no {
		no[i] = -1
	}
	nones := mat.NewDense(wr, 1, no)

	// compute sigmoid
	a.MulElem(nones, z)
	a.Exp(a)
	a.Add(ones, a)
	a.DivElem(ones, a)

	return a
}

func relu(z *mat.Dense) *mat.Dense {
	wr, _ := z.Dims()
	a := mat.NewDense(wr, 1, nil)
	for i := 0; i < wr; i++ {
		if v := z.At(i, 0); v > 0 {
			a.Set(i, 0, v)
		} else {
			a.Set(i, 0, 0)
		}
	}

	return a
}

func denseFromMaps(data map[string][][]float64) map[string]*mat.Dense {
	rets := make(map[string]*mat.Dense)
	for k, m := range data {
		rets[k] = mat.NewDense(len(m), len(m[0]), nil)
		for i, r := range m {
			rets[k].SetRow(i, r)
		}
	}

	return rets
}

func InitializeParameters(filename string) (map[string]*mat.Dense, map[string]*mat.Dense, error) {
	// load from json, used in Local serving
	f, err := os.Open("./assets/" + filename)
	if err != nil {
		// load from json, used in Production
		f, err = os.Open(filename)
		if err != nil {
			// load from json, used in Testing
			f, err = os.Open("../../assets/" + filename)
			if err != nil {
				return nil, nil, fmt.Errorf("open json parameters file error: %s", err)
			}
		}
	}

	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	paramsJSON := &ParamsJSON{}
	// no need to handle error, since we're
	// dealing with the internal data
	err = json.Unmarshal(bytes, paramsJSON)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal json to params object error: %s", err)
	}

	parameters := denseFromMaps(paramsJSON.Parameters)
	scales := denseFromMaps(paramsJSON.Scales)

	return parameters, scales, nil
}

func PredictLogisticInstance(X []float64, parameters map[string]*mat.Dense, scales map[string]*mat.Dense) (float64, error) {
	a := mat.NewDense(len(X), 1, X)
	a.Sub(a, scales["mean"])
	a.DivElem(a, scales["std"])
	L := len(parameters) / 2
	for l := 1; l <= L; l++ {
		wr, _ := parameters["W"+strconv.Itoa(l)].Dims()
		z := mat.NewDense(wr, 1, nil)
		z.Mul(parameters["W"+strconv.Itoa(l)], a)
		z.Add(z, parameters["b"+strconv.Itoa(l)])
		if l != L {
			// relu forward
			a = relu(z)
		} else {
			// sigmoid forward
			a = sigmoid(z)
		}
	}

	// checking indices limit to only 0, 0
	ar, ac := a.Dims()
	if ar == 0 || ac == 0 {
		return -1, fmt.Errorf("shape of propability is wrong, want (1, 1), got (%s, %s)", ac, ar)
	}
	// transform result to percentage
	return a.At(0, 0), nil
}
