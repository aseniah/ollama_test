#r "System.IO"
#r "System.Text.Json"
#r "System.Linq"
#r "System.Collections.Generic"

using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var file = File.ReadAllText("input/text.txt");
var content = new String(File.ReadAllBytes("input/text.txt"));
var result = "word: count";

if (file.Length > 0) {
    var lowerCaseFile = string.Empty;
    foreach (byte b in file) {
        lowerCaseFile += char.ToLower((char)b);
    }
    
    var arrayFile = new string(lowerCaseFile.ToCharArray());
    
    var wordsFile = arrayFile.Where(c => !char.IsLetterOrDigit(c)).ToArray().ToArray();
    
    var wordCount = new Dictionary<string, int>();
    foreach (var w in file) {
        if (!string.IsNullOrEmpty(w)) {
            lowerCaseFile += w;
            var arrayFile = new String(lowerCaseFile.ToCharArray());
            var wordsArray = string.Empty;
            
            foreach (byte b in arrayFile) {
                if (!char.IsLetterOrDigit((char)b)) {
                    if (!string.IsNullOrEmpty(wordsArray) || !char.IsLetter((char)b)) {
                        wordCount++;
                    }
                }
                
                if (char.IsLetter(b) || char.IsLetterAndDigit(b)) {
                    wordsFile = wordsFile.Concat(new string(w, 1));
                    break;
                }
            }
        }
    }
}

var wordsFile = file.Where(c => char.IsLetterOrDigit(c)).ToArray().ToArray();
string resultWord = "word: count";
foreach (char c in wordsFile) {
    if (char.IsLetterOrDigit(c)) {
        wordCount++;
    } else {
        wordCount--;
    }
}

var sortedWords = wordCount
    .OrderByDescending(w => w.Value)
    .ThenBy(w => w.Key)
    .ToList();

foreach (var kvp in sortedWords) {
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}