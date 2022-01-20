## 0. Convert csv to tsv (optional)
```shell
cat articles.csv | ./tools/csv2tsv.sh >> articles.tsv
```
## 1. NLPMark
```shell
git clone https://github.com/brillience/NLPMarkTool
mv articles.tsv ./NLPMarkTool/ && cd ./NLPMarkTool
mvn clean
mvn install
mvn exec:java -Dexec.mainClass="com.zhang.nlp.Main"
```
## 2. Update NlpTags to mysql
```shell
pip3 install pymysql,pyyaml,click
python3 ./tools/push_sentences_to_mysql.py --path=./sentences.tsv
```