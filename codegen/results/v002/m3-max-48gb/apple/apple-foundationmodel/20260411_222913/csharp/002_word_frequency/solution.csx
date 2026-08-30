using System;
using System.Collections.Generic;
using System.Text.Json;

string[] text = File.ReadAllLines("input/text.txt");
var words = string.Join(" ", text).ToLower().Split(new[] { ' ', '.', '!', '?', ';', ':', ',', '(', ')', '-', '_', '"', "'", '$', '%', '&', '*', '/', '=', '+', '#', '@', '[', ']', '{', '}', '|', '}', '<', '>', '^', '`', '{', '}', '~' }, StringSplitOptions.RemoveEmptyEntries);
var wordCounts = new Dictionary<string, int>();

foreach (var word in words)
{
    var wordNoPunct = word.Replace("[^a-zA-Z]", "");
    if (wordNoPunct.Length > 0)
    {
        wordCounts[wordNoPunct]++;
    }
}

var sortedWords = wordCounts.OrderByDescending(x => x.Value).ThenBy(x => x.Key);
foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}