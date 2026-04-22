using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var lines = File.ReadAllLines("input/text.txt");
var counts = new Dictionary<string, int>();

foreach (var line in lines)
{
    foreach (var word in line.Split(' ', StringSplitOptions.RemoveEmptyEntries))
    {
        string clean = new string(word.ToLower().Where(char.IsLetter).ToArray());
        if (clean.Length > 0)
        {
            counts.TryGetValue(clean, out int c);
            counts[clean] = c + 1;
        }
    }
}

var sorted = counts.OrderByDescending(x => x.Value).ThenBy(x => x.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}