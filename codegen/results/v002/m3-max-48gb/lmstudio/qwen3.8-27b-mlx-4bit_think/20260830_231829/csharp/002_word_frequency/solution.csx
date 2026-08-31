using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string text = File.ReadAllText("input/text.txt");

var words = text
    .Split(new[] { ' ', '\t', '\n', '\r', ',' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => new string(w.Where(c => char.IsLetter(c)).ToArray()).ToLower())
    .Where(w => w.Length > 0);

var counts = words
    .GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word);

foreach (var item in counts)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}