using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string content = File.ReadAllText("input/text.txt");
string[] tokens = content.Split(new[] { ' ', '\t', '\r', '\n', ',' }, StringSplitOptions.RemoveEmptyEntries);

var freq = new Dictionary<string, int>();
foreach (var token in tokens)
{
    string word = string.Concat(token.Where(c => char.IsLetter(c))).ToLowerInvariant();
    if (string.IsNullOrEmpty(word)) continue;
    if (freq.ContainsKey(word))
        freq[word]++;
    else
        freq[word] = 1;
}

var sorted = freq.OrderByDescending(kv => kv.Value).ThenBy(kv => kv.Key);

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}