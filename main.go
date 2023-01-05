//Andrea Labra Orozco CS311 - hw1 
package main

import "fmt"

//function to CalcPressure that takes 3 parameter v, n, t and returns the pressure the gas is under 

func CalcPressure(v,n,t float64) float64{
  //R := 8.3144598 
  //take 3 parameters, v,n,t 
  //to calculate pressure p=nRt/v 
  var r float64 = 8.3144598
  term1 := n * r * t //multiply the values 
  term2 := term1/v  //now take the result from above and divide by v 
  return term2 
}

//(mutating maps lesson, methods)
//take two parameters, e and c 
func Decode(e []byte, c map[byte]byte) []byte {
  //range will iterate over the values in the map. if certain keys are found in the map, if m and i are found in the map e will be returned (e)
  for i, m := range e { if d, ok := c[m]; ok { e[i] = d } };
  return e;
}

func OddParity(nums []int) (bool, bool) {
  //both cases are initially true 
  con1:=true 
  con2:=true 
  //the counter will keep count of the number of 1 in the array 
  counter:=0     
  //this loop will go through all the array and see if the number is 1 and if it is the counter will account that
  for i := 0; i < len(nums); i++ { 
    if nums[i]==1 { 
      counter++
//chnage the first case to false if theres other numbers that is not what we are looking for (0,1)
} else if nums[i]!=0 {  
      con2=false         
    }
  }
  if counter%2==0 {  // change the first conditoion to false if the number of 1 is an even count 
    con1=false
  }             
  return con1, con2
}

func main() {

  //Pressre should give out 2478.9561p 
  fmt.Println(CalcPressure(1.0,1.0,298.15),"pascals");
  //Second case of CalcPressure, after using the online pressure calculator, t=500, v=5
  //n=5 pressure should be approximately 4157.23 p
  fmt.Println(CalcPressure(5,5,500),"pascals"); 
  
  c := map[byte]byte{'e': 'u', 'h': 'f', 'l': 'n', 'o': 'y'};
  fmt.Println(string(Decode([]byte("hello"), c)));   //should print out funny 
  fmt.Println(string(Decode([]byte("heell"), c))); //should print out fuunn 
  
  fmt.Println(OddParity([]int{0,1,1,1})); // should pirnt out true true; since con 1 has an odd number of ones, con 2 true since all values are 0 and 1
  fmt.Println(OddParity([]int{1,1,3,5})); // should prit out false false; since con 1 has an even number of 1s and cond 2 since there are other values other than 0 and 1 
}