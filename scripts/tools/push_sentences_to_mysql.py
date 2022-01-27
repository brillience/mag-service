# pip3 install pymysql,click
import pymysql
import click



def GetDocId(line:str):
    return line.split("\t")[0]

@click.command()
@click.option("--host",help="mysql host")
@click.option("--port",help="mysql port")
@click.option("--database",help="mysql db name")
@click.option("--username",help="db user")
@click.option("--passwd",help="db password")
@click.option("--path",help="sentences.tsv path")
def UpdateToDb(host,port,database,username,passwd,path):
    port = int(port)
    db = pymysql.connect(host=host,port=port,db=database,user=username,password=passwd)
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
    db.close()

def updateToMysql(db,docId,nlpTags):
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
