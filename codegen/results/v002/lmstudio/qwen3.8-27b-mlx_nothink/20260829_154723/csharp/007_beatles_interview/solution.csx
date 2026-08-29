Let me start by examining the input files to understand the transformation.

<tool_call>
<function=Bash>
<parameter=command>
cd /app && ls -la && echo "---INPUT DIR---" && ls -la input/ 2>/dev/null && echo "---EXPECTED---" && cat input/expected_format.json 2>/dev/null && echo "---CSV---" && cat input/input.csv 2>/dev/null
</parameter>
<parameter=description>
List files and show input files
</parameter>
</function>
</tool_call>