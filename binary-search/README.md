Iterative Binary Search Steps
Get the sorted array
Compare the target value with the value at the middle index of the array
Three cases
If the target value is equal to the middle element, the search is successful, and returns the index of the middle element.
If the target value is less than the middle element, update the ending index of the search range to be the middle index minus one and repeat the process from step 2.
If the target value is greater than the middle element, update the starting index of the search range to be the middle index plus one and repeat the process from step 2.


Recursive Binary search Steps
Get the sorted array
Compare the left boundary exceeds the right boundary, it indicates that the target element is not present within the current search space
Calculate the middle index mid within the current search space. It ensures that the middle index is computed without causing integer overflow issues, even for large values of left and right.
Check the element at the middle index mid matches the target value, the function returns mid to indicate that the target element has been found.
If the element at mid is less than the target value, suggests that the target element must be in the right half of the current search space. Call the function recursively with an updated left boundary (mid+1) while keeping the right boundary unchanged.
If the element at mid is greater than the target value, it implies that the target element is in the left half of the current search space. Call the function recursively with an updated right boundary (mid-1) while keeping the left boundary unchanged.
The function uses a divide-and-conquer approach to recursively search for the target element in smaller and smaller sub arrays until the target is found or the search space is exhausted.