package gameOfLife

import (
    "math/bits"
    //"fmt"
)

type Data struct {
    Value []int
    Key []int
}

type RuleSet struct {
    CameToLife []int
    KeepAlive []int
}

type Grid struct {
    neighbors []int
    cells []int
}

/*

features to add:
    choices to define if we apply the livinghood or the battleroyale on the cell/neighbors defined by the user

*/

func Life(data Data, set RuleSet) Data {
    valueGrid := genGrid(data.Value)
    keyGrid := genGrid(data.Key)

    /*
    for i := 0 ; i < (len(valueGrid.cells) - 1) / 2 ; i = i + 1 {
        fmt.Println(valueGrid.neighbors, valueGrid.cells)
        nextValue := battleRoyale(valueGrid.neighbors[:], valueGrid.cells[:], set)
        valueGrid = genGrid(nextValue)
    }
    for i := 0 ; i < (len(keyGrid.cells) - 1) / 2 ; i = i + 1 {
        nextKey := battleRoyale(keyGrid.neighbors[:], keyGrid.cells[:], set)
        keyGrid = genGrid(nextKey)
    }*/
    //fmt.Println(valueGrid, keyGrid)
    
    nextValue := battleRoyale(valueGrid.neighbors[:], valueGrid.cells[:], set)
    //valueGrid = genGrid(nextValue)
    nextKey := battleRoyale(keyGrid.neighbors[:], keyGrid.cells[:], set)
    //keyGrid = genGrid(nextKey)

    //fmt.Println(nextValue, nextKey)
    /*
    nextValue = livingHood(valueGrid.neighbors, valueGrid.cells, set)
    nextKey = livingHood(keyGrid.neighbors, keyGrid.cells, set)

    fmt.Println(nextValue, nextKey)
    */
    return Data{nextValue, nextKey}
}

func genGrid(data []int) Grid {
    grid := Grid{}
    for _, v := range data {
        grid.neighbors = append(grid.neighbors, bits.OnesCount8(uint8(v)))
        grid.cells = append(grid.cells, v % 2)
    }
    return grid
}

/*
two methods can be made:

the battleroyale:
    when a cell is kipped alive OR if a new cell is born
        cell is continuing it's life
    else
        all neighbors are dead

the livinghood:
    when a cell is kipped alive:
        one neighbors dies
    when a cell came to life:
        add one neighbors to the pool
    else
        all neighbors are dead
*/

func battleRoyale(neighbors []int, cells []int, set RuleSet) []int {
    br := make([]int, 0)
    for i := range cells {
        if newBorn(neighbors[i], cells[i], set.CameToLife) {
            br = append(br, neighbors[i])
        } else if lifeGoesOn(neighbors[i], cells[i], set.KeepAlive) {
            br = append(br, neighbors[i])
        } else {
            br = append(br, 0)
        }
    }
    return br
}

func livingHood(neighbors []int, cells []int, set RuleSet) []int {
    lh := make([]int, 0)
    for i := range cells {
        if newBorn(neighbors[i], cells[i], set.CameToLife) && !(lifeGoesOn(neighbors[i], cells[i], set.KeepAlive)) {
            lh = append(lh, neighbors[i] + 1)
        } else if lifeGoesOn(neighbors[i], cells[i], set.KeepAlive) && !(newBorn(neighbors[i], cells[i], set.CameToLife)) {
            lh = append(lh, neighbors[i] - 1)
        } else {
            lh = append(lh, 0)
        }
    }
    return lh
}

func newBorn(n int, c int, r []int) bool {
    for _, v := range r {
        if n == v && c == 0 {
            return true
        }
    }
    return false
}

func lifeGoesOn(n int, c int, r []int) bool {
    for _, v := range r {
        if n == v && c == 1 {
            return true
        }
    }
    return false
}
