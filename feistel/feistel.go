package feistel

import (
    "fmt"
    "os"
    "strconv"
    //"math"
    // custom librairies
    "cryptomata/gameOfLife"
)

type Parameters struct {
    linput []int
    rinput []int
    keys [][]int
}

/*
to experiment:
pas a rulle that is exponentially growing cells in the grid, overpopulating the grid
then we uses a rule that is destructive but not the absolute reverse of the first one
this will change a lots of things in the grid, we may add a configuration thing like seeding or something else to make the thing more random, thus less predictable
another thing that can be used to make the program more robust is by using different rules on the keys
making the seed from the key may be cooler i don't know where to start
*/

func Feistel(input []string) string {
    lr := input[0]
    l := sliceToIntSlice( lr[:len(lr)/2] )
    r := sliceToIntSlice( lr[len(lr)/2:] )

    ks := stringToIntSlice(input[1:])

    params := Parameters{ l, r, ks }

    for _, k := range params.keys {
        //fmt.Println(params.linput, params.rinput, k)
        // the program won't work well here, actually the nextLn do not work atm
        params = nextLn(params.linput, params.rinput, k)
    }
    params.linput, params.rinput = params.rinput, params.linput

    /* 
    actually, we may loses the last integer of one of the slices (linput & rinput) we have
    it's intended, i wanted lossy data as much as i can to modify most of the string before returning it to the user
    */
    //fmt.Println(params)
    t := make([]rune, 0)
    ret := ""
    for i := 0 ; i < ( (len(params.linput) - 1) / 2) ; i = i + 1 {
        t = append(t, rune(params.linput[i] + params.rinput[i]))
        ret = fmt.Sprintf("%s%x", ret, byte(t[i]))
    }

    return ret

}

func sliceToIntSlice(sl string) []int {
    ret := make([]int, 0)
    for _, s := range sl {
        ret = append(ret, int(rune(s)))
    }
    return ret
}

func stringToIntSlice(s []string) [][]int {
    ret := make([][]int, 0)
    for _, k := range s {
        t := make([]int, 0)
        for _, r := range k {
            t = append(t, int(r))
        }
        ret = append(ret, t)
    }
    return ret
}

func nextLn(l []int, r []int, k []int) Parameters {
    
    born := make([]int, 0)
    alive := make([]int, 0)

    rules := gameOfLife.RuleSet{
        CameToLife: append(born, 1, 2, 3, 4),
        KeepAlive: append(alive, 2, 3),
    }

    data := gameOfLife.Data{
        Value: r,
        Key: k,
    }
    /*
    for i := 0 ; i < (len(data.Value) + len(data.Key)) ; i = i + 1 {
        data = gameOfLife.Life(data, rules)
    }
    */
    data = gameOfLife.Life(data, rules)

    newValue := addWholeSlice(data.Value)
    newKey := addWholeSlice(data.Key)

    newLn := addWholeSlice(l)
    newRn := newValue + newKey

    ret := Parameters{
        linput: r,
        // it can't be everytime computed, sometimes, l has lower lenght than the rn value thus it's raising an out of bound
        rinput: intToIntSlice( ( intSliceToInt(r) ^ (newLn & newRn) )),
    }
    return ret
}

func addWholeSlice(sl []int) int {
    ret := 0
    for _, v := range sl {
        ret = ret + v
    }
    return ret
}

func intToIntSlice(v int) []int {
    t := strconv.Itoa(v)
    ret := make([]int, 0)
    for _, c := range t {
        ret = append(ret, int(rune(c)))
    }
    return ret
}

func intSliceToInt(sl []int) int {
    t := ""
    ret := 0
    for _, v := range sl {
        t = t + strconv.Itoa(v)
        temp, _ := strconv.Atoi(t)
        ret = ret + temp
    }
    return ret
}

func abort(s string) {
    fmt.Println(s)
    os.Exit(1)
}
