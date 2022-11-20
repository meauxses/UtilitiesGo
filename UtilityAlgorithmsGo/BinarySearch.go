/*Procedure binary_search
   A ← sorted array
   n ← size of array
   x ← value to be searched

   Set lowerBound = 1
   Set upperBound = n 

   while x not found
      if upperBound < lowerBound 
         EXIT: x does not exists.
   
      set midPoint = lowerBound + ( upperBound - lowerBound ) / 2
      
      if A[midPoint] < x
         set lowerBound = midPoint + 1
         
      if A[midPoint] > x
         set upperBound = midPoint - 1 

      if A[midPoint] = x 
         EXIT: x found at location midPoint
   end while
   
end procedure*/
package BinarySearch
import "fmt"

func BinarySearch (sortedArray []T, value int) int{
    length := len(sortedArray)
    lower := 0
    upper := length - 1
    returnIndex := -1

    for lower != upper {
        mid := lower + upper / 2
        if sortedArray[mid] == value {
            returnIndex = mid
        } else if value > sortedArray[mid] {
            lower = mid + 1
        } else if value < sortedArray[mid] {
            upper = mid - 1
        }
    }

    return returnIndex
}