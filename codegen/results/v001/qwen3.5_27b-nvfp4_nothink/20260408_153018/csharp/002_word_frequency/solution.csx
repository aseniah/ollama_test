#r "System.IO"
#r "System.Collections"
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var text = File.ReadAllText("input/text.txt");
var words = Regex.Matches(text, @"[a-zA-Z]+")
    .Cast<Match>()
    .Select(m => m.Value.ToLower());

var counts = new Dictionary<string, int>();

foreach (var word in words)
{
    if (counts.ContainsKey(word))
    {
        counts[word]++;
    }
    else
    {
        counts[word] = 1;
    }
}

var sorted = counts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .Select(x => $"{x.Key}: {x.Value}")
    .ToArray();

foreach (var line in sorted)
{
    Console.WriteLine(line);
}