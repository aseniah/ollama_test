using System;
using System.IO;
using System.Text.RegularExpressions;
using System.Collections.Generic;
using System.Linq;

string text = File.ReadAllText("input/text.txt").ToLowerInvariant();
var regex = new Regex(@"[a-z]+");
var matches = regex.Matches(text);
Dictionary<string, int> counts = new Dictionary<string, int>();
foreach (Match m in matches)
{
    string word = m.Value;
    if (counts.ContainsKey(word))
    {
        counts[word] += 1;
    }
    else
    {
        counts[word] = 1;
    }
}

var result = counts
    .Select(x => new { Key = x.Key, Value = x.Value })
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var item in result)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}