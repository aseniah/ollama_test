using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var words = new Dictionary<string, int>();
var lines = File.ReadAllLines("input/text.txt");

foreach (var line in lines)
{
    var tokens = line.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (var token in tokens)
    {
        // Remove punctuation and convert to lowercase
        var cleanWord = new string(token.Where(c => char.IsLetter(c)).ToArray()).ToLower();
        
        if (!string.IsNullOrEmpty(cleanWord))
        {
            if (words.ContainsKey(cleanWord))
                words[cleanWord]++;
            else
                words[cleanWord] = 1;
        }
    }
}

// Sort by count descending, then alphabetically ascending
var sortedWords = words.OrderBy(pair => -pair.Value).ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}