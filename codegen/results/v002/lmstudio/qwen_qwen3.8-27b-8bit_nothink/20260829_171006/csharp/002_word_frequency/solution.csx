using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string[] words = File.ReadAllText("input/text.txt")
    .Split(new[] { '\n', '\r', '\t', ' ' }, StringSplitOptions.RemoveEmptyEntries);

// Extract only letters from each word token
var wordTokens = words
    .Select(w => new string(w.Where(c => char.IsLetter(c)).ToArray()))
    .Where(w => w.Length > 0)
    .Select(w => w.ToLower());

var wordCounts = wordTokens
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

var sortedWords = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key);

foreach (var kv in sortedWords)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}