using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var text = File.ReadAllText("input/text.txt");
var words = text.Split(new char[] { ' ', '\t', '\n', '\r', ',', '.', '!', '?', ';', ':' }, StringSplitOptions.RemoveEmptyEntries)
                .Select(word => new string(word.Where(char.IsLetter).ToArray()).ToLowerInvariant())
                .Where(word => !string.IsNullOrEmpty(word));

var wordCounts = new Dictionary<string, int>();

foreach (var word in words)
{
    if (wordCounts.ContainsKey(word))
    {
        wordCounts[word]++;
    }
    else
    {
        wordCounts[word] = 1;
    }
}

var sortedWords = wordCounts.OrderByDescending(kvp => kvp.Value).ThenBy(kvp => kvp.Key);

foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}