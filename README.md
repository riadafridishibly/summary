# Print directory summary [WIP]

Print the directory summary. Like,
- How many files
  - How many video files
  - How many jgegs
  - etc.
- Size of the files

Controllable with flags


Optimistic implementation!

I'm mostly interested in a file that has an extension. In the future, I'll use mime-type to detect file type, for now, it only handles extension and total size.

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
        .sample:     13 [     23 kB]
            .go:      5 [    2.8 kB]
          .pack:      1 [    1.6 kB]
           .idx:      1 [    1.2 kB]
            .md:      1 [     915 B]
     .gitignore:      1 [     278 B]
           .sum:      1 [     177 B]
```
