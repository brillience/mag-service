#!/bin/bash
pip3 install pymysql,click
python3 ./tools/dn_init.py  \
      --host=127.0.0.1 \
      --port=33069 \
      --username=root  \
      --passwd=WQAOIaiona8X \
