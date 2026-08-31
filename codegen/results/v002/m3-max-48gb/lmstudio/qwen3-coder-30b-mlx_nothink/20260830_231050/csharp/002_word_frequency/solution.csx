using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var lines = File.ReadAllLines("input/text.txt");
var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    var words = line.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (var word in words)
    {
        // Convert to lowercase and remove punctuation
        var cleanWord = new string(word.Where(c => char.IsLetter(c)).ToArray()).ToLower();
        
        if (string.IsNullOrEmpty(cleanWord))
            continue;
            
        if (wordCounts.ContainsKey(cleanWord))
            wordCounts[cleanWord]++;
        else
            wordCounts[cleanWord] = 1;
    }
}

// Sort by count descending, then alphabetically ascending
var sortedWords = wordCounts.OrderByDescending(kvp => kvp.Value).ThenBy(kvp => kvp.Key);

foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}