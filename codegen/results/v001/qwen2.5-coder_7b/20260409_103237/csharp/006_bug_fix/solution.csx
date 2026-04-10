using System.Collections.Generic;

List<string> FizzBuzz(int n) {
    var result = new List<string>();
    for (int i = 1; i <= n; i++) { // Fixed: changed i = 0 to i = 1 and < n to <= n
        if (i % 15 == 0) result.Add("FizzBuzz");
        else if (i % 3 == 0) result.Add("Fizz");
        else if (i % 5 == 0) result.Add("Buzz");
        else result.Add(i.ToString());
    }
    return result;
}

foreach (var v in FizzBuzz(15)) Console.WriteLine(v);