package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	ApacheCommonLog = "%s - %s [%s] \"%s %s %s\" %d %d"
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	ApacheCombinedLog = "%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\""
	// ApacheErrorLog : [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] [client %s:%d] %s"
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3164Log = "<%d>%s %s %s[%d]: %s"
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	RFC5424Log = "<%d>%d %s %s %s %d ID%d %s %s"
	// CommonLogFormat : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	CommonLogFormat = "%s - %s [%s] \"%s %s %s\" %d %d"
	// JSONLogFormat : {"host": "{host}", "user-identifier": "{user-identifier}", "datetime": "{datetime}", "method": "{method}", "request": "{request}", "protocol": "{protocol}", "status", {status}, "bytes": {bytes}, "referer": "{referer}"}
	JSONLogFormat = `{"host":"%s", "user-identifier":"%s", "datetime":"%s", "method": "%s", "request": "%s", "protocol":"%s", "status":%d, "bytes":%d, "referer": "%s"}`
	// JavaLogFormat  : 2020-10-19 04:34:15,389 org.apache.skywalking.oap.server.starter.OAPServerBootstrap -571727 [main] ERROR [] - method [HEAD], host [http://localhost:9200], URI [/], status line [HTTP/1.1 503 Service Unavailable]
	// org.elasticsearch.ElasticsearchStatusException: method [HEAD], host [http://localhost:9200], URI [/], status line [HTTP/1.1 503 Service Unavailable]
	// at org.elasticsearch.client.RestHighLevelClient.parseResponseException(RestHighLevelClient.java:625) ~[elasticsearch-rest-high-level-client-6.3.2.jar:6.3.2]
	// at org.elasticsearch.client.RestHighLevelClient.performRequest(RestHighLevelClient.java:535) ~[elasticsearch-rest-high-level-client-6.3.2.jar:6.3.2]
	// at org.elasticsearch.client.RestHighLevelClient.ping(RestHighLevelClient.java:275) ~[elasticsearch-rest-high-level-client-6.3.2.jar:6.3.2]
	// at org.apache.skywalking.oap.server.library.client.elasticsearch.ElasticSearchClient.connect(ElasticSearchClient.java:121) ~[library-client-6.6.0.1.jar:6.6.0.1]
	// at org.apache.skywalking.oap.server.storage.plugin.elasticsearch.StorageModuleElasticsearchProvider.start(StorageModuleElasticsearchProvider.java:131) ~[storage-elasticsearch-plugin-6.6.0.1.jar:6.6.0.1]
	JavaLogFormat = "%s,%d %s.%s.%s.%s.%s.%s.%s -%d [main] ERROR [] - method [%s], host [%s:%d], URI [%s], status line [%s %d %s]\n%s.%s.%s.%s.%s.%s.%s: method [%s] host [%s:%d], URI [%s], status line [%s %d %s]\n at %s.%s.%s.%s.%s.%s.%s(%s.java:%d) ~ [%s-%s-%s-%d.%d.%d.%d.jar:%d.%d.%d.%d]\n at %s.%s.%s.%s.%s.%s.%s(%s.java:%d) ~ [%s-%s-%s-%d.%d.%d.%d.jar:%d.%d.%d.%d]"
)

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCommonLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCombinedLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(30, 100000),
		gofakeit.URL(),
		gofakeit.UserAgent(),
	)
}

// NewApacheErrorLog creates a log string with apache error log format
func NewApacheErrorLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheErrorLog,
		t.Format(ApacheError),
		gofakeit.Word(),
		gofakeit.LogLevel("apache"),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		gofakeit.IPv4Address(),
		gofakeit.Number(1, 65535),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC3164Log creates a log string with syslog (RFC3164) format
func NewRFC3164Log(t time.Time) string {
	return fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format(RFC3164),
		strings.ToLower(gofakeit.Username()),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC5424Log creates a log string with syslog (RFC5424) format
func NewRFC5424Log(t time.Time) string {
	return fmt.Sprintf(
		RFC5424Log,
		gofakeit.Number(0, 191),
		gofakeit.Number(1, 3),
		t.Format(RFC5424),
		gofakeit.DomainName(),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 1000),
		"-", // TODO: structured data
		gofakeit.HackerPhrase(),
	)
}

// NewCommonLogFormat creates a log string with common log format
func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewJSONLogFormat creates a log string with json log format
func NewJSONLogFormat(t time.Time) string {
	return fmt.Sprintf(
		JSONLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.URL(),
	)
}

// NewJavaLogFormat creates a log string with json log format
func NewJavaLogFormat(t time.Time) string {
	return fmt.Sprintf(
		JavaLogFormat,
		t.Format(RFC3164),
		gofakeit.Number(111, 999),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Number(111111, 999999),
		gofakeit.HTTPMethod(),
		gofakeit.URL(),
		gofakeit.Number(1024, 30000),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Sentence(3),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.HTTPMethod(),
		gofakeit.URL(),
		gofakeit.Number(1024, 30000),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Sentence(3),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Number(1, 9999),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Number(1, 9999),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
		gofakeit.Number(1, 9),
	)
}
