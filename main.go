package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var TPLL float64 = 0
	var T float64 = 0
	var TA float64 = 0
	var IA float64 = 0
	var DQ float64 = 0
	var TFQ float64 = 0
	var TAQ float64 = 1
	var P float64 = 0
	var SPQ float64 = 0
	var PTO []float64
	var PTE float64 = 0
	var STLL float64 = 0
	var STS float64 = 0
	var STA float64 = 0
	var COP float64 = 0
	var TPS []float64
	var STO []float64
	var ITO []float64
	var NC int = 6
	var NQ int = 36
	var NT int = 0
	var CLL int = 0
	var NV []int
	var TF float64 = 600000000
	var HV float64 = 999999999
	var franja int = 2

	for i := 0; i < NC; i++ {
		TPS = append(TPS, HV)
		NV = append(NV, 0)
		STO = append(STO, 0)
		PTO = append(PTO, 0)
		ITO = append(ITO, 0)
	}

	fin := true
	for fin {
		x := getMenorTPS(TPS)
		if TPLL <= TPS[x] {
			T = TPLL
			CLL++
			NT++
			STLL += T
			IA = getIaByHour(franja)
			TPLL = T + IA
			i := getMenorFila(NV)
			NV[i]++

			if NV[i] == 1 {
				if T >= TFQ {
					TA = generoTA()
					TPS[i] = T + TA
					STO[i] = STO[i] + (T - ITO[i])

				} else {
					TPS[i] = T + TAQ
				}
			} else {
				if NT > NQ && T > TFQ {
					DQ = generoDQ()
					TFQ = T + DQ

					for i := 0; i < NC; i++ {
						TPS[i] = T + TAQ
						STA += TAQ
					}
				}
			}
		} else {
			T = TPS[x]
			NV[x]--

			if NV[x] >= 1 {
				if T < TFQ {
					P = getPrecio()
					SPQ += P
					TPS[x] = T + TAQ
					STA += TAQ
				} else {
					TA = generoTA()
					TPS[x] = T + TA
					STA += TA
				}
			} else {
				ITO[x] = T
				TPS[x] = HV
			}

			NT -= 1
			STS += T
		}

		if T >= TF {
			TPLL = HV
			if NT == 0 {
				generarResultados(STO, &PTO, T, &COP, NC, &PTE, STS, STLL, STA, CLL)
				imprimirResultados(PTO, NC, NQ, PTE, SPQ, COP, T, franja)
				fin = false
			}
		}
	}
}

func imprimirResultados(PTO []float64, NC, NQ int, PTE, SPQ, COP, T float64, franja int) {
	for i := 0; i < len(PTO); i++ {
		fmt.Printf("Colas: %d, Quiebre: %d, PTO: %f\n", NC, NQ, PTO[i])
	}

	PMQ := getPerdidaMensual(SPQ, franja, T)
	fmt.Printf("Colas: %d, Quiebre: %d, PTE: %f, PMQ: %f, COP: %f\n", NC, NQ, PTE, PMQ, COP)
}

func getPerdidaMensual(SPQ float64, franja int, T float64) float64 {
	if franja == 2 {
		return SPQ / (T / 1080000)
	}

	return SPQ / (T / 540000)
}

func generarResultados(STO []float64, PTO *[]float64, T float64, COP *float64, NC int, PTE *float64, STS, STLL, STA float64, CLL int) {
	for i := 0; i < len(STO); i++ {
		(*PTO)[i] = (STO[i] / T) * 100
	}
	*COP = float64(800000 * NC)
	*PTE = float64((STS - STLL - STA) / float64(CLL))
}

func getPrecio() float64 {
	r := rand.Float64()
	if r <= 0.0017 {
		return 3071
	} else if r <= 0.0495 {
		return 1500
	} else if r <= 0.1453 {
		return 800
	} else if r <= 0.41 {
		return 2335
	}

	return 500
}

func generoDQ() float64 {
	r := rand.Float64()
	return 180 / (math.Pow((1 - r), 0.4764))
}

func generoTA() float64 {
	r := rand.Float64()
	return 14 / (math.Pow((1 - r), 0.9225))
}

func getMenorFila(NV []int) int {
	j := 0
	for i := 0; i < len(NV); i++ {
		if NV[j] > NV[i] {
			j = i
		}
	}

	return j
}

func getIaByHour(franja int) float64 {
	if franja == 1 {
		return iaFranja1()
	} else if franja == 2 {
		return iaFranja2()
	}

	return iaFranja3()
}

func iaFranja3() float64 {
	r := rand.Float64()
	return -15.3846 * math.Log10(1-r)
}

func iaFranja2() float64 {
	r := rand.Float64()
	return -6.65336 * math.Log10(1-r)
}

func iaFranja1() float64 {
	r := rand.Float64()
	return -18.9358 * math.Log10(1-r)
}

func getMenorTPS(TPS []float64) int {
	j := 0
	for i := 0; i < len(TPS); i++ {
		if TPS[j] > TPS[i] {
			j = i
		}
	}

	return j
}
