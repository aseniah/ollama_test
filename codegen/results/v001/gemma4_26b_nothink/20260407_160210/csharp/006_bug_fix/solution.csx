using System;
using System.Collections.Generic;

List<string> FizzBuzz(int n) {
    var result = new List<string>();
    // The bug was that the loop started at 0 and ended at n-1.
    // FizzBuzz should iterate from 1 to n inclusive.
    for (int i = 1; i <= n; i++) {
        if (i % 15 == 0) result.Add("FizzBuzz");
        else if (i % 3 == 0) result.Add("Fizz");
        else if (i % 5 == 0) result.Add("Buzz");
        else result.Add(i.ToString());
    }
    return result;
}

foreach (var v in FizzBuzz(15)) 
{
    Console.WriteLine(v);
}