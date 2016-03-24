# gomerge
A tool to do file merging.

### Download
[https://github.com/fredyw/gomerge/releases](https://github.com/fredyw/gomerge/releases)

### Usage
    ./gomerge --template test.template --output test.txt

### Example
**Template file**

    begin
        line 1
    {{ Merge "part1.txt" }}
        line 2
    {{ Merge "part2.txt" }}
        line 3
    end

`Merge` is a special function that reads the content of a file into the template file.

**part1.txt**

        From part1: hello world
        From part1: bye world

**part2.txt**

        From part2: foo
        From part2: bar

**Output file**

    This is the first paragraph.
    This is the second paragraph.

        From part1: hello world
        From part1: bye world

    This is another paragraph.

        From part2: foo
        From part2: bar

    This is the last paragraph.