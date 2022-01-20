# pip3 install pymysql,pyyaml,click
import os
import yaml
import pymysql
import click

config_path= os.path.abspath(os.path.join(os.getcwd(),"../../rpc/etc/mag.yaml"))

with open(config_path) as f:
    config = yaml.load(stream=f, Loader=yaml.FullLoader)
host = config['Mysql']['DataSource'].split('(')[-1].split(')')[0].split(':')[0]
port = int(config['Mysql']['DataSource'].split('(')[-1].split(')')[0].split(':')[1])
database = config['Mysql']['DataSource'].split('?')[0].split('/')[-1]
username = config['Mysql']['DataSource'].split(':')[0]
passwd = config['Mysql']['DataSource'].split(":")[1].split('@')[0]

db = pymysql.connect(host=host,port=port,db=database,user=username,password=passwd)

def GetDocId(line:str):
    return line.split("\t")[0]

@click.command()
@click.option("--path",help="sentences.tsv path")
def UpdateToDb(path:str):
    preDocId = ""
    nlpTags = ""
    with open(path) as f:
        for line in f:
            if preDocId=="":
                preDocId=GetDocId(line)
            curDocId = GetDocId(line)
            if curDocId==preDocId:
                nlpTags=nlpTags+line
            else:
                updateToMysql(preDocId,nlpTags)
                preDocId=curDocId
                nlpTags=line
    updateToMysql(preDocId,nlpTags)

def updateToMysql(docId:str,nlpTags:str):
    # 先判断当前doc是否存在
    cursor = db.cursor()
    querySql = "select * from `nlpTags` where `doc_id`=%s"
    cursor.execute(querySql,(docId))
    res = cursor.fetchall()
    if len(res)==0:
        # 该文档不存在，则插入
        insertSql = "insert into `nlpTags` (`doc_id`,`nlp_tags`) values (%s,%s)"
        cursor.execute(insertSql,(docId,nlpTags))
        db.commit()
    else:
        updateSql = "update `nlpTags` set `nlp_tags`=%s where `doc_id`=%s"
        cursor.execute(updateSql,(nlpTags,docId))
        db.commit()
    print("[INFO]: Update to mysql {} Finished!".format(docId))
    return

            
            

if __name__ == '__main__':
    UpdateToDb()
    db.close()