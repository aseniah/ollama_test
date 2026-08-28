using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

var lines = File.ReadAllLines("input/text.txt");
var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    var lowerLine = line.ToLower();
    var words = Regex.Matches(lowerLine, @"[a-z]+")
        .Cast<System.Text.RegularExpressions.Match>()
        .Select(m => m.Value);
    
    foreach (var word in words)
    {
        if (wordCounts.ContainsKey(word))
            wordCounts[word]++;
        else
            wordCounts[word] = 1;
    }
}

var sorted = wordCounts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var word in sorted)
    Console.WriteLine($"{word.Key}: {word.Value}");