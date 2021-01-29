# Print directory summary [WIP]

Print the directory summary. Like,
- How many files
  - How many video files
  - How many jgegs
  - etc.
- Size of the files

Controllable with flags


Optimistic implementation!

I'm mostly interested in file that has extension. In future I'll use mimetype to detect file type, for now it only handles extension and total size.

## How to run

```
$ git clone https://github.com/riadafridishibly/summary.git
$ cd summary
$ go build 
$ go install

$ # Now run summary from anywhere
$ summary

Unique file types: 8

      Extension  Count       Size
      ---------  -----       ----
     .gitignore:      1 [     278 B]
            .go:      2 [    2.0 kB]
           .idx:      1 [    1.2 kB]
            .md:      1 [     518 B]
           .mod:      1 [      98 B]
          .pack:      1 [    1.6 kB]
        .sample:     13 [     23 kB]
           .sum:      1 [     177 B]
```