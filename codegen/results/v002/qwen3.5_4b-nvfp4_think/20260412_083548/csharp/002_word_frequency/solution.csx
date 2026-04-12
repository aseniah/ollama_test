using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var text = File.ReadAllLines("input/text.txt");
var words = new Dictionary<string, int>();

foreach (var line in text)
{
    var wordsInLine = line.Split(' ');
    foreach (var word in wordsInLine)
    {
        word = word.Trim();
        if (string.IsNullOrEmpty(word)) continue;
        
        string cleanWord = word.AsSpan().Where(char.IsLetter).ToArray()
            .ToString().ToLowerInvariant();
        if (string.IsNullOrEmpty(cleanWord)) continue;
        
        var existing = words[cleanWord];
        if (!existing.HasValue) words.Add(cleanWord, 0);
        else { var oldVal = words[cleanWord] ?? 0; words[cleanWord] = oldVal + 1; }
    }
}

foreach (var kvp in words.OrderBy(kv => kv.Value).ThenBy(kv => -kv.Key.Length))
{
    System.Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}