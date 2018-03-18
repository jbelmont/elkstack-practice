# Elastic Search, Logstash and Kibana 

If you want to generate movies.json then do the following command:

`go run main.go > movies.json`

If you want to generate tags.json then do the following command:

`cd tags; go run main.go > tag.json`

Install elasticsearch via homebrew:

```bash
brew install elasticsearch
```

Install kibana via homebrew:

```bash
brew install kibana
```

Install logstash via homebrew:

```bash
brew install logstash
```

Make sure that all the services are up and running:

```bash
brew services start kibana
brew services start elasticsearch
brew services start logstash
```

You should see all the services up and running:

```bash
brew services list
```

Make sure that the elastic stack is up and running:

```curl
curl -X GET 127.0.0.1:9200
{
  "name" : "M43AiJz",
  "cluster_name" : "elasticsearch_jean-marcelbelmont",
  "cluster_uuid" : "8dHbumffQ3iAOUU32Sx1Jw",
  "version" : {
    "number" : "6.2.2",
    "build_hash" : "10b1edd",
    "build_date" : "2018-02-16T19:01:30.685723Z",
    "build_snapshot" : false,
    "lucene_version" : "7.2.1",
    "minimum_wire_compatibility_version" : "5.6.0",
    "minimum_index_compatibility_version" : "5.0.0"
  },
  "tagline" : "You Know, for Search"
}
```

You can see all the indices available with the following command:

```curl
curl -X GET 127.0.0.1:9200/_cat/indices?v
health status index                           uuid                   pri rep docs.count docs.deleted store.size pri.store.size
yellow open   logstash-2017.05.04             hfY9yL1vTt-IQGE8G591VQ   5   1      16762            0      7.1mb          7.1mb
green  open   .monitoring-es-6-2018.03.18     FWPFXqVBTWmt84HfT4-_Pw   1   0       9903            0      5.7mb          5.7mb
yellow open   logstash-2017.05.02             T6DtDxYJQEm_rwaXg_YLfA   5   1      16278            0        7mb            7mb
green  open   .monitoring-kibana-6-2018.03.18 -tr0EU_0T16tskWdvbGBmg   1   0        314            0    197.5kb        197.5kb
green  open   .triggered_watches              zcs8Nqi2SQ6sUVI2TsBTUg   1   0          0            0     16.2kb         16.2kb
yellow open   app                             oRwQ4n5aSOqFZ_sWPg_iuA   5   1          1            0      5.4kb          5.4kb
yellow open   tags                            CY18aLQuQt-VsLZDce2kZw   5   1       1296            0    342.4kb        342.4kb
yellow open   movies                          141JB-VbTxylwb6IFP6JKA   5   1       9125            0      1.3mb          1.3mb
yellow open   logstash-2017.04.30             TH83PAHZTL6utR_wgxXU9A   5   1      14166            0      6.3mb          6.3mb
yellow open   ratings                         4QXKvkNsTmGVayapswxwcA   5   1     100004            0     13.4mb         13.4mb
green  open   .watches                        XMyhvohjTiSj1sG-ek97nw   1   0          6            0     33.3kb         33.3kb
yellow open   logstash-2017.05.01             8Cs8mEyXSuW_jTQXdUXubg   5   1      15948            0      6.9mb          6.9mb
yellow open   logs                            lXA5H-V_QSOrySGg6r6R1Q   5   1          1            0      5.8kb          5.8kb
yellow open   logstash-2017.05.05             e1hwj1ddQDK69ZmKaRFi5A   5   1      18646            0      7.7mb          7.7mb
yellow open   testindex                       RaI2XDStSBqDEz-jgl-5RQ   3   1          0            0       783b           783b
green  open   .watcher-history-7-2018.03.18   mfe3zaV3Qvib5XLPBA1W0A   1   0        502            0    730.4kb        730.4kb
green  open   .monitoring-alerts-6            IdSXd-D9RNqlGYXcWThGiA   1   0          1            0      6.1kb          6.1kb
yellow open   shakespeare                     EjHjJW8WTpCe4ZyeOhAJXA   5   1     111396            0     21.6mb         21.6mb
green  open   .kibana                         aZ2Fhv-PSsOf5BUTlDdhkw   1   0        143            1    142.5kb        142.5kb
yellow open   logstash-2017.05.03             nklpkzxxSRCWOL990fD2Qg   5   1      21172            0      9.3mb          9.3mb
```

As you can tell there is a list of indices here

We will add some movies via the go script I created:

```curl
curl -X PUT -H "Content-Type: application/json" 127.0.0.1:9200/_bulk?pretty --data-binary @movies.json
```

This imports movies json file into elastic

Here we will create a new movie entry to movie index:

```curl
curl -X POST -H "Content-Type: application/json" 127.0.0.1:9200/movies/movie/180895/_update -d '
{
  "doc" : {
    "title": "BoneHead Man"
  }
}'
```

Here is an example of using filters in elastic search:

```curl
curl -X GET -H 'Content-Type: application/json' '127.0.0.1:9200/logstash-2017.05.04/_search?size=0&pretty' -d '
{
  "query": {
    "match": {
    "agent": "Googlebot"
   }
},
  "aggs": {
    "timestamp": {
      "date_histogram": {
        "field": "@timestamp", "interval": "minute"
      }
    }
  }
}'
```

This query will query and match via a user agent of `Googlebot` and will do an aggregation based on timestamp and use an interval of a minute

Instead of using `curl` you can use the Kibana Dashboard to help you do both queries and create indices and list indices and much more

![Kibana Dashboard](images/kibana-dev-tools.png)
