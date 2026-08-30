using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var content = File.ReadAllLines("input/text.txt");
var words = new Dictionary<string, int>();

foreach (var line in content)
{
    var lineLower = line.ToLowerInvariant();
    
    foreach (var c in lineLower)
    {
        if (char.IsLetter(c))
        {
            continue;
        }
    }

    var chars = lineLower.Select(c => char.IsLetter(c) ? c : ' ').ToArray();
    var text = new string(chars);
    
    if (text.Length == 0) continue;

    // Split by whitespace first, then strip non-letters from each token
    var tokens = text.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    foreach (var token in tokens)
    {
        if (string.IsNullOrEmpty(token)) continue;
        
        // Keep only letters
        string cleaned = new string(token.Where(char.IsLetter).ToArray());
        
        if (cleaned.Length == 0) continue;
        
        if (!words.ContainsKey(cleaned))
        {
            words[cleaned] = 0;
        }
        words[cleaned]++;
    }
}

var sortedWords = words
    .OrderByDescending(k => k.Value)
    .ThenBy(k => k.Key)
    .ToList();

foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}