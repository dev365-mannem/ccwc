<h1>Build Your Own wc Tool</h1>

```
>go build && go install

>ccwc -c data/test.txt
  342190 data/test.txt

>ccwc -l data/test.txt
    7145 test.txt

>ccwc -w data/test.txt
   58164 test.txt

>ccwc -m test.txt
  339292 test.txt

>ccwc test.txt
    7145   58164  342190 test.txt

>cat data/test.txt | wc -m
  339292
```
