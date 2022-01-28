#!/bin/bash
python3 ./tools/push_sentences_to_mysql.py \
      --host=127.0.0.1 \
      --port=33069 \
      --database=mag_server \
      --username=root  \
      --passwd=WQAOIaiona8X \
      --path=./sentences.tsv