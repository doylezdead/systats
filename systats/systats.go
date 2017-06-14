package systats

import (
    "os/exec"
    "bytes"
    "strings"
    "strconv"
)

func GetDiskFree() map[string]float64 {
    diskMap := make(map[string]float64, 0.0)
    cmd := exec.Command("df")
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Run()
    result := strings.Split(out.String(), "\n")

    for i:=1; i<len(result); i++ {
        fields := strings.Fields(result[i])
        if len(fields) < 6 {
            break
        }
        numer,_ := strconv.ParseFloat(fields[2], 64)
        denom,_ := strconv.ParseFloat(fields[3], 64)
        diskMap[fields[5]] = numer/denom
    }
    return diskMap
}

