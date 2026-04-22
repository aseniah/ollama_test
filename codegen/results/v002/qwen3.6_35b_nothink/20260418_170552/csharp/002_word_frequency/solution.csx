using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string[] lines = File.ReadAllLines("input/text.txt");
string[] words = lines
    .SelectMany(line => line.Split(' ', StringSplitOptions.RemoveEmptyEntries))
    .Select(word => word.ToLower())
    .Select(word => new string(word.Where(char.IsLetter).ToArray()))
    .Where(word => word.Length > 0)
    .ToArray();

var freq = words
    .GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(g => g.Count)
    .ThenBy(g => g.Word)
    .ToArray();

foreach (var f in freq)
{
    Console.WriteLine($"{f.Word}: {f.Count}");
}