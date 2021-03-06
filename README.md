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
20+20*20 = 420.000000                               
1*2+3*4 = 14.000000                                 
10^10 = 10000000000.000000                          
1/0 = +Inf                                          
100 = 100.000000                                    
5+(6-3) = 8.000000                                  
(1-1)+(1-1) = 0.000000                              
(1+2)-3 = 0.000000        
1+(2+3+(4+5)+((6+7)+8)+9)+10 = 55.000000
(2+2)*2 = 8.000000
```

Further details (additional prints from main.go inputs):
```
equation: 1+1-1
calculation tree: ((1.000000 + 1.000000) - 1.000000)                   
calculated: (1.000000 + 1.000000) = 2.000000                           
calculated: ((1.000000 + 1.000000) - 1.000000) = 1.000000              
therefore calculated: 1+1-1 = 1.000000                                 
                                                                       
equation: 1+6*7-8                                                      
calculation tree: ((1.000000 + (6.000000 * 7.000000)) - 8.000000)      
calculated: (6.000000 * 7.000000) = 42.000000                          
calculated: (1.000000 + (6.000000 * 7.000000)) = 43.000000             
calculated: ((1.000000 + (6.000000 * 7.000000)) - 8.000000) = 35.000000
therefore calculated: 1+6*7-8 = 35.000000                              
                                                                       
equation: 1*2+3                                                        
calculation tree: ((1.000000 * 2.000000) + 3.000000)                   
calculated: (1.000000 * 2.000000) = 2.000000                           
calculated: ((1.000000 * 2.000000) + 3.000000) = 5.000000              
therefore calculated: 1*2+3 = 5.000000                                 
                                                                       
equation: 2+2*2                                                        
calculation tree: (2.000000 + (2.000000 * 2.000000))                   
calculated: (2.000000 * 2.000000) = 4.000000                           
calculated: (2.000000 + (2.000000 * 2.000000)) = 6.000000              
therefore calculated: 2+2*2 = 6.000000                                 

equation: 1+6*5/7
calculation tree: (1.000000 + ((6.000000 * 5.000000) / 7.000000))
calculated: (6.000000 * 5.000000) = 30.000000
calculated: ((6.000000 * 5.000000) / 7.000000) = 4.285714
calculated: (1.000000 + ((6.000000 * 5.000000) / 7.000000)) = 5.285714
therefore calculated: 1+6*5/7 = 5.285714

equation: 20+20*20
calculation tree: (20.000000 + (20.000000 * 20.000000))
calculated: (20.000000 * 20.000000) = 400.000000
calculated: (20.000000 + (20.000000 * 20.000000)) = 420.000000
therefore calculated: 20+20*20 = 420.000000

equation: 1*2+3*4
calculation tree: ((1.000000 * 2.000000) + (3.000000 * 4.000000))
calculated: (1.000000 * 2.000000) = 2.000000
calculated: (3.000000 * 4.000000) = 12.000000
calculated: ((1.000000 * 2.000000) + (3.000000 * 4.000000)) = 14.000000
therefore calculated: 1*2+3*4 = 14.000000

equation: 10^10
calculation tree: (10.000000 ^ 10.000000)
calculated: (10.000000 ^ 10.000000) = 10000000000.000000
therefore calculated: 10^10 = 10000000000.000000

equation: 1/0
calculation tree: (1.000000 / 0.000000)
calculated: (1.000000 / 0.000000) = +Inf
therefore calculated: 1/0 = +Inf

equation: 100
calculation tree: (100.000000)
calculated: (100.000000) = 100.000000
therefore calculated: 100 = 100.000000

equation: 5+(6-3)
calculation tree: (5.000000 + (6.000000 - 3.000000))
calculated: (6.000000 - 3.000000) = 3.000000
calculated: (5.000000 + (6.000000 - 3.000000)) = 8.000000
therefore calculated: 5+(6-3) = 8.000000

equation: (1-1)+(1-1)
calculation tree: ((1.000000 - 1.000000) + (1.000000 - 1.000000))
calculated: (1.000000 - 1.000000) = 0.000000
calculated: (1.000000 - 1.000000) = 0.000000
calculated: ((1.000000 - 1.000000) + (1.000000 - 1.000000)) = 0.000000
therefore calculated: (1-1)+(1-1) = 0.000000

equation: (1+2)-3
calculation tree: ((1.000000 + 2.000000) - 3.000000)
calculated: (1.000000 + 2.000000) = 3.000000
calculated: ((1.000000 + 2.000000) - 3.000000) = 0.000000
therefore calculated: (1+2)-3 = 0.000000

equation: 1+(2+3+(4+5)+((6+7)+8)+9)+10
calculation tree: ((1.000000 + ((((2.000000 + 3.000000) + (4.000000 + 5.000000)) + ((6.000000 + 7.000000) + 8.000000)) + 9.000000)) + 10.000000)
calculated: (2.000000 + 3.000000) = 5.000000
calculated: (4.000000 + 5.000000) = 9.000000
calculated: ((2.000000 + 3.000000) + (4.000000 + 5.000000)) = 14.000000
calculated: (6.000000 + 7.000000) = 13.000000
calculated: ((6.000000 + 7.000000) + 8.000000) = 21.000000
calculated: (((2.000000 + 3.000000) + (4.000000 + 5.000000)) + ((6.000000 + 7.000000) + 8.000000)) = 35.000000
calculated: ((((2.000000 + 3.000000) + (4.000000 + 5.000000)) + ((6.000000 + 7.000000) + 8.000000)) + 9.000000) = 44.000000
calculated: (1.000000 + ((((2.000000 + 3.000000) + (4.000000 + 5.000000)) + ((6.000000 + 7.000000) + 8.000000)) + 9.000000)) = 45.000000
calculated: ((1.000000 + ((((2.000000 + 3.000000) + (4.000000 + 5.000000)) + ((6.000000 + 7.000000) + 8.000000)) + 9.000000)) + 10.000000) = 55.000000
therefore calculated: 1+(2+3+(4+5)+((6+7)+8)+9)+10 = 55.000000

equation: (2+2)*2
calculation tree: ((2.000000 + 2.000000) * 2.000000)
calculated: (2.000000 + 2.000000) = 4.000000
calculated: ((2.000000 + 2.000000) * 2.000000) = 8.000000
therefore calculated: (2+2)*2 = 8.000000
```