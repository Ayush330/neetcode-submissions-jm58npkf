import "slices"

func carFleet(target int, position []int, speed []int) int {
    type pair struct{
        posn  int
        speed int
    }
    array := make([]pair, len(speed))
    for i:=0; i<len(speed); i++{
        array[i] = pair{position[i], speed[i]}
    }
    slices.SortFunc(array, func(a, b pair)int{
        return a.posn - b.posn
    })

    stack := []pair{}
    stack = append(stack, array[len(array)-1])

    for i:=len(array)-2; i>=0; i--{
        top := stack[len(stack)-1]
        speedTop := float64(target-top.posn)/ float64(top.speed)
        speedCurr := float64(target-array[i].posn)/ float64(array[i].speed)
        fmt.Println("Speed Top: ", speedTop, " Speed Current: ", speedCurr)
        if speedCurr > speedTop{
            stack = append(stack, array[i])
        }
    }
    return len(stack)
}
