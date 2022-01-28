# pip3 install pymysql,click
import pymysql
import click



@click.command()
@click.option("--host",help="mysql host")
@click.option("--port",help="mysql port")
@click.option("--username",help="db user")
@click.option("--passwd",help="db password")
@click.option("--path",help="sentences.tsv path")
def init_db(host,port,database,username,passwd,path):
    port = int(port)
    db = pymysql.connect(host=host,port=port,user=username,password=passwd,charset='UTF8')
    print("[INFO] strat create database mag_server...")
    cursor = db.cursor()
    cursor.execute("CREATE DATABASE mag_server")
    print("[INFO] Done!")
    cursor.execute("DROP TABLE IF EXISTS nlpTags")
    print("[INFO] Start create table nlptags...")
    cursor.execute("""
    CREATE TABLE `nlpTags` (
    `doc_id` varchar(255) NOT NULL DEFAULT '',
    `nlp_tags` text NOT NULL DEFAULT '' COMMENT 'json字符串',
    PRIMARY KEY (`doc_id`),
    INDEX `doc_id`(`doc_id`(255))
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8
        """)
    print("[INFO] Done!")
    print("[INFO] Start create table user...")
    cursor.execute("""
    CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL DEFAULT '' COMMENT '账户',
    `nick` varchar(255)  NOT NULL DEFAULT '' COMMENT '昵称',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `number_unique` (`username`)
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4
        """)
    print("[INFO] Done!")
    cursor.close()
    db.commit()
    print("[INFO] Strat add admin user...")
    insertSql = "insert into `user` (`username`,`password`) values (%s,%s)"
    cursor.execute(insertSql,("admin","123456"))
    db.commit()
    print("[INFO] Done! API user:admine; passwd:123456;")
    db.close()

if __name__ == '__main__':
    init_db()
