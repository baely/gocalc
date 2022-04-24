# Go Calc

Go calculator to parse and calculate math equations.

Works by parsing the equation into a calculation tree and then working through
the tree to calculate

Output from main.go:
```
1+1-1 = 1.000000
1+6*7-8 = 35.000000
1*2+3 = 5.000000   
2+2*2 = 6.000000   
1+6*5/7 = 5.285714 
```

Further details (additional prints from main.go inputs):
```
equation: 1+1-1
calculation tree: ((1.000000 + 1.000000) - 1.000000)             
calculated: 1.000000 + 1.000000 = 2.000000                       
calculated: 2.000000 - 1.000000 = 1.000000                       
therefore calculated: 1+1-1 = 1.000000                           
                                                                 
equation: 1+6*7-8                                                
calculation tree: ((1.000000 + (6.000000 * 7.000000)) - 8.000000)
calculated: 6.000000 * 7.000000 = 42.000000                      
calculated: 1.000000 + 42.000000 = 43.000000                     
calculated: 43.000000 - 8.000000 = 35.000000                     
therefore calculated: 1+6*7-8 = 35.000000                        
                                                                 
equation: 1*2+3                                                  
calculation tree: ((1.000000 * 2.000000) + 3.000000)             
calculated: 1.000000 * 2.000000 = 2.000000                       
calculated: 2.000000 + 3.000000 = 5.000000                       
therefore calculated: 1*2+3 = 5.000000                           
                                                                 
equation: 2+2*2                                                  
calculation tree: (2.000000 + (2.000000 * 2.000000))             
calculated: 2.000000 * 2.000000 = 4.000000                       
calculated: 2.000000 + 4.000000 = 6.000000                       
therefore calculated: 2+2*2 = 6.000000                           

equation: 1+6*5/7
calculation tree: (1.000000 + ((6.000000 * 5.000000) / 7.000000))
calculated: 6.000000 * 5.000000 = 30.000000
calculated: 30.000000 / 7.000000 = 4.285714
calculated: 1.000000 + 4.285714 = 5.285714
therefore calculated: 1+6*5/7 = 5.285714
```