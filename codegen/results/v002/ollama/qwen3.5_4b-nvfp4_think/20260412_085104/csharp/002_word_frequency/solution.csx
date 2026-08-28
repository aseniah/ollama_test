using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/text.txt").ToList();
var allText = string.Join("", lines);

var words = allText
    .Where(char.IsLetter)
    .Select(char.ToLower)
    .ToArray()
    .GroupBy((_, _) => true, (w1, w2) => char.GetUnicodeCategory(w1)) // Placeholder to collect chars
    .SelectMany(w => new[] { w }).ToList();

var wordCounts = new Dictionary<string, int>();

foreach (char c in allText.Where(char.IsLetterOrDigit).Where(c => c != ' ') 
                           .Where(c => c != ',').Where(c => c != '.')) {
    // Reconstructing words is complex - let's use a simpler approach
    
    // This won't work properly - need actual word extraction logic
}

// Better implementation:
var wordChars = new List<char>();

foreach (char c in allText) {
    if (char.IsLetter(c)) {
        wordChars.Add(char.ToLower(c));
    }
}

// Actually, let's process each character properly
var wordsList = new List<string>();
var tempWord = new StringBuilder();

foreach (char c in allText) {
    if (char.IsLetterOrDigit(c)) {
        tempWord.Append(char.ToLower(c));
    } else {
        if (tempWord.Length > 0) {
            wordsList.Add(tempWord.ToString());
            tempWord.Clear();
        }
    }
}

// Final cleanup
var validWords = wordsList.Where(w => w.Length > 0).ToList();
var wordCounts = new Dictionary<string, int>();

foreach (var word in validWords) {
    // Extract each letter-only character from the processed text
    var chars = allText.Where(c => char.IsLetter(c));
    var extractedWord = string.Empty;
    
    foreach (char c in chars) {
        extractedWord += char.ToLower(c);
    }
    
    if (!string.IsNullOrEmpty(extractedWord)) {
        wordCounts[extractedWord]++;
    }
}

var sortedWords = wordCounts
    .OrderByDescending(w => w.Value)
    .ThenBy(w => w.Key)
    .ToList();

foreach (var pair in sortedWords) {
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}

// Actually, this is getting too complex. Let me simplify: