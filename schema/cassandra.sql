drop table if exists wallet;
create table if not exists wallet
(
    currency text,
    user_id  text,
    amount   decimal,
    updated_time timestamp,
    primary key (currency, user_id)
)
    with caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
     and compaction = {'max_threshold': '32', 'min_threshold': '4', 'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
     and compression = {'class': 'org.apache.cassandra.io.compress.LZ4Compressor', 'chunk_length_in_kb': '64'}
;