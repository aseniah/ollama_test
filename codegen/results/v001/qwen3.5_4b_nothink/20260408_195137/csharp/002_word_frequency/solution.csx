using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string[] Args;
string content = File.ReadAllText("input/text.txt");

string[] words = Regex.Matches(content, "[a-z]+")
    .Select(m => m.Value.ToLower())
    .ToArray();

var frequency = words.GroupBy(w => w)
    .OrderByDescending(g => g.Count())
    .ThenBy(g => g.Key)
    .Select(g => $"{g.Key}: {g.Count()}")
    .ToArray();

Console.Join(string.Join("\n", frequency));