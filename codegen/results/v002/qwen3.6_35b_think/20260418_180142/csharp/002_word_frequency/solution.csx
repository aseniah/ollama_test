using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;

var lines = File.ReadAllLines("input/text.txt");
var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    var words = line.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    foreach (var word in words)
    {
        var cleanWord = string.Concat(word.ToLower().Where(char.IsLetter).ToArray());
        if (cleanWord.Length > 0)
        {
            wordCounts[cleanWord] = wordCounts.GetValueOrDefault(cleanWord, 0) + 1;
        }
    }
}

var sorted = wordCounts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key);

foreach (var item in sorted)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}