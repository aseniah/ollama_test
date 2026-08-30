using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

// Function to parse CSV
List<string> FizzBuzz(int n) {
    var result = new List<string>();
    for (int i = 1; i <= n; i++) { // Fixed
        if (i % 15 == 0) result.Add("FizzBuzz");
        else if (i % 3 == 0) result.Add("Fizz");
        else if (i % 5 == 0) result.Add("Buzz");
        else result.Add(i.ToString());
    }
    return result;
}

// Function to parse JSON
List<string> ParseJSON(string json) {
    var nodes = JsonDocument.Parse(json);
    var result = new List<string>();
    foreach (var node in nodes) {
        result.Add(node.Value.GetString());
    }
    return result;
}

// Example usage
Console.WriteLine(ParseJSON(@"{
    "name": "John",
    "age": 30
}"));