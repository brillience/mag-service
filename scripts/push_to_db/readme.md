## 0. Convert csv to tsv (optional)
```shell
cat *.csv | ./tools/csv2tsv.sh >> *.tsv
```
## 1. NLPMark
```shell
    git clone https://github.com/brillience/NLPMarkTool
    mv *.tsv ./NLPMarkTool/ && cd ./NLPMarkTool
    mv *.tsv articles.tsv
    mvn clean
    mvn install
    mvn exec:java -Dexec.mainClass="com.zhang.nlp.Main"
```
## 2. Update NlpTags to mysql