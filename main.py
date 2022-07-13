import time
from datetime import datetime
  
from influxdb_client import InfluxDBClient, Point, WritePrecision
from influxdb_client.client.write_api import SYNCHRONOUS

# You can generate an API token from the "API Tokens Tab" in the UI
token = "AoYJomd8FtlAfG1jUE_h0TRxabhRa15HlZcgzEoHFK-CPgtszef9fgRjxlMHWbXXPDDBCoQSdA0IhQ6qoLW-OQ=="
org = "dek_d"
bucket = "writer_transaction"

def read_data():
    with open('data/novel_transaction_2016 (2).csv') as f:
        return [x.split(',') for x in f.readlines()[1:]]

with InfluxDBClient(url="http://localhost:8086", token=token, org=org) as client:

    a = read_data()
    # sequence = []
    for metric in a:
        # # print(metric[8])
        time = metric[8].strip('"')
        novel_id = metric[2].strip('"')
        buyer_id = metric[1].strip('"')
        coin = metric[6].strip('"')
        x = "2022-01-02 13:33:00"
        dt = datetime().for datetime().strptime(x, '%Y-%m-%d %H:%M:%S')
        print(dt)
        # value = f"transaction,host=host1 time={time} novel_id={int(novel_id)} buyer_id={int(buyer_id)} coin={int(coin)}"
        # sequence.append(value)
        point = Point("transaction_3").field("novel_id", int(novel_id)).field("buyer_id", int(buyer_id)).field("coin", int(coin)).time(time)
        # write_api = client.write_api(write_options=SYNCHRONOUS)
        # write_api.write(bucket, org, point)
    # print(sequence)
    # write_api.write(bucket, org, sequence)

