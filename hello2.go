package main
import "fmt"
import "os"
func main() {
//s := "abc"
//s += " 123"
//fmt.Println(s[:])

//sNum := "1"
//num, _ := strconv.Atoi(sNum)

//num := 1
//sNum := string(num)

//s1 := [3]int{7,8,9}
//reverse(s1)
//fmt.Printf(s1)
//n,err := strconv.Atoi(os.Args[1])
//if err != nil{
//  log.Fatalf("can;t read ang1 :"%s, err)

m := map[string]int{}

file,  := os.Open(fileName)

defer file.Close()

}
finc genRandomNums(n int) []int{
  nums := []int{}
  for i := 0; i< n; i++{
    nums = append(nums, rand.Int())
  }
  return nums
}

//func reverse(s []int){
//  for i,j := 0, len(s)-1; i<j; i,j=i+1,j-1{
//    s[i],s[j] = s[j],s[i]
//  }
//}
