using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string text = File.ReadAllText("input/text.txt");
text = text.ToLower();
text = new Regex(@"[^a-z\s]").Replace(text, "");
string[] words = text.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);

var freq = new Dictionary<string, int>();
foreach (var word in words)
{
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