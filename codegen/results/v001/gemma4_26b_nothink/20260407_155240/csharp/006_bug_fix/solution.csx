using System;
using System.Collections.Generic;

List<string> FizzBuzz(int n) {
    var result = new List<string>();
    // Fixed: The loop should start at 1 and include n to correctly represent the sequence up to n.
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