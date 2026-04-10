using System.Text.Json;
using System.Text.Json.Nodes;

List<string> FixFizzBuzz(int n) {
    var result = new List<string>();
    for (int i = 1; i <= n; i++) {
        if (i % 5 == 0 && i % 3 == 0) result.Add("FizzBuzz");
        else if (i % 5 == 0) result.Add("Buzz");
        else if (i % 3 == 0) result.Add("Fizz");
        else result.Add(i.ToString());
    }
    return result;
}

foreach (var v in FixFizzBuzz(15)) {
    Console.WriteLine(v);
}