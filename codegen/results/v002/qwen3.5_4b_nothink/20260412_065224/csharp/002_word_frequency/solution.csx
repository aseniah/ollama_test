#r "System.Collections.Generic"

#r "System.Linq"

#r "System.IO"

#r "System.Text.RegularExpressions"

#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using System.Text.Json;

string content = File.ReadAllText("input/text.txt");

string[] lines = content.Split('\n');

var wordFrequencies = new Dictionary<string, int>();

foreach (var line in lines)
{
    var lowerLine = line.ToLowerInvariant();
    
    MatchCollection matches;
    using (var regex = new Regex(@"\b[A-Za-z]+\b"))
    {
        matches = regex.Matches(lowerLine);
    }
    
    foreach (var match in matches)
    {
        var word = match.Value.Trim();
        if (word == "") continue;
        
        if (!wordFrequencies.ContainsKey(word))
        {
            wordFrequencies[word] = 0;
        }
        
        wordFrequencies[word]++;
    }
}

var sortedWords = wordFrequencies.OrderByDescending(w => w.Value, w => w.Key)
                                 .ToList();

foreach (var item in sortedWords)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}