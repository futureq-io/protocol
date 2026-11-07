# Protocol

`proto/consumer.proto` defines the consumer stream. `SubscribeInit.auto_ack`
defaults to `false`, preserving manual ACK/NACK behavior. A worker may opt in
with `true`; the broker can reject the stream with `PermissionDenied` according
to its configuration. The setting belongs to the stream, so reconnects must
request it again; durable subscription creation is unchanged.

An accepting broker sends initial metadata `subscription-auto-ack: true` along
with `subscription-generation`. Clients must require this confirmation before
omitting ACK/NACK frames, so older brokers cannot silently ignore the option.
On an accepted auto-ack stream the broker completes each recipient after a
successful gRPC send, without waiting for application processing, and ignores
all manual ACKs and NACKs. Handler errors cannot request redelivery. Other
recipients retain their independent delivery obligations. A failed send is not
acknowledged; a crash before completion is persisted can still cause duplicates.

Regenerate the consumer Go bindings with protoc and protoc-gen-go v1.36.11:

```sh
protoc -I proto --go_out=proto/go --go_opt=paths=source_relative proto/consumer.proto
```

