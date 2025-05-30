# AMQP Output Plugin

This plugin writes to an Advanced Message Queuing Protocol v0.9.1 broker.
A prominent implementation of this protocol is [RabbitMQ][rabbitmq].

> [!NOTE]
> This plugin does not bind the AMQP exchange to a queue.

For an introduction check the [AMQP concepts page][amqp_concepts] and the
[RabbitMQ getting started guide][rabbitmq_getting_started].

⭐ Telegraf v0.1.9
🏷️ messaging
💻 all

[amqp_concepts]: https://www.rabbitmq.com/tutorials/amqp-concepts.html
[rabbitmq]: https://www.rabbitmq.com
[rabbitmq_getting_started]: https://www.rabbitmq.com/getstarted.html

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md#plugins

## Secret-store support

This plugin supports secrets from secret-stores for the `username` and
`password` option.
See the [secret-store documentation][SECRETSTORE] for more details on how
to use them.

[SECRETSTORE]: ../../../docs/CONFIGURATION.md#secret-store-secrets

## Configuration

```toml @sample.conf
# Publishes metrics to an AMQP broker
[[outputs.amqp]]
  ## Brokers to publish to.  If multiple brokers are specified a random broker
  ## will be selected anytime a connection is established.  This can be
  ## helpful for load balancing when not using a dedicated load balancer.
  brokers = ["amqp://localhost:5672/influxdb"]

  ## Maximum messages to send over a connection.  Once this is reached, the
  ## connection is closed and a new connection is made.  This can be helpful for
  ## load balancing when not using a dedicated load balancer.
  # max_messages = 0

  ## Exchange to declare and publish to.
  exchange = "telegraf"

  ## Exchange type; common types are "direct", "fanout", "topic", "header", "x-consistent-hash".
  # exchange_type = "topic"

  ## If true, exchange will be passively declared.
  # exchange_passive = false

  ## Exchange durability can be either "transient" or "durable".
  # exchange_durability = "durable"

  ## Additional exchange arguments.
  # exchange_arguments = { }
  # exchange_arguments = {"hash_property" = "timestamp"}

  ## Authentication credentials for the PLAIN auth_method.
  # username = ""
  # password = ""

  ## Auth method. PLAIN and EXTERNAL are supported
  ## Using EXTERNAL requires enabling the rabbitmq_auth_mechanism_ssl plugin as
  ## described here: https://www.rabbitmq.com/plugins.html
  # auth_method = "PLAIN"

  ## Metric tag to use as a routing key.
  ##   ie, if this tag exists, its value will be used as the routing key
  # routing_tag = "host"

  ## Static routing key.  Used when no routing_tag is set or as a fallback
  ## when the tag specified in routing tag is not found.
  # routing_key = ""
  # routing_key = "telegraf"

  ## Delivery Mode controls if a published message is persistent.
  ##   One of "transient" or "persistent".
  # delivery_mode = "transient"

  ## Static headers added to each published message.
  # headers = { }
  # headers = {"database" = "telegraf", "retention_policy" = "default"}

  ## Connection timeout.  If not provided, will default to 5s.  0s means no
  ## timeout (not recommended).
  # timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Optional Proxy Configuration
  # use_proxy = false
  # proxy_url = "localhost:8888"

  ## If true use batch serialization format instead of line based delimiting.
  ## Only applies to data formats which are not line based such as JSON.
  ## Recommended to set to true.
  # use_batch_format = false

  ## Content encoding for message payloads, can be set to "gzip" to or
  ## "identity" to apply no encoding.
  ##
  ## Please note that when use_batch_format = false each amqp message contains only
  ## a single metric, it is recommended to use compression with batch format
  ## for best results.
  # content_encoding = "identity"

  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  # data_format = "influx"
```

### Routing

If `routing_tag` is set, and the tag is defined on the metric, the value of the
tag is used as the routing key.  Otherwise the value of `routing_key` is used
directly.  If both are unset the empty string is used.

Exchange types that do not use a routing key, `direct` and `header`, always use
the empty string as the routing key.

Metrics are published in batches based on the final routing key.

### Proxy

If you want to use a proxy, you need to set `use_proxy = true`. This will
use the system's proxy settings to determine the proxy URL. If you need to
specify a proxy URL manually, you can do so by using `proxy_url`, overriding
the system settings.
